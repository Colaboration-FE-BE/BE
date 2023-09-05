package education

type Core struct {
	ID       int    `validate:"required"`
	Type     string `sql:"type:ENUM('INFORMATICS', 'NON-INFORMATICS')" gorm:"column:type"`
	Major    string `validate:"required"`
	MenteeId string `validate:"required"`
}

type EmergencyDataInterface interface {
	// SelectAllClass() ([]Core, error)
}

type EmergencyServiceInterface interface {
	// GetAllClass() ([]Core, error)
}
