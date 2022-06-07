package rest

import (
	"log"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/go-playground-2/internals/repositories"
	"github.com/go-playground-2/internals/services"
)

func Run() {
	db := repositories.InitDB()
	defer db.Close()

	er := repositories.NewEmployeeRepo()
	ar := repositories.NewAttendanceRepo()
	es := services.NewEmployeeServices(er)
	as := services.NewAttendanceServices(ar)

	handler := httpHandler{db: db, es: es, as: as}
	router := chi.NewRouter()

	router.Post("/employee", handler.registerEmployee)
	router.Get("/employee", handler.getAllEmployees)
	router.Get("/employee/{id}", handler.getEmployeeById)
	router.Patch("/employee/{id}", handler.updateEmployee)
	router.Delete("/employee/{id}", handler.deleteEmployee)

	router.Post("/attendance/attend", handler.attend)
	router.Post("/attendance/leave", handler.leave)
	router.Get("/attendance", handler.getAllAttendances)

	log.Println("listening 8080")
	http.ListenAndServe("localhost:8080", router)
}
