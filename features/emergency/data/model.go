package data

import (
	"immersive-dash-4/features/emergency"
	_menteeData "immersive-dash-4/features/mentee/data"
)

type EmergencyData struct {
	ID       int `gorm:"primaryKey;NotNull" json:"id"`
	Name     string
	Phone    string
	Status   string
	MenteeId string
	Mentee   _menteeData.Mentee `json:"location,omitempty" gorm:"foreignKey:MenteeId;references:ID"`
}

// mapping struct model to struct core
func ModelToCore(dataModel EmergencyData) emergency.Core {

	return emergency.Core{
		ID:     dataModel.ID,
		Name:   dataModel.Name,
		Phone:  dataModel.Phone,
		Status: dataModel.Status,
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
