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

// func (s *SqlStore) Save(dentist Dentist) (Dentist, error) {
// 	query := "INSERT INTO Dentists (last_name, first_name, registration_number) VALUES ($1, $2, $3);"
// 	var idDentist int
// 	err := s.DB.QueryRow(query, dentist.LastName, dentist.FirstName, dentist.RegistrationNumber).Scan(&Dentist)
// 	if err != nil {
// 		return Dentist{}, err
// 	}

// 	dentist.Dentist = Dentist
// 	return dentist, nil
// }

// func (s *SqlStore) Delete(id int) error {
// 	query := fmt.Sprintf("DELETE FROM Dentists WHERE idDentist = %d;", id)
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
