package service

import "time"

type CreateProductsReq struct {
	ProductName string    `json:"product_name" csv:"product_name"`
	UnitPrice   float64   `json:"unit_price" csv:"unit_price"`
	InStock     int       `json:"in_stock" csv:"in_stock"`
	CreatedAt   time.Time `json:"-" csv:"-"`
	CreatedBy   int       `json:"created_by" csv:"created_by"`
	UpdatedAt   time.Time `json:"-" csv:"-"`
	UpdatedBy   int       `json:"updated_by" csv:"updated_by"`
}

type CreateProductsResp struct {
	Status string `json:"status"`
}
