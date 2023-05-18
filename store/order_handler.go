package store

import (
	"encoding/json"
	"platform/src/github.com/winmorre/platform/http/actionresults"
	"platform/src/github.com/winmorre/platform/http/handling"
	"platform/src/github.com/winmorre/platform/sessions"
	"platform/src/github.com/winmorre/platform/validation"
	"sportsstore/models"
	"sportsstore/store/cart"
	"strings"
)

type OrderHandler struct {
	cart.Cart
	sessions.Session
	Repository   models.Repository
	URLGenerator handling.URLGenerator
	validation.Validator
}

type OrderTemplateContext struct {
	models.ShippingDetails
	ValidationErrors [][]string
	CancelUrl        string
}

func (oh OrderHandler) GetCheckout() actionresults.ActionResult {
	context := OrderTemplateContext{}
	jsonData := oh.Session.GetValueDefault("checkout_details", "")

	if jsonData != nil {
		json.NewDecoder(strings.NewReader(jsonData.(string))).Decode(&context)
	}

	context.CancelUrl = mustGenerateUrl(oh.URLGenerator, CartHandler.GetCart)
	return actionresults.NewTemplateAction("checkout.html", context)
}

func (oh OrderHandler) PostCheckout(details models.ShippingDetails) actionresults.ActionResult {
	valid, errors := oh.Validator.Validate(details)

	if !valid {
		ctx := OrderTemplateContext{
			ShippingDetails:  details,
			ValidationErrors: [][]string{},
		}

		for _, err := range errors {
			ctx.ValidationErrors = append(ctx.ValidationErrors, []string{err.FieldName, err.Error.Error()})
		}
		builder := strings.Builder{}
		json.NewEncoder(&builder).Encode(ctx)

		oh.Session.SetValue("checkout_details", builder.String())
		redirectUrl := mustGenerateUrl(oh.URLGenerator, OrderHandler.GetCheckout)
		return actionresults.NewRedirectAction(redirectUrl)
	} else {
		oh.Session.SetValue("check_details", "")
	}
	order := models.Order{
		ShippingDetails: details,
		Products:        []models.ProductSelection{},
	}

	for _, cl := range oh.Cart.GetLines() {
		order.Products = append(order.Products, models.ProductSelection{
			Quantity: cl.Quantity,
			Product:  cl.Product,
		})
	}

	oh.Repository.SaveOrder(&order)
	oh.Cart.Reset()

	targetUrl, _ := oh.URLGenerator.GenerateUrl(OrderHandler.GetSummary, order.ID)
	return actionresults.NewRedirectAction(targetUrl)
}

func (oh OrderHandler) GetSummary(id int) actionresults.ActionResult {
	targetUrl, _ := oh.URLGenerator.GenerateUrl(ProductHandler.GetProducts, 0, 1)

	return actionresults.NewTemplateAction("checkout_summary.html", struct {
		ID        int
		TargetUrl string
	}{ID: id, TargetUrl: targetUrl})
}
