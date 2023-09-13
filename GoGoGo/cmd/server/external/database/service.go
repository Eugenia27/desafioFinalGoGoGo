package database

import (
	"GoGoGo/internal/products"
	"database/sql"
	"fmt"
)

type SqlStore struct {
	*sql.DB
}

func NewDatabase(db *sql.DB) *SqlStore {
	return &SqlStore{db}
}

func (s *SqlStore) GetByID(id int) (products.Product, error) {
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
}
