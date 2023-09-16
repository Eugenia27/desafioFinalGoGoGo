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

// func (s *SqlStore) Modify(id int, dentist patients.Denstist) (patients.Patient, error) {
// 	query := fmt.Sprintf("UPDATE Patients SET last_name = '%s', first_name = %s, registration_number = '%s' WHERE idDenstist = %v;", dentist.LastName, dentist.FirstName, dentist.RegistrationNumber, dentist.idDenstist)
// 	stmt, err := s.DB.Prepare(query)
// 	if err != nil {
// 		return patients.Patient{}, err
// 	}

// 	_, err = stmt.Exec()
// 	if err != nil {
// 		return patients.Patient{}, err
// 	}

// 	return dentist, nil
// }

// func (s *SqlStore) Save(dentist Patient) (Patient, error) {
// 	query := "INSERT INTO Patients (last_name, first_name, registration_number) VALUES ($1, $2, $3);"
// 	var idPatient int
// 	err := s.DB.QueryRow(query, dentist.LastName, dentist.FirstName, dentist.RegistrationNumber).Scan(&Patient)
// 	if err != nil {
// 		return Patient{}, err
// 	}

// 	dentist.Patient = Patient
// 	return dentist, nil
// }

// func (s *SqlStore) Delete(id int) error {
// 	query := fmt.Sprintf("DELETE FROM Patients WHERE idPatient = %d;", id)
// 	_, err := s.DB.Exec(query)
// 	if err != nil {
// 		return err
// 	}

// 	return nil
// }

/* func (s *SqlStore) GetByID(id int) (products.Product, error) {
	var productReturn products.Product

	query := fmt.Sprintf("SELECT * FROM products WHERE id = %d;", id)
	row := s.DB.QueryRow(query)
	err := row.Scan(&productReturn.ID, &productReturn.Name, &productReturn.Quantity, &productReturn.CodeValue,
		&productReturn.IsPublished, &productReturn.Expiration, &productReturn.Price)
	if err != nil {
		return products.Product{}, err
	}
	return productReturn, nil
}

func (s *SqlStore) Modify(id int, product products.Product) (products.Product, error) {
	query := fmt.Sprintf("UPDATE products SET name = '%s', quantity = %v, code_value = '%s',"+
		" is_published = %v, expiration = '%s', price = %v WHERE id = %v;", product.Name, product.Quantity,
		product.CodeValue, product.IsPublished, product.Expiration, product.Price, product.ID)
	stmt, err := s.DB.Prepare(query)
	if err != nil {
		return products.Product{}, err
	}

	_, err = stmt.Exec()
	if err != nil {
		return products.Product{}, err
	}

	return product, nil
} */
