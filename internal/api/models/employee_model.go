package models

import (
	"time"

	"github.com/milfan/go-boilerplate/internal/api/entities"
	"gorm.io/gorm"
)

type Employee struct {
	ID        uint64         `gorm:"column:id;primaryKey"`
	EmpCode   string         `gorm:"column:emp_code"`
	EmpName   string         `gorm:"column:emp_name"`
	CreatedBy string         `gorm:"column:created_by"`
	CreatedAt time.Time      `gorm:"column:created_at"`
	UpdatedBy string         `gorm:"column:updated_by"`
	UpdatedAt time.Time      `gorm:"column:updated_at"`
	DeletedAt gorm.DeletedAt `gorm:"column:deleted_at"`
}

func (m Employee) Entity() *entities.Employee {
	return entities.EmployeeEntity(
		m.ID,
		m.EmpCode,
		m.EmpName,
		m.CreatedBy,
		m.CreatedAt,
		m.UpdatedBy,
		m.UpdatedAt,
	)
}

func ToEmployeeModel(e entities.Employee) *Employee {
	return &Employee{
		ID:        e.ID(),
		EmpCode:   e.EmpCode(),
		EmpName:   e.EmpName(),
		CreatedBy: e.CreatedBy(),
		CreatedAt: e.CreatedAt(),
		UpdatedBy: e.UpdatedBy(),
		UpdatedAt: e.UpdatedAt(),
	}
}
