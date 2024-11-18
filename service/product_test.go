package service

import (
	"app/domain/dao"
	"app/domain/dto"
	"app/repository"
	"reflect"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	"go.uber.org/mock/gomock"
)

func TestProductServiceImpl_GetAllProducts(t *testing.T) {
	mockctl := gomock.NewController(t)
	defer mockctl.Finish()

	tests := []struct {
		name         string
		functionCall func(mock *repository.MockProductRepository)
		want         []dto.Product
		expectError  error
	}{
		{
			name: "List Product",
			functionCall: func(mock *repository.MockProductRepository) {
				mock.EXPECT().FindAll().Return([]dao.Product{
					{
						ID:          1,
						SKU:         "PROD",
						Name:        "chair",
						Description: "stylish Chair",
						Quantity:    3000,
						UnitPrice:   2000,
						CreatedAt:   time.Now(),
						UpdatedAt:   time.Now(),
					}}, nil)
			},
			want: []dto.Product{{
				ID:          1,
				SKU:         "PROD",
				Name:        "chair",
				Description: "stylish Chair",
				Quantity:    3000,
				UnitPrice:   2000,
			}},
		},
		{
			name: "List Empty Product",
			functionCall: func(mock *repository.MockProductRepository) {
				mock.EXPECT().FindAll().Return([]dao.Product{}, nil)
			},
			want: []dto.Product{},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			mockRepo := repository.NewMockProductRepository(mockctl)
			tt.functionCall(mockRepo)
			svc := ProductServiceImpl{
				repo: mockRepo,
			}
			got, err := svc.GetAllProducts()
			assert.Equal(t, tt.want, got)
			assert.Equal(t, tt.expectError, err)
		})
	}
}

func TestProductServiceImpl_GetProduct(t *testing.T) {
	type fields struct {
		repo repository.ProductRepository
	}
	type args struct {
		id uint64
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    dto.Product
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			p := ProductServiceImpl{
				repo: tt.fields.repo,
			}
			got, err := p.GetProduct(tt.args.id)
			if (err != nil) != tt.wantErr {
				t.Errorf("ProductServiceImpl.GetProduct() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("ProductServiceImpl.GetProduct() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestProductServiceImpl_CreateProduct(t *testing.T) {
	type fields struct {
		repo repository.ProductRepository
	}
	type args struct {
		createObj dto.CreateProduct
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    dto.Product
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			p := ProductServiceImpl{
				repo: tt.fields.repo,
			}
			got, err := p.CreateProduct(tt.args.createObj)
			if (err != nil) != tt.wantErr {
				t.Errorf("ProductServiceImpl.CreateProduct() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("ProductServiceImpl.CreateProduct() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestProductServiceImpl_DeleteProduct(t *testing.T) {
	type fields struct {
		repo repository.ProductRepository
	}
	type args struct {
		id uint64
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			p := ProductServiceImpl{
				repo: tt.fields.repo,
			}
			if err := p.DeleteProduct(tt.args.id); (err != nil) != tt.wantErr {
				t.Errorf("ProductServiceImpl.DeleteProduct() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestProductServiceImpl_UpdateProduct(t *testing.T) {
	type fields struct {
		repo repository.ProductRepository
	}
	type args struct {
		id        uint64
		updateObj dto.UpdateProduct
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			p := ProductServiceImpl{
				repo: tt.fields.repo,
			}
			if err := p.UpdateProduct(tt.args.id, tt.args.updateObj); (err != nil) != tt.wantErr {
				t.Errorf("ProductServiceImpl.UpdateProduct() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestNewProductService(t *testing.T) {
	type args struct {
		repo repository.ProductRepository
	}
	tests := []struct {
		name string
		args args
		want ProductService
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NewProductService(tt.args.repo); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewProductService() = %v, want %v", got, tt.want)
			}
		})
	}
}
