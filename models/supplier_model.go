package models

// Supplier struct is model outcome for data
type Supplier struct {
	ID      int64    `form:"id" json:"id"`
	Code    string   `form:"code" json:"code"`
	Name    string   `form:"name" json:"name"`
	Address string   `form:"address" json:"address"`
	Status  string   `form:"status" json:"status"`
	Faktur  []Faktur `json:"faktur"`
}
