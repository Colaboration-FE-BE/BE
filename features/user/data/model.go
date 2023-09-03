package data

import (
	"fmt"
	_teamData "immersive-dash-4/features/team/data"
	"time"

	"github.com/google/uuid"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

type role string

type User struct {
	ID          string `gorm:"primaryKey;type:varchar(255)" json:"id"`
	Fullname    string
	Email       string `gorm:"unique"`
	Password    string
	Role        role `sql:"type:ENUM('ADMIN', 'DEFAULT')" gorm:"column:role"`
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
	password := []byte("password")

	// Hashing the password with the default cost of 10
	hashedPassword, err := bcrypt.GenerateFromPassword(password, bcrypt.DefaultCost)
	if err != nil {
		panic(err)
	}

	// // Comparing the password with the hash
	// err = bcrypt.CompareHashAndPassword(hashedPassword, password)
	// fmt.Println(err) // nil means it is a match
	var userSeeder = []User{
		{
			ID:          id.String(),
			Fullname:    "Admin",
			Password:    string(hashedPassword),
			PhoneNumber: "08839839293",
			Email:       "admin2@altera.co.id",
			IdTeam:      3, Role: "ADMIN",
			Status: true},
	}

	result := db.Create(userSeeder) // pass a slice to insert multiple row
	fmt.Println(result)
	// result.Error                    // returns error
	// result.RowsAffected             // returns inserted records count
}

func (u *User) BeforeCreate(tx *gorm.DB) (err error) {
	u.CreatedAt = time.Now().UTC()
	return
}

// TableName overrides
func (User) TableName() string {
	return "users"
}

type UserDataInterface interface {
}

type UserServiceInterface interface {
}
