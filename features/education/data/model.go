package data

import (
	"immersive-dash-4/features/education"
	_menteeData "immersive-dash-4/features/mentee/data"
	"time"
)

type Education struct {
	ID       int `gorm:"primaryKey;NotNull" json:"id"`
	Type     string
	Major    string
	Graduate time.Time
	MenteeId string
	Mentee   _menteeData.Mentee `json:"location,omitempty" gorm:"foreignKey:MenteeId;references:ID"`
}

// mapping struct model to struct core
func ModelToCore(dataModel Education) education.Core {

	return education.Core{
		ID:       dataModel.ID,
		Type:     dataModel.Type,
		Major:    dataModel.Major,
		MenteeId: dataModel.MenteeId,
	}
}

// mapping []model to []core
// func ListModelToCore(dataModel []Class) []class.Core {
// 	var result []class.Core
// 	for _, v := range dataModel {
// 		result = append(result, ModelToCore(v))
// 	}
// 	return result
// }
