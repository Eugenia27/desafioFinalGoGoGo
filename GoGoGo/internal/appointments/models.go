package appointments

type Appointment struct {
	AppointmentsID   int    `json:"idAppointments"`
	Date             string `json:"date"`
	Notes            string `json:"notes" `
	DentistIdDentist int    `json:"Dentists_idDentist"`
	PatientIdPatient int    `json:"Patients_idPatient"`
}

type AppointmentDTO struct {
	Date        string `json:"date"`
	Notes       string `json:"notes" `
	DentistName string `json:"dentist" `
	PatientName string `json:"patient" `
}
