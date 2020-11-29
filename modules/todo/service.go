package todo

type (
	// IService interface package todo
	IService interface {
		GetAllTodos() ([]*Todo, error)
		GetTodosByUserID(int) ([]*Todo, error)
		GetATodo(*Todo) (*Todo, error)
		CreateTodo(*createTodoInput, int) (*Todo, error)
		CompleteTodo(int) (*Todo, error)
		DeleteTodo(int) error
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

func (s *service) GetTodosByUserID(userID int) ([]*Todo, error) {
	return s.repo.FindAllByUserID(userID)
}

func (s *service) GetATodo(todoCond *Todo) (*Todo, error) {
	return s.repo.FindBy(todoCond)
}

func (s *service) CreateTodo(req *createTodoInput, userID int) (*Todo, error) {
	data := Todo{
		Title:  req.Title,
		Detail: req.Detail,
		IsDone: false,
		UserID: userID,
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
