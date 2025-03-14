package entities

import (
	"strconv"
	"time"
)

type Employee struct {
	BaseEntity
	id      uint64
	empCode string
	empName string
}

func (e *Employee) ID() uint64 {
	return e.id
}

func (e *Employee) EmpCode() string {
	return e.empCode
}

func (e *Employee) EmpName() string {
	return e.empName
}

func GenerateEmpCode(lastId uint64) string {
	code := "EMP-"
	year, month, _ := time.Now().Date()

	s := strconv.FormatUint(uint64(year+99), 10)
	code = code + s
	u := strconv.FormatUint(uint64(month+12), 10)
	code = code + u
	o := strconv.FormatUint(lastId, 10)
	code = code + o
	return code
}

func EmployeeEntity(
	id uint64,
	empCode, empName string,
	createdBy string,
	createdAt time.Time,
	updatedBy string,
	updatedAt time.Time,
) *Employee {
	return &Employee{
		BaseEntity: BaseEntity{
			createdBy: createdBy,
			createdAt: createdAt,
			updatedBy: updatedBy,
			updatedAt: updatedAt,
		},
		id:      id,
		empCode: empCode,
		empName: empName,
	}
}

func NewEmployee(
	empName string,
	createdBy string,
) *Employee {
	return &Employee{
		BaseEntity: New(createdBy),
		empName:    empName,
	}
}
