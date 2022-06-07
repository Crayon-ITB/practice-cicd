package repositories

import (
	"context"
	"database/sql"
	"errors"
	"log"

	"github.com/go-playground-2/internals/models"
)

type AttendanceRepo interface {
	GetAll(ctx context.Context) ([]models.Attendance, error)
	Create(ctx context.Context, idEmp uint32, date string, attends bool) (models.Attendance, error)
}

type PostgresAttendance struct{}

func (pa *PostgresAttendance) GetAll(ctx context.Context) ([]models.Attendance, error) {
	if db, ok := ctx.Value("db").(*sql.DB); ok {
		rows, err := db.Query("SELECT id, emp_id, date, attends FROM attendances")
		if err != nil {
			return []models.Attendance{}, err
		}
		defer rows.Close()

		var allAttendances []models.Attendance
		for rows.Next() {
			var attendance models.Attendance
			rows.Scan(&attendance.Id, &attendance.IdEmp, &attendance.Date, &attendance.Attends)
			allAttendances = append(allAttendances, attendance)
		}
		return allAttendances, nil
	}
	return []models.Attendance{}, errors.New("Database retrieval error")
}

func (pa *PostgresAttendance) Create(ctx context.Context, idEmp uint32, date string, attends bool) (models.Attendance, error) {
	if db, ok := ctx.Value("db").(*sql.DB); ok {
		log.Println(idEmp, date, attends)
		querystring := "INSERT INTO attendances (emp_id, date, attends) VALUES ($1, $2, $3) RETURNING id, emp_id, date, attends"
		var newAttendance models.Attendance
		if err := db.QueryRow(querystring, idEmp, date, attends).Scan(&newAttendance.Id, &newAttendance.IdEmp, &newAttendance.Date, &newAttendance.Attends); err != nil && err != sql.ErrNoRows {
			return models.Attendance{}, err
		}
		return newAttendance, nil
	}
	return models.Attendance{}, errors.New("Database retrieval error")
}

func NewAttendanceRepo() AttendanceRepo {
	return &PostgresAttendance{}
}
