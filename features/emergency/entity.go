package emergency

type Core struct {
	ID     int    `validate:"required"`
	Name   string `validate:"required"`
	Phone  string `validate:"required"`
	Status string `validate:"required"`
}

type EmergencyDataInterface interface {
	// SelectAllClass() ([]Core, error)
}

type EmergencyServiceInterface interface {
	// GetAllClass() ([]Core, error)
}
