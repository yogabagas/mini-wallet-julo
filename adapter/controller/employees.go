package controller

import (
	"context"

	"os"

	"github.com/gocarina/gocsv"
	"github.com/yogabagas/jatis-BE/domain/service"
	employees "github.com/yogabagas/jatis-BE/service/employees/usecase"
)

type EmployeesControllerImpl struct {
	employeesUsecase employees.EmployeesUsecase
}

type EmployeesController interface {
	CreateEmployees(ctx context.Context, req []service.CreateEmployeesReq) (resp service.CreateEmployeesResp, err error)
}

func NewEmployeesController(employeesUsecase employees.EmployeesUsecase) EmployeesController {
	return &EmployeesControllerImpl{employeesUsecase: employeesUsecase}
}

func (cc *EmployeesControllerImpl) CreateEmployees(ctx context.Context, req []service.CreateEmployeesReq) (resp service.CreateEmployeesResp, err error) {

	file, err := os.Open("csv_employees.csv")
	if err != nil {
		return service.CreateEmployeesResp{
			Status: "failed",
		}, err
	}

	if err := gocsv.Unmarshal(file, &req); err != nil {
		return service.CreateEmployeesResp{
			Status: "failed",
		}, err
	}

	return cc.employeesUsecase.CreateEmployees(ctx, req)
}
