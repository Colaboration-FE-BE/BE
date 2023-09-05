package social_media

type Core struct {
	ID       int    `validate:"required"`
	Email    string `validate:"required"`
	Telegram string `validate:"required"`
	Discord  string `validate:"required"`
	Phone    string `validate:"required"`
}

type EmergencyDataInterface interface {
	// SelectAllClass() ([]Core, error)
}

type EmergencyServiceInterface interface {
	// GetAllClass() ([]Core, error)
}
