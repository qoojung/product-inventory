package config

import (
	"fmt"
	"os"

	"app/domain/dao"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func DatabaseConn() *gorm.DB {

	connStr := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s",
		os.Getenv("POSTGRES_HOST"), os.Getenv("POSTGRES_PORT"), os.Getenv("POSTGRES_USER"), os.Getenv("POSTGRES_PASSWORD"), os.Getenv("POSTGRES_DB"))
	db, err := gorm.Open(postgres.Open(connStr), &gorm.Config{})
	if err != nil {
		panic(err)
	}
	db.AutoMigrate(dao.Product{})
	return db
}
