package store

import (
	"platform/src/github.com/winmorre/platform/http/actionresults"
	"platform/src/github.com/winmorre/platform/http/handling"
	"sportsstore/models"
	"sportsstore/store/cart"
)

type CartHandler struct {
	models.Repository
	cart.Cart
	handling.URLGenerator
}

type CartTemplateContext struct {
	cart.Cart
	ProductListUrl string
	CartUrl        string
	CheckoutUrl    string
	RemoveUrl      string
}

func (ch CartHandler) GetCart() actionresults.ActionResult {
	return actionresults.NewTemplateAction("cart.html", CartTemplateContext{
		Cart:           ch.Cart,
		ProductListUrl: ch.mustGenerateUrl(ProductHandler.GetProducts, 0, 1),
		RemoveUrl:      ch.mustGenerateUrl(CartHandler.PostRemoveFromCart),
		CheckoutUrl:    ch.mustGenerateUrl(OrderHandler.GetCheckout),
	})
}

type CartProductReference struct {
	ID int
}

func (ch CartHandler) PostAddToCart(ref CartProductReference) actionresults.ActionResult {
	p := ch.Repository.GetProduct(ref.ID)
	ch.Cart.AddProduct(p)
	return actionresults.NewRedirectAction(ch.mustGenerateUrl(CartHandler.GetCart))
}

func (ch CartHandler) PostRemoveFromCart(ref CartProductReference) actionresults.ActionResult {
	ch.Cart.RemoveLineForProduct(ref.ID)
	return actionresults.NewRedirectAction(ch.mustGenerateUrl(CartHandler.GetCart))
}

func (ch CartHandler) mustGenerateUrl(method interface{}, data ...interface{}) string {
	url, err := ch.URLGenerator.GenerateUrl(method, data...)
	if err != nil {
		panic(err)
	}
	return url
}

func (ch CartHandler) GetWidget() actionresults.ActionResult {
	return actionresults.NewTemplateAction("cart_widget.html", CartTemplateContext{
		Cart:    ch.Cart,
		CartUrl: ch.mustGenerateUrl(CartHandler.GetCart),
	})
}
