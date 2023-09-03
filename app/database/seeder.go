package database

import (
	_teamData "immersive-dash-4/features/team/data"
	_userData "immersive-dash-4/features/user/data"

	"gorm.io/gorm"
)

func RunSeeder(db *gorm.DB) {
	_teamData.TeamSeeder(db)
	_userData.UserSeeder(db)
}
