package database

import (
	_menteeData "immersive-dash-4/features/mentee/data"
	// _teamData "immersive-dash-4/features/team/data"
	// _userData "immersive-dash-4/features/user/data"

	"gorm.io/gorm"
)

func RunSeeder(db *gorm.DB) {
	// _teamData.TeamSeeder(db)
	// _userData.UserSeeder(db)
	_menteeData.MenteeSeeder(db)
}
