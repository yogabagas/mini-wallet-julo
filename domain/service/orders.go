package service

import "time"

type CreateOrdersReq struct {
	CustomerId          int       `json:"customer_id" csv:"customer_id"`
	EmployeeId          int       `json:"employee_id" csv:"employee_id"`
	OrderDate           time.Time `json:"order_date" csv:"order_date"`
	PurchaseOrderNumber string    `json:"purchase_order_no" csv:"purchase_order_no"`
	ShipDate            time.Time `json:"ship_date" csv:"ship_date"`
	ShippingMethodId    int       `json:"shipping_method_id" csv:"shipping_method_id"`
	FreightCharge       float64   `json:"freight_charge" csv:"freight_charge"`
	Taxes               float64   `json:"taxes" csv:"taxes"`
	PaymentReceived     int       `json:"payment_received" csv:"payment_received"`
	Comment             string    `json:"comment" csv:"comment"`
	CreatedAt           time.Time `json:"-" csv:"-"`
	CreatedBy           int       `json:"created_by" csv:"created_by"`
	UpdatedAt           time.Time `json:"-" csv:"-"`
	UpdatedBy           int       `json:"updated_by" csv:"updated_by"`
}

type CreateOrdersResp struct {
	Status string `json:"status"`
}
