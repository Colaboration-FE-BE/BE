package data

import (
	"fmt"
	"immersive-dash-4/app/middlewares"
	_teamData "immersive-dash-4/features/team/data"
	"immersive-dash-4/features/user"
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

// type role string

type User struct {
	ID          string `gorm:"primaryKey;type:varchar(255);NotNull" json:"id"`
	Fullname    string
	Email       string `gorm:"unique"`
	Password    string
	Role        string `sql:"type:ENUM('ADMIN', 'DEFAULT')" gorm:"column:role"`
	PhoneNumber string
	IdTeam      int
	Team        _teamData.Team `json:"location,omitempty" gorm:"foreignKey:IdTeam;references:ID"`
	Status      bool
	IsDeleted   bool
	CreatedAt   time.Time
	UpdatedAt   time.Time
}

func UserSeeder(db *gorm.DB) {
	id := uuid.New()

	// Hashing the password with the default cost of 10
	password := "secret"
	hashedPassword, _ := middlewares.HashedPassword(password)

	// // Comparing the password with the hash
	// err = bcrypt.CompareHashAndPassword(hashedPassword, password)
	// fmt.Println(err) // nil means it is a match
	var userSeeder = []User{
		{
			ID:          id.String(),
			Fullname:    "Admin",
			Password:    string(hashedPassword),
			PhoneNumber: "08839839293",
			Email:       "admin@altera.co.id",
			IdTeam:      3, Role: "ADMIN",
			Status: true},
	}

	result := db.Create(userSeeder) // pass a slice to insert multiple row
	fmt.Println(result)
	// result.Error                    // returns error
	// result.RowsAffected             // returns inserted records count
}

// func (u *User) BeforeCreate(tx *gorm.DB) (err error) {
// 	u.CreatedAt = time.Now().UTC()
// 	return
// }

// // TableName overrides
// func (User) TableName() string {
// 	return "users"
// }

// Mapping struct core to struct model
func CoreToModel(dataCore user.Core) User {
	return User{
		ID:          dataCore.ID,
		Fullname:    dataCore.Fullname,
		Email:       dataCore.Email,
		Role:        dataCore.Role,
		IdTeam:      int(dataCore.IdTeam),
		Password:    dataCore.Password,
		Status:      dataCore.Status,
		PhoneNumber: dataCore.PhoneNumber,
	}
}

// mapping struct model to struct core
func ModelToCore(dataModel User) user.Core {
	return user.Core{
		ID:          dataModel.ID,
		Fullname:    dataModel.Fullname,
		Role:        dataModel.Role,
		IdTeam:      dataModel.IdTeam,
		Email:       dataModel.Email,
		Password:    dataModel.Password,
		PhoneNumber: dataModel.PhoneNumber,
		CreatedAt:   dataModel.CreatedAt,
		UpdatedAt:   dataModel.UpdatedAt,
	}
}

func DetailUser(dataModel User) user.Core {
	return user.Core{
		ID:        dataModel.ID,
		Email:     dataModel.Email,
		Fullname:  dataModel.Fullname,
		IdTeam:    dataModel.IdTeam,
		Role:      dataModel.Role,
		Status:    dataModel.Status,
		CreatedAt: dataModel.CreatedAt,
		UpdatedAt: dataModel.UpdatedAt,
	}
}

// mapping []model to []core
func ListModelToCore(dataModel []User) []user.Core {
	var result []user.Core
	for _, v := range dataModel {
		result = append(result, ModelToCore(v))
	}
	return result
}
