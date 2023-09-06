package social_media

type Core struct {
	ID       int    `validate:"required"`
	Email    string `validate:"required"`
	Telegram string `validate:"required"`
	Discord  string `validate:"required"`
	Phone    string `validate:"required"`
}

type SocialMediaDataInterface interface {
	// SelectAllClass() ([]Core, error)
}

type SocialMediaServiceInterface interface {
	// GetAllClass() ([]Core, error)
}
