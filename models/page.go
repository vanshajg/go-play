package models

type Page struct {
	Data             *[]Comment `json:"data"`
	Last             bool       `json:"last"`
	TotalElements    int        `json:"totalElements"`
	TotalPages       int        `json:"totalPages"`
	Size             int        `json:"size"`
	Page             int        `json:"page"`
	NumberOfElements int        `json:"numberOfElements"`
}
