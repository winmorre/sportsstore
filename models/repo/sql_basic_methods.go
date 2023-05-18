package repo

import "sportsstore/models"

func (sr *SqlRepository) GetProduct(id int) (p models.Product) {
	row := sr.Commands.GetProduct.QueryRowContext(sr.Context, id)

	if row.Err() == nil {
		var err error
		if p, err = scanProduct(row); err != nil {
			sr.Logger.Panicf("Cannot scan data : %v", err.Error())
		}
	} else {
		sr.Logger.Panicf("Cannot exec GetProduct command: %v", row.Err().Error())
	}
	return
}

func (sr *SqlRepository) GetProducts() (results []models.Product) {
	rows, err := sr.Commands.GetProducts.QueryContext(sr.Context)

	if err == nil {
		if results, err = scanProducts(rows); err != nil {
			sr.Logger.Panicf("Cannot scan data: %v", err.Error())
			return
		}
	} else {
		sr.Logger.Panicf("Cannot exec GetProducts command: %V", err)
	}
	return
}

func (sr *SqlRepository) GetCategories() []models.Category {
	results := make([]models.Category, 0, 10)
	rows, err := sr.Commands.GetCategories.QueryContext(sr.Context)
	if err == nil {
		for rows.Next() {
			c := models.Category{}
			if err := rows.Scan(&c.ID, &c.CategoryName); err != nil {
				sr.Logger.Panicf("Cannot scan data: %v", err.Error())
			}
			results = append(results, c)
		}
	} else {
		sr.Logger.Panicf("Cannot exec GetCategories command: %v", err)
	}
	return results
}
