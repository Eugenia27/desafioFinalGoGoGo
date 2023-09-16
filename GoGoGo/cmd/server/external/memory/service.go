package memory

import (
	"GoGoGo/internal/products"
	"encoding/json"
	"errors"
	"os"
)

// Service provides functionalities related to memory products storage
type Service struct {
	products []products.Product
}

func NewService(path string) (*Service, error) {
	byteFile, err := os.ReadFile(path)
	if err != nil {
		return nil, err
	}
	var products []products.Product
	err = json.Unmarshal(byteFile, &products)
	if err != nil {
		return nil, err
	}
	return &Service{products: products}, nil
}

func (s *Service) GetByID(id int) (products.Product, error) {
	for _, value := range s.products {
		if value.ID == id {
			return value, nil
		}
	}
	return products.Product{}, errors.New("not found")
}

func (s *Service) Modify(id int, product products.Product) (products.Product, error) {
	for i, value := range s.products {
		if value.ID == id {
			s.products = append(s.products[:i], s.products[i+1:]...)
			s.products = append(s.products, product)
			return product, nil
		}
	}
	return products.Product{}, errors.New("not found")
}
