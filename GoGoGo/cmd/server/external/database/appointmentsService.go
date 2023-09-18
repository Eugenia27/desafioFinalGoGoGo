package database

import (
	"GoGoGo/internal/appointments"
	"database/sql"
	"fmt"
)

type AppointmentRepository struct {
	*sql.DB
}

func NewAppointmentRepository(db *sql.DB) *AppointmentRepository {
	return &AppointmentRepository{db}
}

func (s *AppointmentRepository) GetByID(id int) (appointments.Appointment, error) {
	var appointmentReturn appointments.Appointment

	query := fmt.Sprintf("SELECT * FROM Appointments WHERE idAppointments = %d;", id)
	row := s.DB.QueryRow(query)
	err := row.Scan(&appointmentReturn.AppointmentsID,
		&appointmentReturn.Date,
		&appointmentReturn.Notes,
		&appointmentReturn.DentistIdDentist,
		&appointmentReturn.PatientIdPatient)
	if err != nil {
		return appointments.Appointment{}, err
	}

	return appointmentReturn, nil
}

func (s *AppointmentRepository) Save(appointment appointments.Appointment) (appointments.Appointment, error) {
	query := fmt.Sprintf("INSERT INTO Appointments (`date`, notes, Dentists_idDentist, Patients_idPatient) VALUES ( ?, ?, ?, ?);")
	stmt, err := s.DB.Prepare(query)
	if err != nil {
		return appointments.Appointment{}, err
	}

	result, err := stmt.Exec(appointment.Date, appointment.Notes, appointment.DentistIdDentist, appointment.PatientIdPatient)
	if err != nil {
		return appointments.Appointment{}, err
	}
	insertedId, _ := result.LastInsertId()
	appointment.AppointmentsID = int(insertedId)

	return appointment, nil
}

func (s *AppointmentRepository) Delete(id int) error {
	query := fmt.Sprintf("DELETE FROM Appointments WHERE idAppointments = %d;", id)
	_, err := s.DB.Exec(query)
	if err != nil {
		return err
	}

	return nil
}

func (s *AppointmentRepository) ModifyByID(id int, appointment appointments.Appointment) (appointments.Appointment, error) {
	query := fmt.Sprintf("UPDATE Appointments SET date = ?, notes = ?, Dentists_idDentist = ?, Patients_idPatient = ?  WHERE idAppointments = ?;")
	stmt, err := s.DB.Prepare(query)
	if err != nil {
		return appointments.Appointment{}, err
	}
	fmt.Println(query)
	_, err = stmt.Exec(appointment.Date, appointment.Notes, appointment.DentistIdDentist, appointment.PatientIdPatient, id)
	if err != nil {
		return appointments.Appointment{}, err
	}
	appointment.AppointmentsID = id

	return appointment, nil
}

// func (s *AppointmentRepository) GetByAppointmentsByPatientCredential(credentialId string) ([]appointments.Appointment, error) {
// 	var appointmentReturn []appointments.Appointment
// 	var patientId int

// 	patientId = patientsService.GetByCredentialID(credentialId)

// 	query := fmt.Sprintf("SELECT * FROM Appointments WHERE idAppointments = %d;", id)
// 	row := s.DB.QueryRow(query)
// 	err := row.Scan(&appointmentReturn.AppointmentsID,
// 		&appointmentReturn.Date,
// 		&appointmentReturn.Notes,
// 		&appointmentReturn.DentistIdDentist,
// 		&appointmentReturn.PatientIdPatient)
// 	if err != nil {
// 		return appointments.Appointment{}, err
// 	}

// 	return appointmentReturn, nil
// }
