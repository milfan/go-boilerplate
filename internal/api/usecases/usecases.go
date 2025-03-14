package api_usecases

import (
	"github.com/milfan/go-boilerplate/internal/api/repositories"
	api_web_usecases "github.com/milfan/go-boilerplate/internal/api/usecases/web"
)

type (
	Usecases struct {
		WebUsecases api_web_usecases.WebUsecases
	}
)

func LoadUsecases(
	repo repositories.Repositories,
) Usecases {
	return Usecases{
		WebUsecases: api_web_usecases.RegisterWebUsecases(repo),
	}
}
