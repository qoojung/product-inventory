package mapper

import (
	"app/domain/dao"
	"app/domain/dto"
)

func ToProductDAO(productDto dto.Product) dao.Product {
	return dao.Product{
		ID:          productDto.ID,
		SKU:         productDto.SKU,
		Name:        productDto.Name,
		Description: productDto.Description,
		Quantity:    productDto.Quantity,
		UnitPrice:   productDto.UnitPrice,
	}
}

func CreateProductToProductDAO(productDto dto.CreateProduct) dao.Product {
	return dao.Product{
		SKU:         productDto.SKU,
		Name:        productDto.Name,
		Description: productDto.Description,
		Quantity:    productDto.Quantity,
		UnitPrice:   productDto.UnitPrice,
	}
}
func UpdateProductToMap(productDto dto.UpdateProduct) map[string]interface{} {
	return map[string]interface{}{
		"sku":         productDto.SKU,
		"name":        productDto.Name,
		"description": productDto.Description,
		"quantity":    productDto.Quantity,
		"unit_price":  productDto.UnitPrice,
	}
}

func ToProductDTO(productDao dao.Product) dto.Product {
	return dto.Product{
		ID:          productDao.ID,
		SKU:         productDao.SKU,
		Name:        productDao.Name,
		Description: productDao.Description,
		Quantity:    productDao.Quantity,
		UnitPrice:   productDao.UnitPrice,
	}
}

func ToProductDTOs(productDaos []dao.Product) []dto.Product {
	productDTOs := make([]dto.Product, len(productDaos))

	for i, item := range productDaos {
		productDTOs[i] = ToProductDTO(item)
	}

	return productDTOs
}

func ToProductDAOs(productDTOs []dto.Product) []dao.Product {
	productDaos := make([]dao.Product, len(productDTOs))

	for i, item := range productDTOs {
		productDaos[i] = ToProductDAO(item)
	}

	return productDaos
}
