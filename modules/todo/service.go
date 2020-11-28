package todo

type (
	// IService interface package todo
	IService interface {
		GetAllTodos() ([]*Todo, error)
		GetATodo(id int) (*Todo, error)
		CreateTodo(req *createTodoInput) (*Todo, error)
		CompleteTodo(id int) (*Todo, error)
		DeleteTodo(id int) error
	}

	service struct {
		repo IRepository
	}
)

// newService todo
func newService(repo IRepository) IService {
	return &service{repo}
}

func (s *service) GetAllTodos() ([]*Todo, error) {
	return s.repo.FindAll()
}

func (s *service) GetATodo(id int) (*Todo, error) {
	return s.repo.FindByID(id)
}

func (s *service) CreateTodo(req *createTodoInput) (*Todo, error) {
	data := Todo{
		Title:  req.Title,
		Detail: req.Detail,
		IsDone: false,
	}

	return s.repo.Save(&data)
}

func (s *service) CompleteTodo(id int) (*Todo, error) {
	todo, err := s.repo.FindByID(id)
	if err != nil {
		return nil, err
	}

	todo.IsDone = true
	return s.repo.Update(todo)
}

func (s *service) DeleteTodo(id int) error {
	return s.repo.Delete(id)
}
