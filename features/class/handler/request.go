package handler

import (
	"immersive-dash-4/features/class"
	"time"
)

type ClassRequest struct {
	ID           uint      `json:"id"`
	Name         string    `json:"name"`
	PicId        string    `json:"pic_id"`
	StartDate    time.Time `json:"start_date"`
	GraduateDate time.Time `json:"graduate_date"`
}

func RequestToCore(input ClassRequest) class.Core {
	return class.Core{
		ID:           input.ID,
		Name:         input.Name,
		PicId:        input.PicId,
		StartDate:    input.StartDate,
		GraduateDate: input.GraduateDate,
	}
}
