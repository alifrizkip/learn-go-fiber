package todo

import "gorm.io/gorm"

// IRepository package todo
type (
	IRepository interface {
		FindAll() ([]*Todo, error)
		FindBy(todoCond *Todo) (*Todo, error)
		FindByID(id int) (*Todo, error)
		Save(todo *Todo) (*Todo, error)
		Update(todo *Todo) (*Todo, error)
		Delete(id int) error
	}

	repository struct {
		db *gorm.DB
	}
)

// newRepository ...
func newRepository(db *gorm.DB) IRepository {
	return &repository{db}
}

func (r *repository) FindAll() ([]*Todo, error) {
	var todos []*Todo

	if err := r.db.Find(&todos).Error; err != nil {
		return nil, err
	}

	return todos, nil
}

func (r *repository) FindBy(todoCond *Todo) (*Todo, error) {
	var todo Todo

	if err := r.db.Where(todoCond).First(&todo).Error; err != nil {
		return nil, err
	}

	return &todo, nil
}

func (r *repository) FindByID(id int) (*Todo, error) {
	return r.FindBy(&Todo{ID: id})
}

func (r *repository) Save(todo *Todo) (*Todo, error) {
	if err := r.db.Create(&todo).Error; err != nil {
		return nil, err
	}

	return todo, nil
}

func (r *repository) Update(todo *Todo) (*Todo, error) {
	if err := r.db.Save(&todo).Error; err != nil {
		return nil, err
	}

	return todo, nil
}

func (r *repository) Delete(id int) error {
	return r.db.Delete(&Todo{}, id).Error
}
