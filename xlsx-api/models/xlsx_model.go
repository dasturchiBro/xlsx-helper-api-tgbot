package models

type XLSXRequest struct {
	Header []string `json:"header"`
	Criteria []string `json:"criteria"`
	Students []Student `json:"students"`
}

type Student struct {
	Name string `json:"name"`
	Points []float64 `json:"points"`
}