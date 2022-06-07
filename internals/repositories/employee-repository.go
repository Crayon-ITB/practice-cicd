package repositories

import (
	"context"
	"database/sql"
	"errors"

	"github.com/go-playground-2/internals/models"
)

type EmployeeRepo interface {
	Create(ctx context.Context, name string) (models.Employee, error) // Q: mending return struct apa json?
	GetAll(ctx context.Context) ([]models.Employee, error)
	GetById(ctx context.Context, id uint32) (models.Employee, error)
	Update(ctx context.Context, id uint32, name string) (models.Employee, error)
	Delete(ctx context.Context, id uint32) error
}

type PostgresEmployee struct{}

func (pe *PostgresEmployee) Create(ctx context.Context, name string) (models.Employee, error) {
	if db, ok := ctx.Value("db").(*sql.DB); ok {
		var newEmployee models.Employee
		if err := db.QueryRow("INSERT INTO employees (name) VALUES ($1) RETURNING id, name", name).Scan(&newEmployee.Id, &newEmployee.Name); err != nil && err != sql.ErrNoRows {
			return models.Employee{}, err
		}
		return newEmployee, nil
	}
	return models.Employee{}, errors.New("Database retrieval error")
}

func (pe *PostgresEmployee) GetAll(ctx context.Context) ([]models.Employee, error) {
	if db, ok := ctx.Value("db").(*sql.DB); ok {
		rows, err := db.Query("SELECT id, name FROM employees")
		if err != nil {
			return []models.Employee{}, err
		}
		defer rows.Close()

		var allEmployees []models.Employee // Q: read rows repetitive, kalo mau modular gimana?
		for rows.Next() {
			var employee models.Employee
			rows.Scan(&employee.Id, &employee.Name)
			allEmployees = append(allEmployees, employee)
		}
		return allEmployees, nil
	}
	return []models.Employee{}, errors.New("Database retrieval error")
}

func (pe *PostgresEmployee) GetById(ctx context.Context, id uint32) (models.Employee, error) {
	if db, ok := ctx.Value("db").(*sql.DB); ok {
		var employee models.Employee
		if err := db.QueryRow("SELECT id, name FROM employees WHERE id = $1", id).Scan(&employee.Id, &employee.Name); err != nil && err != sql.ErrNoRows {
			return models.Employee{}, err
		}
		return employee, nil
	}
	return models.Employee{}, errors.New("Database retrieval error")
}

func (pe *PostgresEmployee) Update(ctx context.Context, id uint32, name string) (models.Employee, error) {
	if db, ok := ctx.Value("db").(*sql.DB); ok {
		var employee models.Employee
		if err := db.QueryRow("UPDATE employees SET name = $1 WHERE id = $2 RETURNING id, name", name, id).Scan(&employee.Id, &employee.Name); err != nil && err != sql.ErrNoRows {
			return models.Employee{}, err
		}
		return employee, nil
	}
	return models.Employee{}, errors.New("Database retrieval error")
}

func (pe *PostgresEmployee) Delete(ctx context.Context, id uint32) error {
	if db, ok := ctx.Value("db").(*sql.DB); ok {
		_, err := db.Exec("DELETE FROM employees WHERE id = $1", id)
		if err != nil {
			return err
		}
		return nil
	}
	return errors.New("Database retrieval error")
}

func NewEmployeeRepo() EmployeeRepo {
	return &PostgresEmployee{}
}
