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

type FeedbackDataInterface interface {
	// SelectAllClass() ([]Core, error)
}

type FeedbackServiceInterface interface {
	// GetAllClass() ([]Core, error)
}
