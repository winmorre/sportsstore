package repo

import "sportsstore/models"

func (sr *SqlRepository) SaveProduct(p *models.Product) {

	if p.ID == 0 {
		result, err := sr.Commands.SaveProduct.ExecContext(sr.Context, p.Name, p.Description, p.Category.ID, p.Price)

		if err == nil {
			id, err := result.LastInsertId()

			if err == nil {
				p.ID = int(id)
				return
			} else {
				sr.Logger.Panicf("Cannot  get inserted ID: %v", err.Error())
			}
		} else {
			sr.Logger.Panicf("Cannot exec SaveProduct command: %v", err.Error())
		}
	} else {
		result, err := sr.Commands.UpdateProduct.ExecContext(sr.Context, p.Name, p.Description, p.Category.ID, p.Price, p.ID)

		if err == nil {
			affected, err := result.RowsAffected()
			if err == nil && affected != 1 {
				sr.Logger.Panicf("Got unexpected row affected: %v", affected)
			} else if err != nil {
				sr.Logger.Panicf("Cannot get rows affected: %v", err)
			}
		} else {
			sr.Logger.Panicf("Cannot exec Update command: %v", err.Error())
		}
	}
}
