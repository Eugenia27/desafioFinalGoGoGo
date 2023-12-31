package patients

type Repository interface {
	Save(patient Patient) (Patient, error)
	GetByID(id int) (Patient, error)
	ModifyByID(id int, patient Patient) (Patient, error)
	Delete(id int) error
}

type Service struct {
	repository Repository
}

func NewService(repository Repository) *Service {
	return &Service{repository: repository}
}

func (s *Service) Save(patient Patient) (Patient, error) {
	return s.repository.Save(patient)
}

func (s *Service) GetByID(id int) (Patient, error) {
	return s.repository.GetByID(id)
}

func (s *Service) ModifyByID(id int, patient Patient) (Patient, error) {
	return s.repository.ModifyByID(id, patient)
}

func (s *Service) Delete(id int) error {
	return s.repository.Delete(id)
}
