package controller

import (
	"context"

	"os"

	"github.com/gocarina/gocsv"
	"github.com/yogabagas/jatis-BE/domain/service"
	customers "github.com/yogabagas/jatis-BE/service/customers/usecase"
)

type CustomersControllerImpl struct {
	customersUsecase customers.CustomersUsecase
}

type CustomersController interface {
	CreateCustomers(ctx context.Context, req []service.CreateCustomersReq) (resp service.CreateCustomersResp, err error)
}

func NewCustomersController(customersUsecase customers.CustomersUsecase) CustomersController {
	return &CustomersControllerImpl{customersUsecase: customersUsecase}
}

func (cc *CustomersControllerImpl) CreateCustomers(ctx context.Context, req []service.CreateCustomersReq) (resp service.CreateCustomersResp, err error) {

	file, err := os.Open("csv_customers.csv")
	if err != nil {
		return service.CreateCustomersResp{
			Status: "failed",
		}, err
	}

	if err := gocsv.Unmarshal(file, &req); err != nil {
		return service.CreateCustomersResp{
			Status: "failed",
		}, err
	}

	return cc.customersUsecase.CreateCustomers(ctx, req)
}
