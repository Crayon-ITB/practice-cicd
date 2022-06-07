package rest

import (
	"context"
	"database/sql"
	"encoding/json"
	"log"
	"net/http"
	"strconv"

	"github.com/go-chi/chi/v5"
	"github.com/go-playground-2/internals/services"
)

type httpHandler struct { // biar bisa masukin db ke handler?
	db *sql.DB
	es services.EmployeeServices
	as services.AttendanceServices
}

type EmployeeRegister struct {
	Name string `json:name`
}

type EmployeeUpdate struct {
	Name string `json:name`
}

func (h *httpHandler) registerEmployee(w http.ResponseWriter, r *http.Request) {
	ctx := context.WithValue(r.Context(), "db", h.db)
	var data EmployeeRegister
	err := json.NewDecoder(r.Body).Decode(&data)
	if err != nil {
		log.Println(err)
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	employee, err := h.es.Create(ctx, data.Name)
	if err != nil {
		log.Println(err)
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	w.Header().Set("content-type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(employee)
}

func (h *httpHandler) getAllEmployees(w http.ResponseWriter, r *http.Request) {
	ctx := context.WithValue(r.Context(), "db", h.db)
	employees, err := h.es.GetAll(ctx)
	if err != nil {
		log.Println(err)
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	w.Header().Set("content-type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(employees)
}

func (h *httpHandler) getEmployeeById(w http.ResponseWriter, r *http.Request) {
	ctx := context.WithValue(r.Context(), "db", h.db)
	id, err := strconv.Atoi(chi.URLParam(r, "id"))
	if err != nil {
		log.Println(err)
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	employee, err := h.es.GetById(ctx, uint32(id))
	if err != nil {
		log.Println(err)
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	w.Header().Set("content-type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(employee)
}

func (h *httpHandler) updateEmployee(w http.ResponseWriter, r *http.Request) {
	ctx := context.WithValue(r.Context(), "db", h.db)
	id, err := strconv.Atoi(chi.URLParam(r, "id"))
	if err != nil {
		log.Println(err)
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	var data EmployeeUpdate
	json.NewDecoder(r.Body).Decode(&data)
	employee, err := h.es.Update(ctx, uint32(id), data.Name)
	if err != nil {
		log.Println(err)
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	w.Header().Set("content-type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(employee)
}

func (h *httpHandler) deleteEmployee(w http.ResponseWriter, r *http.Request) {
	ctx := context.WithValue(r.Context(), "db", h.db)
	id, err := strconv.Atoi(chi.URLParam(r, "id"))
	if err != nil {
		log.Println(err)
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	err = h.es.Delete(ctx, uint32(id))
	if err != nil {
		log.Println(err)
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	w.WriteHeader(http.StatusOK)
}

type Attendance struct {
	idEmp uint32 `json:emp_id`
	date  string `json:date`
}

func (h *httpHandler) attend(w http.ResponseWriter, r *http.Request) {
	ctx := context.WithValue(r.Context(), "db", h.db)
	var attendance Attendance
	err := json.NewDecoder(r.Body).Decode(&attendance) // lol kok gamau decode
	if err != nil {
		log.Println(err)
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	log.Println("body: ", attendance)
	res, err := h.as.Attend(ctx, attendance.idEmp, attendance.date)
	if err != nil {
		log.Println(err)
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	w.Header().Set("content-type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

func (h *httpHandler) leave(w http.ResponseWriter, r *http.Request) {
	ctx := context.WithValue(r.Context(), "db", h.db)
	var attendance Attendance
	err := json.NewDecoder(r.Body).Decode(&attendance)
	if err != nil {
		log.Println(err)
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	res, err := h.as.Leave(ctx, attendance.idEmp, attendance.date)
	if err != nil {
		log.Println(err)
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	w.Header().Set("content-type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

func (h *httpHandler) getAllAttendances(w http.ResponseWriter, r *http.Request) {
	ctx := context.WithValue(r.Context(), "db", h.db)
	attendances, err := h.as.GetAll(ctx)
	if err != nil {
		log.Println(err)
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	w.Header().Set("content-type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(attendances)
}
