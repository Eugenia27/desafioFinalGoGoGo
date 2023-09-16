package patients

// Repository is an interface that we used to indicate to some user how to implement
// a repository for products.
type Repository interface {
	GetByID(id int) (Patient, error)
	Modify(id int, patient Patient) (Patient, error)
}

// Service provides all functionalities related to products.

type Service struct {
	repository Repository
}

func NewService(repository Repository) *Service {
	return &Service{repository: repository}
}

func (s *Service) GetByID(id int) (Patient, error) {
	return s.repository.GetByID(id)
}

func (s *Service) ModifyByID(id int, patient Patient) (Patient, error) {
	return s.repository.Modify(id, patient)
}
