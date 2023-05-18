package repo

import (
	"math"
	//"platform/src/github.com/winmorre/platform/services"
	"sportsstore/models"
)

//func RegisterMemoryRepoService() {
//	services.AddSingleton(func() models.Repository {
//		repo := &MemoryRepo{}
//		repo.Seed()
//		return repo
//	})
//}

type MemoryRepo struct {
	products   []models.Product
	categories []models.Category
}

func (r *MemoryRepo) GetProduct(id int) (product models.Product) {
	for _, p := range r.products {
		if p.ID == id {
			product = p
			return
		}
	}
	return
}

func (r *MemoryRepo) GetProducts() (results []models.Product) {
	return r.products
}

func (r *MemoryRepo) GetCategories() (results []models.Category) {
	return r.categories
}

func (r *MemoryRepo) GetProductPage(page, pageSize int) ([]models.Product, int) {
	return getPage(r.products, page, pageSize), len(r.products)
}

func getPage(src []models.Product, page, pageSize int) []models.Product {
	start := (page - 1) * pageSize
	if page > 0 && len(src) > start {
		end := (int)(math.Min((float64)(len(src)), (float64)(start+pageSize)))
		return src[start:end]
	}
	return []models.Product{}
}

func (r *MemoryRepo) GetProductPageCategory(category, page, pageSize int) (products []models.Product, totalAvailable int) {
	if category == 0 {
		return r.GetProductPage(page, pageSize)
	} else {
		filteredProducts := make([]models.Product, 0, len(r.products))

		for _, p := range r.products {
			if p.Category.ID == category {
				filteredProducts = append(filteredProducts, p)
			}
		}
		return getPage(filteredProducts, page, pageSize), len(filteredProducts)
	}
}
