package data

import (
	"immersive-dash-4/features/social_media"

	"gorm.io/gorm"
)

type socialMediaQuery struct {
	db *gorm.DB
}

// SelectAllClass implements class.ClassDataInterface.
// func (repo *socialMediaQuery) SelectAllClass() ([]SocialMedia.Core, error) {
// 	panic("unimplemented")
// }

func New(db *gorm.DB) social_media.SocialMediaDataInterface {
	return &socialMediaQuery{
		db: db,
	}
}

// SelectAllClass implements class.ClassDataInterface.
// func (repo *classQuery) SelectAllClass() ([]class.Core, error) {
// 	var classData []Class
// 	tx := repo.db.Raw("SELECT*FROM classes").Scan(&classData) // select * from users;
// 	if tx.Error != nil {
// 		return nil, tx.Error
// 	}

// 	//mapping dari struct gorm model ke struct core (entity)
// 	var classCore = ListModelToCore(classData)

// 	return classCore, nil
// }
