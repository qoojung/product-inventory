package service

import (
	"app/domain/dao"
	"app/domain/dto"
	"app/repository"
	"app/util"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	"go.uber.org/mock/gomock"
	"gorm.io/gorm"
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
	mockctl := gomock.NewController(t)
	defer mockctl.Finish()

	tests := []struct {
		name         string
		functionCall func(mock *repository.MockProductRepository)
		want         dto.Product
		arg          uint64
		expectError  error
	}{
		{
			name: "Get Product Happy Path",
			functionCall: func(mock *repository.MockProductRepository) {
				mock.EXPECT().FindById(uint64(1)).Return(dao.Product{
					ID:          1,
					SKU:         "PROD",
					Name:        "chair",
					Description: "stylish Chair",
					Quantity:    3000,
					UnitPrice:   2000,
					CreatedAt:   time.Now(),
					UpdatedAt:   time.Now(),
				}, nil)
			},
			arg: uint64(1),
			want: dto.Product{
				ID:          1,
				SKU:         "PROD",
				Name:        "chair",
				Description: "stylish Chair",
				Quantity:    3000,
				UnitPrice:   2000,
			},
		},
		{
			name: "Get Product Fail Path",
			functionCall: func(mock *repository.MockProductRepository) {
				mock.EXPECT().FindById(uint64(1)).Return(dao.Product{}, gorm.ErrRecordNotFound)
			},
			arg:         uint64(1),
			want:        dto.Product{},
			expectError: &util.ApiError{Status: util.NotFound},
		},
		{
			name: "Get Product Fail Path: internal",
			functionCall: func(mock *repository.MockProductRepository) {
				mock.EXPECT().FindById(uint64(1)).Return(dao.Product{}, gorm.ErrInvalidDB)
			},
			arg:         uint64(1),
			want:        dto.Product{},
			expectError: gorm.ErrInvalidDB,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			mockRepo := repository.NewMockProductRepository(mockctl)
			tt.functionCall(mockRepo)
			svc := ProductServiceImpl{
				repo: mockRepo,
			}
			got, err := svc.GetProduct(tt.arg)
			assert.Equal(t, tt.want, got)
			assert.Equal(t, tt.expectError, err)
		})
	}
}

func TestProductServiceImpl_CreateProduct(t *testing.T) {
	mockctl := gomock.NewController(t)
	defer mockctl.Finish()

	tests := []struct {
		name         string
		functionCall func(mock *repository.MockProductRepository)
		want         dto.Product
		arg          dto.CreateProduct
		expectError  error
	}{
		{
			name: "Create Product",
			functionCall: func(mock *repository.MockProductRepository) {
				mock.EXPECT().Save(dao.Product{
					SKU:         "PROD",
					Name:        "chair",
					Description: "stylish Chair",
					Quantity:    3000,
					UnitPrice:   2000,
				}).Return(dao.Product{
					ID:          1,
					SKU:         "PROD",
					Name:        "chair",
					Description: "stylish Chair",
					Quantity:    3000,
					UnitPrice:   2000,
					CreatedAt:   time.Now(),
					UpdatedAt:   time.Now(),
				}, nil)
			},
			arg: dto.CreateProduct{
				SKU:         "PROD",
				Name:        "chair",
				Description: "stylish Chair",
				Quantity:    3000,
				UnitPrice:   2000,
			},
			want: dto.Product{
				ID:          1,
				SKU:         "PROD",
				Name:        "chair",
				Description: "stylish Chair",
				Quantity:    3000,
				UnitPrice:   2000,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			mockRepo := repository.NewMockProductRepository(mockctl)
			tt.functionCall(mockRepo)
			svc := ProductServiceImpl{
				repo: mockRepo,
			}
			got, err := svc.CreateProduct(tt.arg)
			assert.Equal(t, tt.want, got)
			assert.Equal(t, tt.expectError, err)
		})
	}
}

