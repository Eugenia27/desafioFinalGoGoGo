package dentists

type Dentist struct {
	ID                 int    `json:"id"`
	LastName           string `json:"last_name"`
	Name               string `json:"name"`
	RegistrationNumber int    `json:"registration_number" `
}
