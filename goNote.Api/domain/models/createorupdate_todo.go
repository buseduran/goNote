package models

type Todo struct {
	Completed bool   `json:"completed"`
	Body      string `json:"body"`
}
