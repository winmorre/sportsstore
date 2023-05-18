package admin

type DatabaseHandler struct{}

func (dh DatabaseHandler) GetData() string {
	return "This is the database handler"
}
