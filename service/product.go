package service

import (
	"app/domain/dto"
	"app/domain/mapper"
	"app/repository"
	"app/util"
	"errors"

	log "github.com/sirupsen/logrus"
	"gorm.io/gorm"
)

type ProductService interface {
	GetAllProducts() ([]dto.Product, error)
	GetProduct(id uint64) (dto.Product, error)
	CreateProduct(createObj dto.CreateProduct) (dto.Product, error)
	DeleteProduct(id uint64) error
	UpdateProduct(id uint64, updateObj dto.UpdateProduct) error
}

type ProductServiceImpl struct {
	repo repository.ProductRepository
}

func (p ProductServiceImpl) GetAllProducts() ([]dto.Product, error) {
	products, err := p.repo.FindAll()
	if err != nil {
		return []dto.Product{}, err
	}
	return mapper.ToProductDTOs(products), nil

}
func (p ProductServiceImpl) GetProduct(id uint64) (dto.Product, error) {
	daoProduct, err := p.repo.FindById(id)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return dto.Product{}, &util.ApiError{Status: util.NotFound}
		}
		return dto.Product{}, err
	}
	return mapper.ToProductDTO(daoProduct), nil
}
func (p ProductServiceImpl) CreateProduct(createObj dto.CreateProduct) (dto.Product, error) {
	productDAO, err := p.repo.Save(mapper.CreateProductToProductDAO(createObj))
	if err != nil {
		return dto.Product{}, err
	}
	return mapper.ToProductDTO(productDAO), nil
}
func (p ProductServiceImpl) DeleteProduct(id uint64) error {
	row, err := p.repo.Delete(id)
	if err != nil {
		log.Error(err)
		return err
	}
	if row == 0 {
		return &util.ApiError{Status: util.NotFound}
	}
	return nil
}
func (p ProductServiceImpl) UpdateProduct(id uint64, updateObj dto.UpdateProduct) error {
	rows, err := p.repo.UpdateById(id, mapper.UpdateProductToMap(updateObj))
	if err != nil {
		log.Error(err)
		return err
	}
	if rows == 0 {
		return &util.ApiError{Status: util.NotFound}
	}
	return nil
}
func NewProductService(repo repository.ProductRepository) ProductService {
	return &ProductServiceImpl{
		repo: repo,
	}
}
