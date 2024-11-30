package models

type Product struct {
	ID          int64  `json:"id"`
	Name        string `json:"name"`
	Code        string `json:"code"`
	Description string `json:"description"`
	Qty         int    `json:"qty"`
	Active      bool   `json:"active"`
	Deleted     bool   `json:"deleted"`
}
