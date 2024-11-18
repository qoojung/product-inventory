package repository

import (
	"app/domain/dao"

	"gorm.io/gorm"
)

type ProductRepository interface {
	FindById(id uint64) (dao.Product, error)
	FindAll() ([]dao.Product, error)
	Save(product dao.Product) (dao.Product, error)
	UpdateById(id uint64, productMap map[string]interface{}) (int64, error)
	Delete(id uint64) (int64, error)
}
type ProductRepositoryImpl struct {
	db *gorm.DB
}

func (p ProductRepositoryImpl) FindById(id uint64) (dao.Product, error) {
	product := dao.Product{}
	err := p.db.First(&product, "id = ?", id).Error
	return product, err
}
func (p ProductRepositoryImpl) FindAll() ([]dao.Product, error) {
	products := []dao.Product{}
	err := p.db.Find(&products).Error
	return products, err
}
func (p ProductRepositoryImpl) Save(product dao.Product) (dao.Product, error) {
	err := p.db.Create(&product).Error
	return product, err
}
func (p ProductRepositoryImpl) UpdateById(id uint64, productMap map[string]interface{}) (int64, error) {
	result := p.db.Model(&dao.Product{}).Where("id = ?", id).Updates(productMap)
	return result.RowsAffected, result.Error
}

func (p ProductRepositoryImpl) Delete(id uint64) (int64, error) {
	result := p.db.Delete(&dao.Product{}, id)
	return result.RowsAffected, result.Error
}
func NewProductRepository(db *gorm.DB) ProductRepository {
	return &ProductRepositoryImpl{db: db}
}
