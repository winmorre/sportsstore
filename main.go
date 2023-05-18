package main

import (
	"platform/src/github.com/winmorre/platform/http"
	"platform/src/github.com/winmorre/platform/http/handling"
	"platform/src/github.com/winmorre/platform/pipeline"
	"platform/src/github.com/winmorre/platform/pipeline/basic"
	"platform/src/github.com/winmorre/platform/services"
	"platform/src/github.com/winmorre/platform/sessions"
	"sportsstore/admin"
	"sportsstore/models/repo"
	"sportsstore/store"
	"sportsstore/store/cart"
	"sync"
)

func registerServices() {
	services.RegisterDefaultServices()
	//repo.RegisterMemoryRepoService()
	repo.RegisterSqlRepositoryService()
	sessions.RegisterSessionService()
	cart.RegisterCartService()
}

func createPipeline() pipeline.RequestPipeline {
	return pipeline.CreatePipeline(
		&basic.ServicesComponent{},
		&basic.LoggingComponent{},
		&basic.ErrorComponent{},
		&basic.StaticFileComponent{},
		&sessions.SessionComponent{},
		handling.NewRouter(
			handling.HandlerEntry{Handler: store.ProductHandler{}},
			handling.HandlerEntry{Handler: store.CategoryHandler{}},
			handling.HandlerEntry{Handler: store.CartHandler{}},
			handling.HandlerEntry{Handler: store.OrderHandler{}},
			handling.HandlerEntry{Prefix: "admin", Handler: admin.AdminHandler{}},
			handling.HandlerEntry{Prefix: "admin", Handler: admin.ProductsHandler{}},
			handling.HandlerEntry{Prefix: "admin", Handler: admin.CategoriesHandler{}},
			handling.HandlerEntry{Prefix: "admin", Handler: admin.OrdersHandler{}},
			handling.HandlerEntry{Prefix: "admin", Handler: admin.DatabaseHandler{}},
		).AddMethodAlias("/", store.ProductHandler.GetProducts, 0, 1).
			AddMethodAlias("/products[/]?[A-z0-9]*?", store.ProductHandler.GetProducts, 0, 1).
			AddMethodAlias("/admin[/]", admin.AdminHandler.GetSection, ""),
	)
}

func main() {
	registerServices()

	results, err := services.Call(http.Serve, createPipeline())

	if err == nil {
		(results[0].(*sync.WaitGroup)).Wait()
	} else {
		panic(err)
	}
}
