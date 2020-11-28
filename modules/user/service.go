package user

type (
	// IService interface package user
	IService interface {
		GetUserByID(id int) (*User, error)
	}

	// service ...
	service struct {
		repo IRepository
	}
)

// NewService ...
func NewService(repo IRepository) IService {
	return &service{repo}
}

// GetUserByID ...
func (s *service) GetUserByID(id int) (*User, error) {
	return s.repo.FindByID(id)
}