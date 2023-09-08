package service

import (
	"immersive-dash-4/features/social_media"

	"github.com/go-playground/validator/v10"
)

type socialMediaService struct {
	social_media social_media.SocialMediaDataInterface
	validate     *validator.Validate
}

func New(repo social_media.SocialMediaDataInterface) social_media.SocialMediaServiceInterface {
	return &socialMediaService{
		social_media: repo,
		validate:     validator.New(),
	}
}
