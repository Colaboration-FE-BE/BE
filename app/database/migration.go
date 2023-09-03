package database

import (
	_teamData "immersive-dash-4/features/team/data"
	_userData "immersive-dash-4/features/user/data"

	"gorm.io/gorm"
)

// db migration
func InitialMigration(db *gorm.DB) {
	db.AutoMigrate(&_teamData.Team{})
	db.AutoMigrate(&_userData.User{})

	/*
		TODO 2:
		migrate struct item
	*/
}
