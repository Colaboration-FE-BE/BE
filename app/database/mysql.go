package database

import (
	"fmt"
	"immersive-dash-4/app/config"

	// "immersive-dash-4/app/config"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func InitDBMysql(cfg *config.AppConfig) *gorm.DB {
	// declare struct config & variable connectionString
	// usernama:password@tcp(hostdb:portdb)/db_name
	// connectionString := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=utf8&parseTime=True&loc=Local",
	// 	cfg.DB_USERNAME, cfg.DB_PASSWORD, cfg.DB_HOSTNAME, cfg.DB_PORT, cfg.DB_NAME)

	connectionString := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=utf8&parseTime=True&loc=Local",
		"root", "123", "34.101.64.219", 3306, "Immersive-dash-4-1")
	db, err := gorm.Open(mysql.Open(connectionString), &gorm.Config{})

	if err != nil {
		panic(err)
	}

	return db
}
