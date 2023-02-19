package model

import "time"

type Customer struct {
	CustomerId      int
	CompanyName     string
	FirstName       string
	LastName        string
	BillingAddress  string
	City            string
	Province        string
	ZipCode         string
	Email           string
	CompanyWebsite  string
	PhoneNumber     string
	FaxNumber       string
	ShipAddress     string
	ShipCity        string
	ShipProvince    string
	ShipZipCode     string
	ShipPhoneNumber string
	CreatedAt       time.Time
	CreatedBy       int
	UpdatedAt       time.Time
	UpdatedBy       int
}
