package feedback

type Core struct {
	ID       int    `validate:"required"`
	Notes    string `validate:"required"`
	UserId   string `validate:"required"`
	StatusId int    `validate:"required"`
	Proof    string `validate:"required"`
	Phone    string `validate:"required"`
	MenteeId int    `validate:"required"`
}

type EmergencyDataInterface interface {
	// SelectAllClass() ([]Core, error)
}

type EmergencyServiceInterface interface {
	// GetAllClass() ([]Core, error)
}
