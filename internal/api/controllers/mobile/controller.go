package api_mobile_controller

import (
	api_usecases "github.com/milfan/go-boilerplate/internal/api/usecases"
	pkg_response "github.com/milfan/go-boilerplate/pkg/response"
)

type (
	MobileControllers struct {
		OrderController IOrderController
	}
)

func RegisterMobileController(
	pkgResponse pkg_response.IResponse,
	usecases api_usecases.Usecases,
) MobileControllers {

	return MobileControllers{
		OrderController: newOrderController(pkgResponse, usecases.MobileUsecases),
	}
}
