package admin

import (
	"platform/src/github.com/winmorre/platform/http/actionresults"
	"platform/src/github.com/winmorre/platform/http/handling"
	"platform/src/github.com/winmorre/platform/sessions"
	"sportsstore/models"
)

type ProductsHandler struct {
	models.Repository
	handling.URLGenerator
	sessions.Session
}

type ProductTemplateContext struct {
	Products []models.Product
	EditId   int
	EditUrl  string
	SaveUrl  string
}

const PRODUCT_EDIT_KEY string = "product_edit"

func (ph ProductsHandler) GetData() actionresults.ActionResult {
	return actionresults.NewTemplateAction("admin_products.html",
		ProductTemplateContext{
			Products: ph.GetProducts(),
			EditId:   ph.Session.GetValueDefault(PRODUCT_EDIT_KEY, 0).(int),
			EditUrl:  mustGenerateUrl(ph.URLGenerator, ProductsHandler.PostProductEdit),
			SaveUrl:  mustGenerateUrl(ph.URLGenerator, ProductsHandler.PostProductSave),
		})
}

type EditReference struct {
	ID int
}

func (ph ProductsHandler) PostProductEdit(ref EditReference) actionresults.ActionResult {
	ph.Session.SetValue(PRODUCT_EDIT_KEY, ref.ID)

	return actionresults.NewRedirectAction(mustGenerateUrl(ph.URLGenerator, AdminHandler.GetSection, "Products"))
}

type ProductSaveReference struct {
	Id                int
	Name, Description string
	Category          int
	Price             float64
}

func (ph ProductsHandler) PostProductSave(p ProductSaveReference) actionresults.ActionResult {
	ph.Repository.SaveProduct(&models.Product{
		ID: p.Id, Name: p.Name, Description: p.Description,
		Category: &models.Category{ID: p.Category},
		Price:    p.Price,
	})
	ph.Session.SetValue(PRODUCT_EDIT_KEY, 0)
	return actionresults.NewRedirectAction(mustGenerateUrl(ph.URLGenerator, AdminHandler.GetSection, "Products"))
}

func mustGenerateUrl(gen handling.URLGenerator, target interface{}, data ...interface{}) string {
	url, err := gen.GenerateUrl(target, data...)
	if err != nil {
		panic(err)
	}
	return url
}
