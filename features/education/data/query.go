package data

import (
	"immersive-dash-4/features/education"

	"gorm.io/gorm"
)

type educationQuery struct {
	db *gorm.DB
}

// SelectAllClass implements class.ClassDataInterface.
func (*educationQuery) SelectAllClass() ([]education.Core, error) {
	panic("unimplemented")
}

func New(db *gorm.DB) education.EducationDataInterface {
	return &educationQuery{
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
