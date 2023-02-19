package service

import "time"

type CreateOrderDetailsReq struct {
	OrderId   int       `json:"order_id" csv:"order_id"`
	ProductId int       `json:"product_id" csv:"product_id"`
	Qty       int       `json:"qty" csv:"qty"`
	UnitPrice float64   `json:"unit_price" csv:"unit_price"`
	Discount  float64   `json:"discount" csv:"discount"`
	CreatedAt time.Time `json:"-" csv:"-"`
	CreatedBy int       `json:"created_by" csv:"created_by"`
	UpdatedAt time.Time `json:"-" csv:"-"`
	UpdatedBy int       `json:"updated_by" csv:"updated_by"`
}

type CreateOrderDetailsResp struct {
	Status string `json:"status"`
}

type GetOrderDetailsResp struct {
	CustomerName   string  `json:"customer_name"`
	EmployeeName   string  `json:"employee_name"`
	ShippingMethod string  `json:"shipping_method"`
	Qty            int     `json:"qty"`
	Discount       float64 `json:"discount"`
	UnitPrice      float64 `json:"unit_price"`
	SubTotalPrice  float64 `json:"subtotal_price"`
	TotalPayment   float64 `json:"total_payment"`
}
