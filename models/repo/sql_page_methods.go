package repo

import "sportsstore/models"

func (sr *SqlRepository) GetProductPage(page, pageSize int) (products []models.Product, totalAvailable int) {
	rows, err := sr.Commands.GetPage.QueryContext(sr.Context, pageSize, (page*pageSize)-pageSize)
	if err == nil {
		if products, err = scanProducts(rows); err != nil {
			sr.Logger.Panicf("Cannot scan ata: %v", err.Error())
			return
		}
	} else {
		sr.Logger.Panicf("Cannot exec GetProductPage command: %v", err)
		return
	}

	row := sr.Commands.GetPageCount.QueryRowContext(sr.Context)

	if row.Err() == nil {
		if err := row.Scan(&totalAvailable); err != nil {
			sr.Logger.Panicf("Cannot scan data: %v", err.Error())
		}
	} else {
		sr.Logger.Panicf("Cannot exec GetPageCount command: %v", row.Err().Error())
	}
	return
}

func (sr *SqlRepository) GetProductPageCategory(categoryId, page, pageSize int) (products []models.Product, totalAvailable int) {

	if categoryId == 0 {
		return sr.GetProductPage(page, pageSize)
	}

	rows, err := sr.Commands.GetCategoryPage.QueryContext(sr.Context, categoryId, pageSize, (page*pageSize)-pageSize)

	if err == nil {
		if products, err = scanProducts(rows); err != nil {
			sr.Logger.Panicf("Cannot scan data: %v", err.Error())
			return
		}
	} else {
		sr.Logger.Panicf("Cannot exec GetProductPage command: %v", err)
		return
	}

	row := sr.Commands.GetCategoryPageCount.QueryRowContext(sr.Context, categoryId)

	if row.Err() == nil {
		if err := row.Scan(&totalAvailable); err != nil {
			sr.Logger.Panicf("Cannot scn data: %v", err.Error())
		}
	} else {
		sr.Logger.Panicf("Cannot exec GetCategoryPageCount COMMAND: %V", row.Err().Error())
	}
	return
}
