package service

import (
	"immersive-dash-4/features/feedback"

	"github.com/go-playground/validator/v10"
)

type classService struct {
	feedbackData feedback.FeedbackDataInterface
	validate     *validator.Validate
}

func New(repo feedback.FeedbackDataInterface) feedback.FeedbackServiceInterface {
	return &classService{
		feedbackData: repo,
		validate:     validator.New(),
	}
}
