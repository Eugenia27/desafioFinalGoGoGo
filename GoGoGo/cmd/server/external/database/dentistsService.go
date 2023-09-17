package database

import (
	"GoGoGo/internal/dentists"
	"database/sql"
	"fmt"
)

type DentistRepository struct {
	*sql.DB
}

func NewDentistRepository(db *sql.DB) *DentistRepository {
	return &DentistRepository{db}
}


func (s *DentistRepository) GetByID(id int) (dentists.Dentist, error) {
	var dentistReturn dentists.Dentist

	query := fmt.Sprintf("SELECT * FROM Dentists WHERE idDentist = %d;", id)
	row := s.DB.QueryRow(query)
	err := row.Scan(&dentistReturn.DentistID, &dentistReturn.LastName, &dentistReturn.FirstName, &dentistReturn.RegistrationNumber)
	if err != nil {
		return dentists.Dentist{}, err
	}
	return dentistReturn, nil
}

// func (s *SqlStore) Modify(id int, dentist dentists.Denstist) (dentists.Dentist, error) {
// 	query := fmt.Sprintf("UPDATE Dentists SET last_name = '%s', first_name = %s, registration_number = '%s' WHERE idDenstist = %v;", dentist.LastName, dentist.FirstName, dentist.RegistrationNumber, dentist.idDenstist)
// 	stmt, err := s.DB.Prepare(query)
// 	if err != nil {
// 		return dentists.Dentist{}, err
// 	}

// 	_, err = stmt.Exec()
// 	if err != nil {
// 		return dentists.Dentist{}, err
// 	}

// 	return dentist, nil
// }

func (s *SqlStore) Save(dentist dentists.Dentist) (dentists.Dentist, error) {
	query := fmt.Sprintf("INSERT INTO Dentists (last_name, first_name, registration_number) VALUES ( ?, ?, ?);")
	stmt, err := s.DB.Prepare(query)
	if err != nil {
		return dentists.Dentist{}, err
	}

	result, err := stmt.Exec(dentist.LastName, dentist.FirstName, dentist.RegistrationNumber)
	if err != nil {
		return dentists.Dentist{}, err
	}
	insertedId, _ := result.LastInsertId()
	dentist.DentistID = int(insertedId)
	return dentist, nil
}

func (s *SqlStore) Delete(id int) error {
	query := fmt.Sprintf("DELETE FROM Dentists WHERE idDentist = %d;", id)
	_, err := s.DB.Exec(query)
	if err != nil {
		return err
	}

	return nil
}

func (s *SqlStore) ModifyByID(id int, dentist dentists.Dentist) (dentists.Dentist, error) {
	query := fmt.Sprintf("UPDATE Dentists SET first_name = ?, last_name = ?, registration_number = ? WHERE idDentist = ?;")
	stmt, err := s.DB.Prepare(query)
	if err != nil {
		return dentists.Dentist{}, err
	}
	fmt.Println(query)
	_, err = stmt.Exec(dentist.FirstName, dentist.LastName, dentist.RegistrationNumber, id)
	if err != nil {
		return dentists.Dentist{}, err
	}
	dentist.DentistID = id

	return dentist, nil
}
