package models

// import "time"

type Attendance struct {
	Id      uint32 `json:"id"`
	IdEmp   uint32 `json:"emp_id"`
	Date    string `json:"date"`
	Attends bool   `json:"attends"`
}
