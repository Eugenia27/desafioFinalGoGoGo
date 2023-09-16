package patients

import "time"

type Patient struct {
	IDPatient     int       `json:"idPatient"`
	FirstName     string    `json:"first_name"`
	LastName      string    `json:"last_name"`
	Address       string    `json:"address" `
	CredentialID  string    `json:"credential_id"`
	DischargeDate time.Time `json:"discharge_date"`
}
