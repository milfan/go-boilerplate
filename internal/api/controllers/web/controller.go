package api_web_controller

import (
	api_usecases "github.com/milfan/go-boilerplate/internal/api/usecases"
	pkg_response "github.com/milfan/go-boilerplate/pkg/response"
)

type (
	WebControllers struct {
		ProductController IProductController
	}
)

func RegisterWebController(
	pkgResponse pkg_response.IResponse,
	usecases api_usecases.Usecases,
) WebControllers {

	return WebControllers{
		ProductController: newProductController(pkgResponse, usecases.WebUsecases),
	}
}
