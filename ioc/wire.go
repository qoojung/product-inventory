//go:build wireinject
// +build wireinject

package ioc

import (
	"app/controller"
	"app/repository"
	"app/service"

	"github.com/google/wire"
	"gorm.io/gorm"
)

func InitProductApp(db *gorm.DB) controller.ProductController {
	wire.Build(repository.NewProductRepository, service.NewProductService, controller.NewProductController)
	return controller.ProductControllerImpl{}
}
