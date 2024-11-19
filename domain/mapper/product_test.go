package mapper

import (
	"app/domain/dao"
	"app/domain/dto"
	"reflect"
	"testing"
)

func TestToProductDAO(t *testing.T) {
	type args struct {
		productDto dto.Product
	}
	tests := []struct {
		name string
		args args
		want dao.Product
	}{
		{
			name: "toDAO",
			args: args{
				productDto: dto.Product{
					ID:          1,
					SKU:         "PROD",
					Name:        "Car",
					Description: "car",
					Quantity:    10,
					UnitPrice:   10000,
				},
			},
			want: dao.Product{
				ID:          1,
				SKU:         "PROD",
				Name:        "Car",
				Description: "car",
				Quantity:    10,
				UnitPrice:   10000,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := ToProductDAO(tt.args.productDto); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("ToProductDAO() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestCreateProductToProductDAO(t *testing.T) {
	type args struct {
		productDto dto.CreateProduct
	}
	tests := []struct {
		name string
		args args
		want dao.Product
	}{
		{
			name: "ToProductDAO",
			args: args{
				productDto: dto.CreateProduct{
					SKU:         "PROD",
					Name:        "Car",
					Description: "car",
					Quantity:    10,
					UnitPrice:   10000,
				},
			},
			want: dao.Product{
				SKU:         "PROD",
				Name:        "Car",
				Description: "car",
				Quantity:    10,
				UnitPrice:   10000,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := CreateProductToProductDAO(tt.args.productDto); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("CreateProductToProductDAO() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestUpdateProductToMap(t *testing.T) {
	sku := "PROD"
	name := "Car"
	description := "car"
	unitPrice := uint(10000)
	quantity := uint(10)
	type args struct {
		productDto dto.UpdateProduct
	}
	tests := []struct {
		name string
		args args
		want map[string]interface{}
	}{
		{
			name: "Full",
			args: args{
				productDto: dto.UpdateProduct{
					SKU:         &sku,
					Name:        &name,
					Description: &description,
					Quantity:    &quantity,
					UnitPrice:   &unitPrice,
				}},
			want: map[string]interface{}{
				"sku":         sku,
				"name":        name,
				"description": description,
				"quantity":    quantity,
				"unit_price":  unitPrice,
			},
		},
		{
			name: "Partial",
			args: args{
				productDto: dto.UpdateProduct{
					SKU: &sku,
				}},
			want: map[string]interface{}{
				"sku": sku,
			},
		},
		{
			name: "Partial",
			args: args{
				productDto: dto.UpdateProduct{
					Name: &name,
				}},
			want: map[string]interface{}{
				"name": name,
			},
		},
		{
			name: "Partial",
			args: args{
				productDto: dto.UpdateProduct{
					Description: &description,
				}},
			want: map[string]interface{}{
				"description": description,
			},
		},
		{
			name: "Partial",
			args: args{
				productDto: dto.UpdateProduct{
					Quantity: &quantity,
				}},
			want: map[string]interface{}{
				"quantity": quantity,
			},
		},
		{
			name: "Partial",
			args: args{
				productDto: dto.UpdateProduct{
					UnitPrice: &unitPrice,
				}},
			want: map[string]interface{}{
				"unit_price": unitPrice,
			},
		},
		{
			name: "Partial",
			args: args{
				productDto: dto.UpdateProduct{
					Name:        &name,
					Description: &description,
					Quantity:    &quantity,
					UnitPrice:   &unitPrice,
				}},
			want: map[string]interface{}{
				"name":        name,
				"description": description,
				"quantity":    quantity,
				"unit_price":  unitPrice,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := UpdateProductToMap(tt.args.productDto); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("UpdateProductToMap() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestToProductDTO(t *testing.T) {
	type args struct {
		productDao dao.Product
	}
	tests := []struct {
		name string
		args args
		want dto.Product
	}{
		{
			name: "toDAO",
			args: args{
				productDao: dao.Product{
					ID:          1,
					SKU:         "PROD",
					Name:        "Car",
					Description: "car",
					Quantity:    10,
					UnitPrice:   10000,
				},
			},
			want: dto.Product{
				ID:          1,
				SKU:         "PROD",
				Name:        "Car",
				Description: "car",
				Quantity:    10,
				UnitPrice:   10000,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := ToProductDTO(tt.args.productDao); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("ToProductDTO() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestToProductDTOs(t *testing.T) {
	type args struct {
		productDaos []dao.Product
	}
	tests := []struct {
		name string
		args args
		want []dto.Product
	}{
		{
			name: "ProductDAO",
			args: args{
				productDaos: []dao.Product{{
					ID:          1,
					SKU:         "PROD",
					Name:        "Car",
					Description: "car",
					Quantity:    10,
					UnitPrice:   10000,
				}},
			},
			want: []dto.Product{{
				ID:          1,
				SKU:         "PROD",
				Name:        "Car",
				Description: "car",
				Quantity:    10,
				UnitPrice:   10000,
			}},
		},
		{
			name: "ProductDAOEmpty",
			args: args{
				productDaos: []dao.Product{},
			},
			want: []dto.Product{},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := ToProductDTOs(tt.args.productDaos); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("ToProductDTOs() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestToProductDAOs(t *testing.T) {
	type args struct {
		productDTOs []dto.Product
	}
	tests := []struct {
		name string
		args args
		want []dao.Product
	}{
		{
			name: "ProductDAO",
			args: args{
				productDTOs: []dto.Product{{
					ID:          1,
					SKU:         "PROD",
					Name:        "Car",
					Description: "car",
					Quantity:    10,
					UnitPrice:   10000,
				}},
			},
			want: []dao.Product{{
				ID:          1,
				SKU:         "PROD",
				Name:        "Car",
				Description: "car",
				Quantity:    10,
				UnitPrice:   10000,
			}},
		},
		{
			name: "ProductDAOEmpty",
			args: args{
				productDTOs: []dto.Product{},
			},
			want: []dao.Product{},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := ToProductDAOs(tt.args.productDTOs); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("ToProductDAOs() = %v, want %v", got, tt.want)
			}
		})
	}
}
