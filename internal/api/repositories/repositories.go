package repositories

import (
	config_postgres "github.com/milfan/go-boilerplate/configs/postgres"
	"github.com/sirupsen/logrus"
)

type (
	Repositories struct {
		ProductRepositories IProductRepository
		OrderRepositories   IOrdersRepository
	}
)

func LoadRepositories(
	conn config_postgres.Postgres,
	logger *logrus.Logger,
) Repositories {
	return Repositories{
		ProductRepositories: newProductsRepository(conn, logger),
		OrderRepositories:   newOrdersRepository(conn, logger),
	}
}
