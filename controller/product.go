package controller

import (
	"app/domain/dto"
	"app/service"
	"app/util"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	log "github.com/sirupsen/logrus"
)

type ProductController interface {
	GetProduct(ctx *gin.Context)
	GetAllProducts(ctx *gin.Context)
	CreateProduct(ctx *gin.Context)
	DeleteProduct(ctx *gin.Context)
	UpdateProduct(ctx *gin.Context)
	AdjustProductQuantity(ctx *gin.Context)
}

type ProductControllerImpl struct {
	svc service.ProductService
}

func (p ProductControllerImpl) GetAllProducts(ctx *gin.Context) {
	products, err := p.svc.GetAllProducts()
	if err != nil {
		ctx.JSON(http.StatusBadRequest, util.BuildErrorResponseFromError(err))
		log.Error(err)
		return
	}
	resp := util.BuildSuccessResponse(products)
	ctx.JSON(http.StatusOK, resp)
}
func (p ProductControllerImpl) GetProduct(ctx *gin.Context) {
	idStr := ctx.Param("id")
	id, err := strconv.ParseUint(idStr, 10, 64)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, util.BuildErrorResponse(util.Argument))
		return
	}
	product, err := p.svc.GetProduct(id)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, util.BuildErrorResponseFromError(err))
		return
	}
	resp := util.BuildSuccessResponse(product)
	ctx.JSON(http.StatusOK, resp)
}
func (p ProductControllerImpl) CreateProduct(ctx *gin.Context) {
	productReq := dto.CreateProduct{}
	err := ctx.ShouldBindJSON(&productReq)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, util.BuildErrorResponse(util.Argument))
		log.Info(err)
		return
	}
	product, err := p.svc.CreateProduct(productReq)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, util.BuildErrorResponseFromError(err))
		log.Error(err)
	}
	ctx.JSON(http.StatusCreated, util.BuildSuccessResponse(product))
}
func (p ProductControllerImpl) DeleteProduct(ctx *gin.Context) {
	id, err := strconv.ParseUint(ctx.Param("id"), 10, 64)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, util.BuildErrorResponse(util.Argument))
		return
	}
	err = p.svc.DeleteProduct(id)
	if err != nil {
		log.Error(err)
		ctx.JSON(http.StatusBadRequest, util.BuildErrorResponseFromError(err))
		return
	}
	ctx.JSON(http.StatusOK, util.BuildEmptySuccessResponse())
}
func (p ProductControllerImpl) UpdateProduct(ctx *gin.Context) {
	updateReq := dto.UpdateProduct{}
	err := ctx.ShouldBindJSON(&updateReq)
	if err != nil {
		log.Error(err)
		ctx.JSON(http.StatusBadRequest, util.BuildErrorResponse(util.Argument))
		return
	}
	id, err := strconv.ParseUint(ctx.Param("id"), 10, 64)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, util.BuildErrorResponse(util.Argument))
		return
	}

	err = p.svc.UpdateProduct(id, updateReq)
	if err != nil {

		log.Error(err)

		ctx.JSON(http.StatusBadRequest, util.BuildErrorResponseFromError(err))
		return
	}
	ctx.JSON(http.StatusOK, util.BuildEmptySuccessResponse())
}
func (p ProductControllerImpl) AdjustProductQuantity(ctx *gin.Context) {
	adjustReq := dto.AdjustProductQuantity{}
	err := ctx.ShouldBindJSON(&adjustReq)
	if err != nil {
		log.Error(err)
		ctx.JSON(http.StatusBadRequest, util.BuildErrorResponse(util.Argument))
		return
	}
	id, err := strconv.ParseUint(ctx.Param("id"), 10, 64)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, util.BuildErrorResponse(util.Argument))
		return
	}
	err = p.svc.IncrementProductQuantity(id, adjustReq.Value)
	if err != nil {
		log.Error(err)
		ctx.JSON(http.StatusBadRequest, util.BuildErrorResponseFromError(err))
		return
	}
	ctx.JSON(http.StatusOK, util.BuildEmptySuccessResponse())
}
func NewProductController(svc service.ProductService) ProductController {
	return ProductControllerImpl{svc: svc}
}
