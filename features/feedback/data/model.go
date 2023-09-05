package data

import (
	"immersive-dash-4/features/feedback"
	_menteeData "immersive-dash-4/features/mentee/data"
	_statusLogData "immersive-dash-4/features/status_log/data"
	_userData "immersive-dash-4/features/user/data"
)

type Feedback struct {
	ID       int `gorm:"primaryKey;NotNull" json:"id"`
	Notes    string
	UserId   string
	User     _userData.User `json:"user,omitempty" gorm:"foreignKey:UserId;references:ID"`
	StatusId int
	Status   _statusLogData.StatusLog `json:"status,omitempty" gorm:"foreignKey:StatusId;references:ID"`
	Telegram string
	Discord  string
	Phone    string
	MenteeId string
	Mentee   _menteeData.Mentee `json:"mentee,omitempty" gorm:"foreignKey:MenteeId;references:ID"`
}

// mapping struct model to struct core
func ModelToCore(dataModel Feedback) feedback.Core {
	return feedback.Core{
		ID: dataModel.ID,
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
