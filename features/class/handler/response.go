package handler

import "time"

type ClassResponse struct {
	ID           int       `json:"id"`
	Name         string    `json:"name"`
	Pic          string    `json:"pic"`
	StartDate    time.Time `json:"start_date"`
	GraduateDate time.Time `json:"graduate_date"`
}
