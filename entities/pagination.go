package entities

type Pagination struct {
	Limit uint `json:"limit"`
	Page  uint `json:"page"`
}