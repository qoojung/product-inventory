package repository

import (
	"app/domain/dao"
	"testing"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/stretchr/testify/assert"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func CreateDB() (*gorm.DB, sqlmock.Sqlmock) {
	mockDb, mock, _ := sqlmock.New(sqlmock.QueryMatcherOption(sqlmock.QueryMatcherEqual))
	dialector := postgres.New(postgres.Config{
		Conn:       mockDb,
		DriverName: "postgres",
	})
	db, err := gorm.Open(dialector, &gorm.Config{})
	if err != nil {
		panic(err)
	}
	return db, mock
}
func TestProductRepositoryImpl_FindById(t *testing.T) {
	db, mock := CreateDB()
	prod := ProductRepositoryImpl{db: db}
	mock.ExpectQuery(`SELECT * FROM "products" WHERE id = $1 ORDER BY "products"."id" LIMIT $2`).
		WithArgs(1, 1).
		WillReturnRows(sqlmock.NewRows([]string{"id", "sku", "name", "description", "quantity", "unit_price"}).
			AddRow(1, "test", "car", "car", uint(60), uint(60)))
	product, err := prod.FindById(1)
	assert.Equal(t, err, nil)
	assert.Equal(t, dao.Product{ID: 1, SKU: "test", Name: "car", Description: "car", Quantity: 60, UnitPrice: 60}, product)

}

func TestProductRepositoryImpl_FindAll(t *testing.T) {
	db, mock := CreateDB()
	prod := ProductRepositoryImpl{db: db}
	mock.ExpectQuery(`SELECT * FROM "products"`).WithoutArgs().WillReturnRows(sqlmock.NewRows([]string{"id", "sku", "name", "description", "quantity", "unit_price"}).
		AddRow(1, "test", "car", "car", uint(60), uint(60)))
	t.Run("FindAll", func(t *testing.T) {
		product, err := prod.FindAll()
		assert.Equal(t, err, nil)
		assert.Equal(t, []dao.Product{{ID: 1, SKU: "test", Name: "car", Description: "car", Quantity: 60, UnitPrice: 60}}, product)
	})

}

func TestProductRepositoryImpl_DeleteById(t *testing.T) {
	db, mock := CreateDB()
	prod := ProductRepositoryImpl{db: db}
	t.Run("Delete", func(t *testing.T) {
		mock.ExpectBegin()
		mock.ExpectExec(`DELETE FROM "products" WHERE "products"."id" = $1`).WithArgs(1).WillReturnResult(sqlmock.NewResult(1, 1))
		mock.ExpectCommit()
		row, err := prod.Delete(1)
		assert.Equal(t, err, nil)
		assert.Equal(t, row, int64(1))
	})
	t.Run("Delete Fail", func(t *testing.T) {
		mock.ExpectBegin()
		mock.ExpectExec(`DELETE FROM "products" WHERE "products"."id" = $1`).WithArgs(1).WillReturnResult(sqlmock.NewResult(1, 0))
		mock.ExpectCommit()
		row, err := prod.Delete(1)
		assert.Equal(t, err, nil)
		assert.Equal(t, row, int64(0))
	})

}

func TestProductRepositoryImpl_UpdateById(t *testing.T) {
	db, mock := CreateDB()
	prod := ProductRepositoryImpl{db: db}
	t.Run("Update OK", func(t *testing.T) {
		mock.ExpectBegin()
		mock.ExpectExec(`UPDATE "products" SET "unit_price"=$1,"updated_at"=$2 WHERE id = $3`).WithArgs(2000, sqlmock.AnyArg(), 1).WillReturnResult(sqlmock.NewResult(1, 1))
		mock.ExpectCommit()
		row, err := prod.UpdateById(1, map[string]interface{}{"unit_price": uint(2000)})
		assert.Equal(t, err, nil)
		assert.Equal(t, row, int64(1))
	})
	t.Run("Update Fail", func(t *testing.T) {
		mock.ExpectBegin()
		mock.ExpectExec(`UPDATE "products" SET "unit_price"=$1,"updated_at"=$2 WHERE id = $3`).WithArgs(100, sqlmock.AnyArg(), 1).WillReturnResult(sqlmock.NewResult(1, 0))
		mock.ExpectCommit()
		row, err := prod.UpdateById(1, map[string]interface{}{"unit_price": uint(100)})
		assert.Equal(t, err, nil)
		assert.Equal(t, row, int64(0))
	})

}
