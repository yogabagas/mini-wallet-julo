package model

import "time"

type ShippingMethod struct {
	ShippingMethodId   int
	ShippingMethodName string
	CreatedAt          time.Time
	CreatedBy          int
	UpdatedAt          time.Time
	UpdatedBy          int
}
