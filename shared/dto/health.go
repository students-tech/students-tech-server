package dto

const (
	OK = "OK"
)

type (
	Status struct {
		Name   string      `json:"name"`
		Status string      `json:"status"`
		Data   interface{} `json:"data"`
	}
)
