package model

import "time"

type Order struct {
	OrderId             int
	CustomerId          int
	EmployeeId          int
	OrderDate           time.Time
	PurchaseOrderNumber string
	ShipDate            time.Time
	ShippingMethodId    int
	FreightCharge       float64
	Taxes               float64
	PaymentReceived     int
	Comment             string
	CreatedAt           time.Time
	CreatedBy           int
	UpdatedAt           time.Time
	UpdatedBy           int
}
