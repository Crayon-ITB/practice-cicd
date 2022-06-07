package models

type Employee struct {
	Id   uint32 `json:"id,omitempty"`
	Name string `json:"name"`
}
