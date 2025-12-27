package models

import (
	"time"
	"encoding/json"
)

type User struct {
	ID int
	CreatedAt time.Time
	UserID int64
	Stage string
}

type Student struct {
	ID int
	Name string
	ClassID int 
	Points json.RawMessage 
	CreatedAt time.Time
	TemplateID int
	UserID int64
}

type Class struct {
	ID int 
	Name string 
	Grade int
	UserID int64
}

type Template struct {
	ID int
	Name string
	ClassID int
	UserID int64 
	Header []string
	Criteria json.RawMessage 
}

type XLSXStudent struct {
	Name string `json:"name"`
	Points []float64 `json:"points"`
}

type XLSXRequest struct {
	Header []string `json:"header"`
	Criteria []string `json:"criteria"`
	Students []XLSXStudent `json:"students"`
}



type URL struct {
	URL string
}