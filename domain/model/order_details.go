package model

import "time"

type OrderDetail struct {
	OrderDetailId int
	OrderId       int
	ProductId     int
	Qty           int
	UnitPrice     float64
	Discount      float64
	CreatedAt     time.Time
	CreatedBy     int
	UpdatedAt     time.Time
	UpdatedBy     int
}

type ReadOrderDetailsResp struct {
	CustomerName   string
	EmployeeName   string
	ShippingMethod string
	Qty            int
	Discount       float64
	UnitPrice      float64
	SubTotalPrice  float64
	TotalPayment   float64
}
