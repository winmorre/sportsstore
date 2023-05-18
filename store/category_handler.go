package store

import (
	"platform/src/github.com/winmorre/platform/http/actionresults"
	"platform/src/github.com/winmorre/platform/http/handling"
	"sportsstore/models"
)

type CategoryHandler struct {
	Repository   models.Repository
	URLGenerator handling.URLGenerator
}

type categoryTemplateContext struct {
	Categories       []models.Category
	SelectedCategory int
	CategoryUrlFunc  func(int) string
}

func (ch CategoryHandler) GetButtons(selected int) actionresults.ActionResult {
	return actionresults.NewTemplateAction("category_buttons.html",
		categoryTemplateContext{
			Categories:       ch.Repository.GetCategories(),
			SelectedCategory: selected,
			CategoryUrlFunc:  ch.createCategoryFilterFunction(),
		})
}

func (ch CategoryHandler) createCategoryFilterFunction() func(int) string {
	return func(category int) string {
		url, _ := ch.URLGenerator.GenerateUrl(ProductHandler.GetProducts, category, 1)
		return url
	}
}
