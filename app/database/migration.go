package database

import (
	_classData "immersive-dash-4/features/class/data"
	_educationData "immersive-dash-4/features/education/data"
	_emergencyData "immersive-dash-4/features/emergency/data"
	_feedbackData "immersive-dash-4/features/feedback/data"
	_menteeData "immersive-dash-4/features/mentee/data"
	_socialMedia "immersive-dash-4/features/social_media/data"
	_statusLog "immersive-dash-4/features/status_log/data"
	_teamData "immersive-dash-4/features/team/data"
	_userData "immersive-dash-4/features/user/data"

	"gorm.io/gorm"
)

// db migration
func InitialMigration(db *gorm.DB) {
	db.AutoMigrate(&_teamData.Team{})
	db.AutoMigrate(&_userData.User{})
	db.AutoMigrate(&_classData.Class{})
	db.AutoMigrate(&_emergencyData.EmergencyData{})
	db.AutoMigrate(&_socialMedia.SocialMedia{})
	db.AutoMigrate(&_educationData.Education{})
	db.AutoMigrate(&_menteeData.Mentee{})
	db.AutoMigrate(&_statusLog.StatusLog{})
	db.AutoMigrate(&_feedbackData.Feedback{})

	/*
		TODO 2:
		migrate struct item
	*/
}
