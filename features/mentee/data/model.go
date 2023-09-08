package data

import (
	"fmt"
	_classData "immersive-dash-4/features/class/data"
	"immersive-dash-4/features/mentee"

	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Mentee struct {
	ID          string `gorm:"primaryKey;type:varchar(255);NotNull" json:"id"`
	Fullname    string
	IdClass     int
	Class       _classData.Class `json:"location,omitempty" gorm:"foreignKey:IdClass;references:ID"`
	Status      string
	Category    string
	Address     string
	HomeAddress string
	Gender      bool
	CreatedAt   time.Time
	UpdatedAt   time.Time
	DeletedAt   time.Time
}

func MenteeSeeder(db *gorm.DB) {
	id := uuid.New()
	var menteeSeeder = []Mentee{
		{
			ID:          id.String(),
			Fullname:    "Aldi Taher",
			IdClass:     5,
			Class:       _classData.Class{},
			Status:      "Graduated",
			Category:    "NON IT",
			Address:     "Bandung",
			HomeAddress: "Bandung no 50",
			Gender:      true,
			CreatedAt:   time.Time{},
			UpdatedAt:   time.Time{},
			DeletedAt:   time.Now(),
		},
	}

	result := db.Create(menteeSeeder) // pass a slice to insert multiple row
	fmt.Println(result)
	// result.Error                    // returns error
	// result.RowsAffected             // returns inserted records count
}

// func MenteeSeeder(db *gorm.DB) {
// 	id := uuid.New()

// 	// Hashing the password with the default cost of 10
// 	password := "secret"
// 	hashedPassword, _ := middlewares.HashedPassword(password)

// 	// // Comparing the password with the hash
// 	// err = bcrypt.CompareHashAndPassword(hashedPassword, password)
// 	// fmt.Println(err) // nil means it is a match
// 	var userSeeder = []Mentee{
// 		{
// 			ID:          id.String(),
// 			Fullname:    "Admin",
// 			Password:    string(hashedPassword),
// 			PhoneNumber: "08839839293",
// 			Email:       "admin@altera.co.id",
// 			IdTeam:      3, Role: "ADMIN",
// 			Status: true},
// 	}

// 	result := db.Create(userSeeder) // pass a slice to insert multiple row
// 	fmt.Println(result)
// 	// result.Error                    // returns error
// 	// result.RowsAffected             // returns inserted records count
// }

// func (u *User) BeforeCreate(tx *gorm.DB) (err error) {
// 	u.CreatedAt = time.Now().UTC()
// 	return
// }

// // TableName overrides
// func (User) TableName() string {
// 	return "users"
// }

// Mapping struct core to struct model
func CoreToModel(dataCore mentee.Core) Mentee {
	return Mentee{
		ID:          dataCore.ID,
		Fullname:    dataCore.Fullname,
		IdClass:     dataCore.IdClass,
		Class:       _classData.Class{},
		Status:      dataCore.Status,
		Address:     dataCore.Address,
		HomeAddress: dataCore.HomeAddress,
		Gender:      dataCore.Gender,
		CreatedAt:   time.Time{},
		UpdatedAt:   time.Time{},
		DeletedAt:   time.Time{},
	}
}

// mapping struct model to struct core
func ModelToCore(dataModel Mentee) mentee.Core {
	return mentee.Core{
		ID:       dataModel.ID,
		Fullname: dataModel.Fullname,
		IdClass:  dataModel.IdClass,
		// Class:       ,
		Status:      dataModel.Status,
		Address:     dataModel.Address,
		HomeAddress: dataModel.HomeAddress,
		Gender:      dataModel.Gender,
		CreatedAt:   dataModel.CreatedAt,
		UpdatedAt:   dataModel.UpdatedAt,
		DeletedAt:   dataModel.DeletedAt,
	}
}

func DetailMentee(dataModel Mentee) mentee.Core {
	return mentee.Core{
		ID:        dataModel.ID,
		Status:    dataModel.Status,
		CreatedAt: dataModel.CreatedAt,
		UpdatedAt: dataModel.UpdatedAt,
	}
}

// mapping []model to []core
func ListModelToCore(dataModel []Mentee) []mentee.Core {
	var result []mentee.Core
	for _, v := range dataModel {
		result = append(result, ModelToCore(v))
	}
	return result
}
