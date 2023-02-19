package service

import "time"

type CreateShippingMethodsReq struct {
	ShippingMethodName string    `json:"shipping_method" csv:"shipping_method"`
	CreatedAt          time.Time `json:"-" csv:"-"`
	CreatedBy          int       `json:"created_by" csv:"created_by"`
	UpdatedAt          time.Time `json:"-" csv:"-"`
	UpdatedBy          int       `json:"updated_by" csv:"updated_by"`
}

type CreateShippingMethodsResp struct {
	Status string `json:"status"`
}
