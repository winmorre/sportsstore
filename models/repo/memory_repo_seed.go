package repo

import (
	"fmt"
	"math/rand"
	"sportsstore/models"
)

func (r *MemoryRepo) Seed() {
	r.categories = make([]models.Category, 3)

	for i := 0; i < 3; i++ {
		catName := fmt.Sprintf("Category_%v", i+1)
		r.categories[i] = models.Category{ID: i + 1, CategoryName: catName}
	}

	for i := 0; i < 20; i++ {
		name := fmt.Sprintf("Product_%v", i+1)
		price := rand.Float64() * float64(rand.Intn(500))

		cat := &r.categories[rand.Intn(len(r.categories))]

		r.products = append(r.products, models.Product{
			ID:          i + 1,
			Name:        name,
			Price:       price,
			Description: fmt.Sprintf("%v (%v)", cat.CategoryName),
			Category:    cat,
		})
	}
}
