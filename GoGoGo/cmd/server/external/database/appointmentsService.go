package database

import (
	"GoGoGo/internal/appointments"
	"database/sql"
	"fmt"
	"log"
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

func (s *AppointmentRepository) GetByCredentialID(credentialId string) ([]appointments.AppointmentDTO, error) {
	var appointmentDTO []appointments.AppointmentDTO

	query := fmt.Sprintf(`SELECT ap.date
	, ap.notes
	, concat(dt.last_name,", ",dt.first_name) as dentist
	, concat(pt.last_name,", ",pt.first_name) as patient
	FROM finalGogogo.Appointments ap
	left join finalGogogo.Patients pt on ap.Patients_idPatient = pt.idPatient
	left join finalGogogo.Dentists dt on ap.Dentists_idDentist = dt.idDentist
	WHERE pt.credential_id = ?;`)

	rows, err := s.DB.Query(query, credentialId)
	if err != nil {
		return appointmentDTO, err
	}
	defer rows.Close()

	for rows.Next() {
		var appDate, appNotes, appDentist, appPatient string
		if err := rows.Scan(&appDate, &appNotes, &appDentist, &appPatient); err != nil {
			log.Fatal(err)
		}
		fmt.Printf("appDate: %s,  appNotes: %s,  appDentist: %s,  appPatient: %s \n", appDate, appNotes, appDentist, appPatient)
		appointmentDTO = append(appointmentDTO, appointments.AppointmentDTO{Date: appDate, Notes: appNotes, DentistName: appDentist, PatientName: appPatient})
	}

	if err := rows.Err(); err != nil {
		log.Fatal(err)
	}
	return appointmentDTO, nil
}
