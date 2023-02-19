package usecase

import (
	"context"

	"github.com/yogabagas/jatis-BE/domain/model"
	"github.com/yogabagas/jatis-BE/domain/repository"
	"github.com/yogabagas/jatis-BE/domain/service"
)

type CustomersUsecaseImpl struct {
	cutomersRepo repository.CustomersRepository
}

type CustomersUsecase interface {
	CreateCustomers(ctx context.Context, req []service.CreateCustomersReq) (resp service.CreateCustomersResp, err error)
}

func NewCustomersUsecase(customersRepo repository.CustomersRepository) CustomersUsecase {
	return &CustomersUsecaseImpl{cutomersRepo: customersRepo}
}

func (cu *CustomersUsecaseImpl) CreateCustomers(ctx context.Context, req []service.CreateCustomersReq) (resp service.CreateCustomersResp, err error) {

	var reqCustomers []model.Customer

	for _, v := range req {
		reqCustomer := model.Customer{
			CustomerId:      v.CustomerId,
			CompanyName:     v.CompanyName,
			FirstName:       v.FirstName,
			LastName:        v.LastName,
			BillingAddress:  v.BillingAddress,
			City:            v.City,
			Province:        v.Province,
			ZipCode:         v.ZipCode,
			Email:           v.Email,
			CompanyWebsite:  v.CompanyWebsite,
			PhoneNumber:     v.PhoneNumber,
			FaxNumber:       v.FaxNumber,
			ShipAddress:     v.ShipAddress,
			ShipCity:        v.ShipCity,
			ShipProvince:    v.ShipProvince,
			ShipZipCode:     v.ShipZipCode,
			ShipPhoneNumber: v.ShipPhoneNumber,
			CreatedBy:       v.CreatedBy,
			UpdatedBy:       v.UpdatedBy,
		}
		reqCustomers = append(reqCustomers, reqCustomer)
	}

	err = cu.cutomersRepo.InsertCustomers(ctx, reqCustomers)
	if err != nil {
		return service.CreateCustomersResp{
			Status: "failed",
		}, err
	}

	return service.CreateCustomersResp{
		Status: "created",
	}, nil
}
