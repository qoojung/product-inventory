//go:build wireinject
// +build wireinject

package ioc

import (
	"app/repository"
	"app/service"

	"github.com/google/wire"
	"gorm.io/gorm"
)

func InitProductService(db *gorm.DB) service.ProductService {
	wire.Build(repository.NewProductRepository, service.NewProductService)
	return service.ProductServiceImpl{}
}
