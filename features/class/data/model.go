package data

import (
	"immersive-dash-4/features/class"
	_userData "immersive-dash-4/features/user/data"
	"time"

	"gorm.io/gorm"
)

type Class struct {
	gorm.Model
	// ID           string `gorm:"primaryKey;type:varchar(255);NotNull" json:"id"`
	Name         string
	PicId        string
	User         _userData.User `json:"location,omitempty" gorm:"foreignKey:PicId;references:ID"`
	StartDate    time.Time
	GraduateDate time.Time
}

func CoreToModel(dataCore class.Core) Class {
	return Class{
		Model:        gorm.Model{},
		Name:         dataCore.Name,
		PicId:        dataCore.PicId,
		User:         _userData.User{},
		StartDate:    dataCore.StartDate,
		GraduateDate: dataCore.GraduateDate,
	}
}

// mapping struct model to struct core
func ModelToCore(dataModel Class) class.Core {

	return class.Core{
		ID:           dataModel.ID,
		PicId:        dataModel.PicId,
		Name:         dataModel.Name,
		StartDate:    dataModel.StartDate,
		GraduateDate: dataModel.GraduateDate,
	}
}

// mapping []model to []core
func ListModelToCore(dataModel []Class) []class.Core {
	var result []class.Core
	for _, v := range dataModel {
		result = append(result, ModelToCore(v))
	}
	return result
}
