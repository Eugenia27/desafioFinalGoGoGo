package database

import (
	"GoGoGo/internal/patients"
	"database/sql"
	"fmt"
)

type PatientRepository struct {
	*sql.DB
}

func NewPatientRepository(db *sql.DB) *PatientRepository {
	return &PatientRepository{db}
}

func (s *PatientRepository) GetByID(id int) (patients.Patient, error) {
	var patientReturn patients.Patient

	query := fmt.Sprintf("SELECT * FROM Patients WHERE idPatient = %d;", id)
	row := s.DB.QueryRow(query)
	err := row.Scan(&patientReturn.PatientID,
		&patientReturn.FirstName,
		&patientReturn.LastName,
		&patientReturn.Address,
		&patientReturn.CredentialID,
		&patientReturn.DischargeDate)
	if err != nil {
		return patients.Patient{}, err
	}

	return patientReturn, nil
}

func (s *PatientRepository) Save(patient patients.Patient) (patients.Patient, error) {
	query := fmt.Sprintf("INSERT INTO Patients (first_name, last_name, address, credential_id, discharge_date) VALUES ( ?, ?, ?, ?, ?);")
	stmt, err := s.DB.Prepare(query)
	if err != nil {
		return patients.Patient{}, err
	}

	result, err := stmt.Exec(patient.FirstName, patient.LastName, patient.Address, patient.CredentialID, patient.DischargeDate)
	if err != nil {
		return patients.Patient{}, err
	}
	insertedId, _ := result.LastInsertId()
	patient.PatientID = int(insertedId)

	return patient, nil
}

func (s *PatientRepository) Delete(id int) error {
	query := fmt.Sprintf("DELETE FROM Patients WHERE idPatient = %d;", id)
	_, err := s.DB.Exec(query)
	if err != nil {
		return err
	}

	return nil
}

func (s *PatientRepository) ModifyByID(id int, patient patients.Patient) (patients.Patient, error) {
	query := fmt.Sprintf("UPDATE Patients SET first_name = ?, last_name = ?, address = ?, credential_id = ?, discharge_date = ?  WHERE idPatient = ?;")
	stmt, err := s.DB.Prepare(query)
	if err != nil {
		return patients.Patient{}, err
	}
	fmt.Println(query)
	_, err = stmt.Exec(patient.FirstName, patient.LastName, patient.Address, patient.CredentialID, patient.DischargeDate, id)
	if err != nil {
		return patients.Patient{}, err
	}
	patient.PatientID = id

	return patient, nil
}
