package data

import (
	"immersive-dash-4/features/emergency"

	"gorm.io/gorm"
)

type emergencyQuery struct {
	db *gorm.DB
}

// SelectAllClass implements class.ClassDataInterface.
func (*emergencyQuery) SelectAllClass() ([]emergency.Core, error) {
	panic("unimplemented")
}

func New(db *gorm.DB) emergency.EmergencyDataInterface {
	return &emergencyQuery{
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
