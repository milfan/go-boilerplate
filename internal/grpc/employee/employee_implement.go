package grpc_employee

import (
	"context"

	pkg_grpc_employee "github.com/milfan/go-boilerplate/pkg/grpc/employee"
	"github.com/sirupsen/logrus"
)

type EmployeeService struct {
	pkg_grpc_employee.UnimplementedEmployeeGrpcServer
	logger *logrus.Logger
}

func (s *EmployeeService) DetailEmployee(
	ctx context.Context,
	req *pkg_grpc_employee.DetailEmployeeRequest,
) (*pkg_grpc_employee.DetailEmployeeResponse, error) {

	return &pkg_grpc_employee.DetailEmployeeResponse{
		EmpCode: "empCode",
		EmpName: "empName",
	}, nil
}
func New(
	logger *logrus.Logger,
) *EmployeeService {
	return &EmployeeService{
		logger: logger,
	}
}
