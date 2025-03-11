package api_usecases

import (
	api_web_usecases "github.com/milfan/go-boilerplate/internal/api/usecases/web"
)

type (
	Usecases struct {
		WebUsecases api_web_usecases.WebUsecases
	}
)

func LoadUsecases() Usecases {
	return Usecases{
		WebUsecases: api_web_usecases.RegisterWebUsecases(),
	}
}
