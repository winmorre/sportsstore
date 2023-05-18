package cart

import "sportsstore/models"

type CartLine struct {
	models.Product
	Quantity int
}

func (cl *CartLine) GetLineTotal() float64 {
	return cl.Price * float64(cl.Quantity)
}

type Cart interface {
	AddProduct(models.Product)
	GetLines() []*CartLine
	RemoveLineForProduct(id int)
	GetItemCount() int
	GetTotal() float64

	Reset()
}

type BasicCart struct {
	lines []*CartLine
}

func (bc *BasicCart) AddProduct(p models.Product) {
	for _, line := range bc.lines {
		if line.Product.ID == p.ID {
			line.Quantity++
			return
		}
	}
	bc.lines = append(bc.lines, &CartLine{Product: p, Quantity: 1})
}

func (bc *BasicCart) GetLines() []*CartLine {
	return bc.lines
}

func (bc *BasicCart) RemoveLineForProduct(id int) {
	for index, line := range bc.lines {
		if line.Product.ID == id {
			bc.lines = append(bc.lines[0:index], bc.lines[index+1:]...)
		}
	}
}

func (bc *BasicCart) GetItemCount() (total int) {
	for _, l := range bc.lines {
		total += l.Quantity
	}
	return
}

func (bc *BasicCart) GetTotal() (total float64) {
	for _, line := range bc.lines {
		total += float64(line.Quantity) * line.Product.Price
	}
	return
}

func (bc *BasicCart) Reset() {
	bc.lines = []*CartLine{}
}
