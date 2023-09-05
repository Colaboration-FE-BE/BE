package data

import (
	_menteeData "immersive-dash-4/features/mentee/data"
	"immersive-dash-4/features/social_media"
)

type SocialMedia struct {
	ID       int `gorm:"primaryKey;NotNull" json:"id"`
	Email    string
	Telegram string
	Discord  string
	Phone    string
	MenteeId string
	Mentee   _menteeData.Mentee `json:"location,omitempty" gorm:"foreignKey:MenteeId;references:ID"`
}

// mapping struct model to struct core
func ModelToCore(dataModel SocialMedia) social_media.Core {

	return social_media.Core{
		ID:       dataModel.ID,
		Email:    dataModel.Email,
		Telegram: dataModel.Telegram,
		Discord:  dataModel.Discord,
		Phone:    dataModel.Phone,
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
