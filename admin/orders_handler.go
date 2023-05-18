package admin

type OrdersHandler struct{}

func (oh OrdersHandler) GetData() string {
	return "This is the orders handler"
}
