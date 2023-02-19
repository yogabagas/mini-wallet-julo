package model

import "time"

type Employee struct {
	EmployeeId int
	FirstName  string
	LastName   string
	Title      string
	WorkPhone  string
	IsDeleted  bool
	CreatedAt  time.Time
	CreatedBy  int
	UpdatedAt  time.Time
	UpdatedBy  int
}
