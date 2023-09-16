package appointments

import "time"

type Appointment struct {
	IDAppointments   int       `json:"idAppointments"`
	Date             time.Time `json:"date"`
	Notes            string    `json:"quantity" `
	DentistIdDentist int       `json:"Dentist_idDentist"`
	PatientIdPatient int       `json:"Patient_idPatient"`
}
