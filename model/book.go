package model

type Book struct {
	Auth     string `json:"auth"`
	Title    string `json:"title"`
	Category string `json:"category"`
	ISBN     string `json:"isbn"`
}
