package model

import "time"

type Product struct {
	ProductID   int
	ProductName string
	UnitPrice   float64
	InStock     int
	CreatedAt   time.Time
	CreatedBy   int
	UpdatedAt   time.Time
	UpdatedBy   int
}
