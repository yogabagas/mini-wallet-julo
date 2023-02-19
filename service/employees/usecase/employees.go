package usecase

import (
	"context"

	"github.com/yogabagas/jatis-BE/domain/model"
	"github.com/yogabagas/jatis-BE/domain/repository"
	"github.com/yogabagas/jatis-BE/domain/service"
)

type EmployeessUsecaseImpl struct {
	employeesRepo repository.EmployeesRepository
}

type EmployeesUsecase interface {
	CreateEmployees(ctx context.Context, req []service.CreateEmployeesReq) (resp service.CreateEmployeesResp, err error)
}

func NewEmployeesUsecase(employeesRepo repository.EmployeesRepository) EmployeesUsecase {
	return &EmployeessUsecaseImpl{employeesRepo: employeesRepo}
}

func (eu *EmployeessUsecaseImpl) CreateEmployees(ctx context.Context, req []service.CreateEmployeesReq) (resp service.CreateEmployeesResp, err error) {

	var reqEmployees []model.Employee

	for _, v := range req {
		reqEmployee := model.Employee{
			EmployeeId: v.EmployeeId,
			FirstName:  v.FirstName,
			LastName:   v.LastName,
			Title:      v.Title,
			WorkPhone:  v.WorkPhone,
			CreatedBy:  v.CreatedBy,
			UpdatedBy:  v.UpdatedBy,
		}
		reqEmployees = append(reqEmployees, reqEmployee)
	}

	err = eu.employeesRepo.InsertEmployees(ctx, reqEmployees)
	if err != nil {
		return service.CreateEmployeesResp{
			Status: "failed",
		}, err
	}

	return service.CreateEmployeesResp{
		Status: "created",
	}, nil
}
