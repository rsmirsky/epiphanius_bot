package requests

type Holiday struct {
	Id          int    `json:"id"`
	Name        string `json:"name"`
	Description string `json:"description"`
	Date        string `json:"date"` //було стринг
}
