package services

import (
	"context"
	"encoding/json"

	"github.com/go-playground-2/internals/repositories"
)

type AttendanceServices interface { // Q: naming?
	GetAll(ctx context.Context) ([]byte, error) // Q: mending return json apa struct?
	Attend(ctx context.Context, idEmp uint32, date string) ([]byte, error)
	Leave(ctx context.Context, idEmp uint32, date string) ([]byte, error)
}

type AttendanceServicesStruct struct {
	ar repositories.AttendanceRepo
}

func (as *AttendanceServicesStruct) GetAll(ctx context.Context) ([]byte, error) {
	attendances, err := as.ar.GetAll(ctx)
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(attendances)
}

func (as *AttendanceServicesStruct) Attend(ctx context.Context, idEmp uint32, date string) ([]byte, error) {
	attendance, err := as.ar.Create(ctx, idEmp, date, true)
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(attendance)
}

func (as *AttendanceServicesStruct) Leave(ctx context.Context, idEmp uint32, date string) ([]byte, error) {
	attendance, err := as.ar.Create(ctx, idEmp, date, false)
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(attendance)
}

func NewAttendanceServices(ar repositories.AttendanceRepo) AttendanceServices {
	return &AttendanceServicesStruct{ar: ar}
}
