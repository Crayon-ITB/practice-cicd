package services

import (
	"context"
	"encoding/json"

	"github.com/go-playground-2/internals/repositories"
)

type EmployeeServices interface { // Q: naming?
	Create(ctx context.Context, name string) ([]byte, error) // Q: mending return struct apa json?
	GetAll(ctx context.Context) ([]byte, error)
	GetById(ctx context.Context, id uint32) ([]byte, error)
	Update(ctx context.Context, id uint32, name string) ([]byte, error)
	Delete(ctx context.Context, id uint32) error
}

type EmployeeServicesStruct struct {
	er repositories.EmployeeRepo // Q: mending pointer apa nggak?
}

func (es *EmployeeServicesStruct) Create(ctx context.Context, name string) ([]byte, error) {
	newEmployee, err := es.er.Create(ctx, name)
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(newEmployee)
}

func (es *EmployeeServicesStruct) GetAll(ctx context.Context) ([]byte, error) {
	employees, err := es.er.GetAll(ctx)
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(employees)
}

func (es *EmployeeServicesStruct) GetById(ctx context.Context, id uint32) ([]byte, error) {
	employee, err := es.er.GetById(ctx, id)
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(employee)
}

func (es *EmployeeServicesStruct) Update(ctx context.Context, id uint32, name string) ([]byte, error) {
	employee, err := es.er.Update(ctx, id, name)
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(employee)
}

func (es *EmployeeServicesStruct) Delete(ctx context.Context, id uint32) error {
	err := es.er.Delete(ctx, id)
	if err != nil {
		return err
	}
	return nil
}

func NewEmployeeServices(er repositories.EmployeeRepo) EmployeeServices {
	return &EmployeeServicesStruct{er: er}
}
