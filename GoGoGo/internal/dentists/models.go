package dentists

type Dentist struct {
	DentistID          int    `json:"idDentist"`
	LastName           string `json:"last_name"`
	FirstName          string `json:"first_name"`
	RegistrationNumber int    `json:"registration_number" `
}
