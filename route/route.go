package route

import (
	"app/ioc"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func Register(rootRoute *gin.RouterGroup, db *gorm.DB) {
	product := ioc.InitProductApp(db)
	userRouter := rootRoute.Group("/products")
	{
		userRouter.GET("/:id", product.GetProduct)
		userRouter.POST("/", product.CreateProduct)
		userRouter.GET("/", product.GetAllProducts)
		userRouter.PUT("/:id", product.UpdateProduct)
		userRouter.DELETE("/:id", product.DeleteProduct)
	}
}
