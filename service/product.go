package service

import (
	"app/domain/dto"
	"app/domain/mapper"
	"app/repository"
	"app/util"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type ProductService interface {
	GetProduct(ctx *gin.Context)
	GetAllProducts(ctx *gin.Context)
	CreateProduct(ctx *gin.Context)
	DeleteProduct(ctx *gin.Context)
	UpdateProduct(ctx *gin.Context)
}

type ProductServiceImpl struct {
	repo repository.ProductRepository
}

func (p ProductServiceImpl) GetAllProducts(ctx *gin.Context) {
	products, err := p.repo.FindAll()
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"message": "products not found",
		})
		return
	}
	resp := util.BuildSuccessResponse_(mapper.ToProductDTOs(products))
	ctx.JSON(http.StatusOK, resp)
}
func (p ProductServiceImpl) GetProduct(ctx *gin.Context) {
	idStr := ctx.Param("id")
	id, err := strconv.ParseUint(idStr, 10, 64)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"message": "id not found",
		})
		return
	}
	daoProduct, err := p.repo.FindById(id)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"message": "product not found",
		})
		return
	}
	resp := util.BuildSuccessResponse_(mapper.ToProductDTO(daoProduct))
	ctx.JSON(200, resp)
}
func (p ProductServiceImpl) CreateProduct(ctx *gin.Context) {
	productReq := dto.CreateProduct{}
	bindErr := ctx.ShouldBindJSON(&productReq)
	if bindErr != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"message": bindErr.Error(),
		})
		return
	}
	_, err := p.repo.Save(mapper.CreateProductToProductDAO(productReq))
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"message": "product not created",
		})
		return
	}
	ctx.Status(http.StatusCreated)
}
func (p ProductServiceImpl) DeleteProduct(ctx *gin.Context) {
	id, err := strconv.ParseUint(ctx.Param("id"), 10, 64)
	if err != nil {
		ctx.JSON(400, gin.H{
			"message": "id not found",
		})
		return
	}
	row, err := p.repo.Delete(id)
	if row == 0 || err != nil {
		ctx.JSON(400, gin.H{
			"message": "product not deleted",
		})
		return
	}
	ctx.JSON(200, util.BuildEmptySuccessResponse_())
}
func (p ProductServiceImpl) UpdateProduct(ctx *gin.Context) {
	updateReq := dto.UpdateProduct{}
	err := ctx.ShouldBindJSON(&updateReq)
	if err != nil {
		ctx.JSON(400, gin.H{
			"message": err.Error(),
		})
		return
	}
	id, err := strconv.ParseUint(ctx.Param("id"), 10, 64)
	if err != nil {
		ctx.JSON(400, gin.H{
			"message": "id not found",
		})
		return
	}
	rows, err := p.repo.UpdateById(id, mapper.UpdateProductToMap(updateReq))
	if err != nil || rows == 0 {
		ctx.JSON(400, gin.H{
			"message": "product not updated",
		})
		return
	}
	ctx.JSON(200, util.BuildEmptySuccessResponse_())
}
func NewProductService(repo repository.ProductRepository) ProductService {
	return &ProductServiceImpl{
		repo: repo,
	}
}
