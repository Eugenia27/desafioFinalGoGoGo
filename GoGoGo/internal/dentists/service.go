package dentists

type Repository interface {
	Save(dentist Dentist) (Dentist, error)
	GetByID(id int) (Dentist, error)
	ModifyByID(id int, dentist Dentist) (Dentist, error)
	Delete(id int) error
}

type Service struct {
	repository Repository
}

func NewService(repository Repository) *Service {
	return &Service{repository: repository}
}

func (s *Service) Save(dentist Dentist) (Dentist, error) {
	return s.repository.Save(dentist)
}

func (s *Service) GetByID(id int) (Dentist, error) {
	return s.repository.GetByID(id)
}

func (s *Service) ModifyByID(id int, dentist Dentist) (Dentist, error) {
	return s.repository.ModifyByID(id, dentist)
}

func (s *Service) Delete(id int) error {
	return s.repository.Delete(id)
}
