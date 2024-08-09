package dto

type UpdateEmployeeRequestDTO struct {
	EmployeeID  uint
	Username    string
	Email       string
	Password    string
	FirstName   string
	LastName    string
	Position    string
	HourlyPrice float64
}
