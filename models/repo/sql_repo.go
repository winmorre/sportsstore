package repo

import (
	"context"
	"database/sql"
	"platform/src/github.com/winmorre/platform/config"
	"platform/src/github.com/winmorre/platform/logging"
)

type SqlRepository struct {
	config.Configuration
	logging.Logger
	Commands SqlCommands
	*sql.DB
	context.Context
}

type SqlCommands struct {
	Init,
	Seed,
	GetProduct,
	GetProducts,
	GetCategories,
	GetPage,
	GetPageCount,
	GetCategoryPage,
	GetCategoryPageCount,
	GetOrder,
	GetOrderLines,
	GetOrders,
	GetOrdersLines,
	SaveOrder,
	SaveProduct,
	UpdateProduct,
	SaveOrderLine *sql.Stmt
}
