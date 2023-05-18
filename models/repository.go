package models

type Repository interface {
	GetProduct(id int) Product
	GetProducts() []Product
	GetCategories() []Category
	GetProductPage(page, pageSize int) (products []Product, totalAvailable int)
	GetProductPageCategory(categoryId, page, pageSize int) (products []Product, totalAvailability int)
	GetOrder(id int) Order
	GetOrders() []Order
	SaveOrder(order *Order)
	Seed()
	SaveProduct(*Product)
}
