package appointments

// Repository is an interface that we used to indicate to some user how to implement
// a repository for products.

type Repository interface {
	Save(appointment Appointment) (Appointment, error)
	GetByID(id int) (Appointment, error)
	ModifyByID(id int, appointment Appointment) (Appointment, error)
	Delete(id int) error
	GetByCredentialID(credentialId string) ([]AppointmentDTO, error)
}

// Service provides all functionalities related to appointment.
type Service struct {
	repository Repository
}

func NewService(repository Repository) *Service {
	return &Service{repository: repository}
}

func (s *Service) Save(appointment Appointment) (Appointment, error) {
	return s.repository.Save(appointment)
}

func (s *Service) GetByID(id int) (Appointment, error) {
	return s.repository.GetByID(id)
}

func (s *Service) ModifyByID(id int, appointment Appointment) (Appointment, error) {
	return s.repository.ModifyByID(id, appointment)
}

func (s *Service) Delete(id int) error {
	return s.repository.Delete(id)
}

func (s *Service) GetByCredentialID(credentialId string) ([]AppointmentDTO, error) {
	return s.repository.GetByCredentialID(credentialId)
}
