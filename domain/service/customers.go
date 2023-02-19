package service

import "time"

type CreateCustomersReq struct {
	CustomerId      int       `json:"customer_id" csv:"-"`
	CompanyName     string    `json:"company_name" csv:"company_name"`
	FirstName       string    `json:"first_name" csv:"first_name"`
	LastName        string    `json:"last_name" csv:"last_name"`
	BillingAddress  string    `json:"billing_address" csv:"billing_address"`
	City            string    `json:"city" csv:"city"`
	Province        string    `json:"province" csv:"province"`
	ZipCode         string    `json:"zip_code" csv:"zip_code"`
	Email           string    `json:"email" csv:"email"`
	CompanyWebsite  string    `json:"company_website" csv:"company_website"`
	PhoneNumber     string    `json:"phone_number" csv:"phone_number"`
	FaxNumber       string    `json:"fax_number" csv:"fax_number"`
	ShipAddress     string    `json:"ship_address" csv:"ship_address"`
	ShipCity        string    `json:"ship_city" csv:"ship_city"`
	ShipProvince    string    `json:"ship_province" csv:"ship_province"`
	ShipZipCode     string    `json:"ship_zip_code" csv:"ship_zip_code"`
	ShipPhoneNumber string    `json:"ship_phone_number" csv:"ship_phone_number"`
	CreatedAt       time.Time `json:"created_at" csv:"-"`
	CreatedBy       int       `json:"created_by" csv:"-"`
	UpdatedAt       time.Time `json:"updated_at" csv:"-"`
	UpdatedBy       int       `json:"updated_by" csv:"-"`
}

type CreateCustomersResp struct {
	Status string `json:"status"`
}
