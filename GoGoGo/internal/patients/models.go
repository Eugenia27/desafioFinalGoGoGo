package patients

type Patient struct {
	PatientID     int    `json:"idPatient"`
	FirstName     string `json:"first_name"`
	LastName      string `json:"last_name"`
	Address       string `json:"address"`
	CredentialID  string `json:"credential_id"`
	DischargeDate string `json:"discharge_date"`
}
