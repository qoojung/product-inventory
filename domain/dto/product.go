package dto

type Product struct {
	ID          uint   `json:"id"`
	SKU         string `json:"sku"`
	Name        string `json:"name"`
	Description string `json:"description"`
	Quantity    uint   `json:"quantity"`
	UnitPrice   uint   `json:"unit_price"`
}
type CreateProduct struct {
	SKU         string `json:"sku" binding:"required"`
	Name        string `json:"name" binding:"required"`
	Description string `json:"description"`
	Quantity    uint   `json:"quantity" binding:"required"`
	UnitPrice   uint   `json:"unit_price" binding:"required"`
}
type UpdateProduct struct {
	SKU         *string `json:"sku"`
	Name        *string `json:"name"`
	Description *string `json:"description"`
	Quantity    *uint   `json:"quantity"`
	UnitPrice   *uint   `json:"unit_price"`
}
type AdjustProductQuantity struct {
	Value int `json:"value" binding:"required"`
}
