package user

import "gorm.io/gorm"

// IRepository package user
type (
	IRepository interface {
		FindAll() ([]*User, error)
		FindBy(userCond *User) (*User, error)
		FindByID(id int) (*User, error)
		Save(user *User) (*User, error)
		Update(user *User) (*User, error)
		Delete(id int) error
	}

	repository struct {
		db *gorm.DB
	}
)

// NewRepository ...
func NewRepository(db *gorm.DB) IRepository {
	return &repository{db}
}

func (r *repository) FindAll() ([]*User, error) {
	var users []*User

	if err := r.db.Find(&users).Error; err != nil {
		return nil, err
	}

	return users, nil
}

func (r *repository) FindBy(userCond *User) (*User, error) {
	var user User

	if err := r.db.Where(userCond).First(&user).Error; err != nil {
		return nil, err
	}

	return &user, nil
}

func (r *repository) FindByID(id int) (*User, error) {
	return r.FindBy(&User{ID: id})
}

func (r *repository) Save(user *User) (*User, error) {
	if err := r.db.Create(&user).Error; err != nil {
		return nil, err
	}

	return user, nil
}

func (r *repository) Update(user *User) (*User, error) {
	if err := r.db.Save(&user).Error; err != nil {
		return nil, err
	}

	return user, nil
}

func (r *repository) Delete(id int) error {
	return r.db.Delete(&User{}, id).Error
}