func TestProductServiceImpl_DeleteProduct(t *testing.T) {
	mockctl := gomock.NewController(t)
	defer mockctl.Finish()

	tests := []struct {
		name         string
		functionCall func(mock *repository.MockProductRepository)
		arg          uint64
		expectError  error
	}{
		{
			name: "Delete Product Happy Path",
			functionCall: func(mock *repository.MockProductRepository) {
				mock.EXPECT().Delete(uint64(1)).Return(int64(1), nil)
			},
			arg: uint64(1),
		},
		{
			name: "Delete Product Fail Path",
			functionCall: func(mock *repository.MockProductRepository) {
				mock.EXPECT().Delete(uint64(1)).Return(int64(0), nil)
			},
			arg:         uint64(1),
			expectError: &util.ApiError{Status: util.NotFound},
		},
		{
			name: "Delete Product Fail Path",
			functionCall: func(mock *repository.MockProductRepository) {
				mock.EXPECT().Delete(uint64(1)).Return(int64(0), gorm.ErrInvalidTransaction)
			},
			arg:         uint64(1),
			expectError: gorm.ErrInvalidTransaction,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			mockRepo := repository.NewMockProductRepository(mockctl)
			tt.functionCall(mockRepo)
			svc := ProductServiceImpl{
				repo: mockRepo,
			}
			err := svc.DeleteProduct(uint64(1))
			assert.Equal(t, tt.expectError, err)
		})
	}
}
func getPointer[T any](val T) *T {
	return &val
}
func TestProductServiceImpl_UpdateProduct(t *testing.T) {
	mockctl := gomock.NewController(t)
	defer mockctl.Finish()
	type Argument struct {
		update dto.UpdateProduct
		id     uint64
	}
	tests := []struct {
		name         string
		functionCall func(mock *repository.MockProductRepository)
		arg          Argument
		expectError  error
	}{
		{
			name: "Update Product Happy Path",
			functionCall: func(mock *repository.MockProductRepository) {
				mock.EXPECT().UpdateById(uint64(1), map[string]interface{}{
					"sku":         "PROD",
					"name":        "mobile",
					"description": "cutting-edge mobile",
					"quantity":    uint(1),
					"unit_price":  uint(2000),
				}).Return(int64(1), nil)
			},
			arg: Argument{
				update: dto.UpdateProduct{
					SKU:         getPointer("PROD"),
					Name:        getPointer("mobile"),
					Description: getPointer("cutting-edge mobile"),
					Quantity:    getPointer(uint(1)),
					UnitPrice:   getPointer(uint(2000)),
				},
				id: uint64(1),
			},
		},
		{
			name: "Update Product Fail Path",
			functionCall: func(mock *repository.MockProductRepository) {
				mock.EXPECT().UpdateById(uint64(1), map[string]interface{}{
					"sku":         "PROD",
					"name":        "mobile",
					"description": "cutting-edge mobile",
					"quantity":    uint(1),
					"unit_price":  uint(2000),
				}).Return(int64(0), nil)
			},
			arg: Argument{
				update: dto.UpdateProduct{
					SKU:         getPointer("PROD"),
					Name:        getPointer("mobile"),
					Description: getPointer("cutting-edge mobile"),
					Quantity:    getPointer(uint(1)),
					UnitPrice:   getPointer(uint(2000)),
				},
				id: uint64(1),
			},
			expectError: &util.ApiError{
				Status: util.NotFound,
			},
		},
		{
			name: "Update Product Fail Path",
			functionCall: func(mock *repository.MockProductRepository) {
				mock.EXPECT().UpdateById(uint64(1), map[string]interface{}{
					"sku":         "PROD",
					"name":        "mobile",
					"description": "cutting-edge mobile",
					"quantity":    uint(1),
					"unit_price":  uint(2000),
				}).Return(int64(0), gorm.ErrInvalidValue)
			},
			arg: Argument{
				update: dto.UpdateProduct{
					SKU:         getPointer("PROD"),
					Name:        getPointer("mobile"),
					Description: getPointer("cutting-edge mobile"),
					Quantity:    getPointer(uint(1)),
					UnitPrice:   getPointer(uint(2000)),
				},
				id: uint64(1),
			},
			expectError: gorm.ErrInvalidValue,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			mockRepo := repository.NewMockProductRepository(mockctl)
			tt.functionCall(mockRepo)
			svc := ProductServiceImpl{
				repo: mockRepo,
			}
			err := svc.UpdateProduct(tt.arg.id, tt.arg.update)
			assert.Equal(t, tt.expectError, err)
		})
	}
}

func TestProductServiceImpl_IncrementProductQuantity(t *testing.T) {
	mockctl := gomock.NewController(t)
	defer mockctl.Finish()

	tests := []struct {
		name         string
		functionCall func(mock *repository.MockProductRepository)
		expectError  error
	}{
		{
			name: "Update Success",
			functionCall: func(mock *repository.MockProductRepository) {
				mock.EXPECT().IncrementQuantity(uint64(1), 30).Return(int64(1), nil)
			},
		},
		{
			name: "Update Fail Out of stock",
			functionCall: func(mock *repository.MockProductRepository) {
				mock.EXPECT().IncrementQuantity(uint64(1), 30).Return(int64(0), nil)
			},
			expectError: &util.ApiError{Status: util.OperationInvalid},
		},
		{
			name: "Update Fail DB Error",
			functionCall: func(mock *repository.MockProductRepository) {
				mock.EXPECT().IncrementQuantity(uint64(1), 30).Return(int64(0), gorm.ErrInvalidField)
			},
			expectError: gorm.ErrInvalidField,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			mockRepo := repository.NewMockProductRepository(mockctl)
			tt.functionCall(mockRepo)
			svc := ProductServiceImpl{
				repo: mockRepo,
			}
			err := svc.IncrementProductQuantity(uint64(1), 30)
			assert.Equal(t, tt.expectError, err)
		})
	}
}
