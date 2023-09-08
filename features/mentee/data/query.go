package data

import (
	"fmt"
	_classData "immersive-dash-4/features/class"
	"immersive-dash-4/features/mentee"
	"strconv"

	"github.com/labstack/echo/v4"
	"gorm.io/gorm"
)

type menteeQuery struct {
	db *gorm.DB
}

func New(db *gorm.DB) mentee.MenteeDataInterface {
	return &menteeQuery{
		db: db,
	}
}

// SelectAllMentee implements mentee.MenteeDataInterface.
func (repo *menteeQuery) SelectAllMentee(c echo.Context) ([]mentee.Core, error) {
	var menteeData []Mentee
	var classData _classData.Core
	var menteeCore []mentee.Core

	class := c.QueryParam("class")
	classId, _ := strconv.Atoi(class)

	status := c.QueryParam("status")

	category := c.QueryParam("category")

	fmt.Println("Category ", category)

	if class != "" && status != "" && category != "" {
		tx := repo.db.Raw("SELECT*FROM mentees WHERE id_class=? AND status=? AND category =?", classId, status, "%"+category).Scan(&menteeData)

		if tx.Error != nil {
			return nil, tx.Error
		}
	} else if class != "" && status != "" {
		tx := repo.db.Raw("SELECT*FROM mentees WHERE id_class=? AND status=? ", classId, status).Scan(&menteeData)

		if tx.Error != nil {
			return nil, tx.Error
		}
	} else if status != "" && category != "" {
		tx := repo.db.Raw("SELECT*FROM mentees WHERE status=? AND category=? ", status, category).Scan(&menteeData)

		if tx.Error != nil {
			return nil, tx.Error
		}
	} else if class != "" {

		tx := repo.db.Raw("SELECT*FROM mentees WHERE id_class=?", classId).Scan(&menteeData)

		if tx.Error != nil {
			return nil, tx.Error
		}
	} else if status != "" {

		tx := repo.db.Raw("SELECT*FROM mentees WHERE status=?", status).Scan(&menteeData)

		if tx.Error != nil {
			return nil, tx.Error
		}
	} else if category != "" {

		tx := repo.db.Raw("SELECT*FROM mentees WHERE category LIKE?", "%"+category).Scan(&menteeData)

		if tx.Error != nil {
			return nil, tx.Error
		}
	} else {

		tx := repo.db.Raw("SELECT*FROM mentees ").Scan(&menteeData)

		if tx.Error != nil {
			return nil, tx.Error
		}
	}

	for _, value := range menteeData {
		txClass := repo.db.Raw("SELECT name FROM classes WHERE id=? ORDER BY created_at ASC", value.IdClass).Scan(&classData.Name)

		if txClass.Error != nil {
			return nil, txClass.Error
		}
		var menteeValue = mentee.Core{
			ID:          value.ID,
			Fullname:    value.Fullname,
			IdClass:     value.IdClass,
			Class:       classData.Name,
			Status:      value.Status,
			Category:    value.Category,
			Address:     value.Address,
			HomeAddress: value.HomeAddress,
			Gender:      value.Gender,
			CreatedAt:   value.CreatedAt,
			UpdatedAt:   value.UpdatedAt,
			// DeletedAt:   time.Time{},
		}
		menteeCore = append(menteeCore, menteeValue)
	}
	return menteeCore, nil
}
