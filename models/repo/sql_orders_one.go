package repo

import "sportsstore/models"

func (sr *SqlRepository) GetOrder(id int) (order models.Order) {
	order = models.Order{Products: []models.ProductSelection{}}
	row := sr.Commands.GetOrder.QueryRowContext(sr.Context, id)

	if row.Err() == nil {
		err := row.Scan(&order.ID, &order.Name, &order.Name, &order.StreetAddr, &order.City, &order.Zip, &order.Country, &order.Shipped)

		if err != nil {
			sr.Logger.Panicf("Cannot scan order data: %v", err.Error())
			return
		}
		lineRows, err := sr.Commands.GetOrderLines.QueryContext(sr.Context, id)
		if err == nil {
			for lineRows.Next() {
				ps := models.ProductSelection{Product: models.Product{Category: &models.Category{}}}

				err = lineRows.Scan(&ps.Quantity, &ps.Product.ID, &ps.Product.Name, &ps.Product.Description, &ps.Product.Price, &ps.Product.Category.ID, &ps.Product.Category.CategoryName)

				if err == nil {
					order.Products = append(order.Products, ps)
				} else {
					sr.Logger.Panicf("Cannot scan order line data: %v", err.Error())
				}
			}
		} else {
			sr.Logger.Panicf("Cannot exec GetOrderLines command: %v", err.Error())
		}
	} else {
		sr.Logger.Panicf("Cannot exec GetOrder Command: %v", row.Err().Error())
	}
	return
}
