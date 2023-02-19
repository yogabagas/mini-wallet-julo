package service

import "time"

type CreateEmployeesReq struct {
	EmployeeId int       `json:"employee_id" csv:"employee_id"`
	FirstName  string    `json:"first_name" csv:"first_name"`
	LastName   string    `json:"last_name" csv:"last_name"`
	Title      string    `json:"title" csv:"title"`
	WorkPhone  string    `json:"work_phone" csv:"work_phone"`
	IsDeleted  bool      `json:"-" csv:"-"`
	CreatedAt  time.Time `json:"-" csv:"-"`
	CreatedBy  int       `json:"created_by" csv:"created_by"`
	UpdatedAt  time.Time `json:"-" csv:"-"`
	UpdatedBy  int       `json:"updated_by" csv:"updated_by"`
}

type CreateEmployeesResp struct {
	Status string `json:"status"`
}
