package dao

import "time"

type Product struct {
	ID          uint   `gnorm:"primaryKey"`
	SKU         string `gorm:"type:varchar(64);unique_index"`
	Name        string `gorm:"type:varchar(64)"`
	Description string `gorm:"type:varchar(128)"`
	Quantity    uint
	UnitPrice   uint
	CreatedAt   time.Time
	UpdatedAt   time.Time
}
