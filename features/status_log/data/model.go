package data

import (
	"time"
)

type role string

type StatusLog struct {
	ID         int `gorm:"primaryKey"`
	StatusName string
	Date       time.Time
	CreatedAt  time.Time
	UpdatedAt  time.Time
}

// func TeamSeeder(db *gorm.DB) {
// 	var teamSeeder = []Team{
// 		{Name: "Academic"},
// 		{Name: "People"},
// 		{Name: "Placement"},
// 		{Name: "Admission"},
// 		{Name: "Mentor"},
// 	}

// 	result := db.Create(teamSeeder) // pass a slice to insert multiple row
// 	fmt.Println(result)
// 	// result.Error                    // returns error
// 	// result.RowsAffected             // returns inserted records count
// }

//   result.Error        // returns error
//   result.RowsAffected // returns inserted records count

type StatusLogDataInterface interface {
}

type StatusLogServiceInterface interface {
}
