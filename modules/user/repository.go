package user

import (
	"auth-with-clean-architecture/modules/user/entity"
	"auth-with-clean-architecture/pkg/password"

	"gorm.io/gorm"
)

type Repository struct {
	db *gorm.DB
}

type RepositoryInterface interface {
	List() ([]entity.User, error)
	Create(user *entity.User) error
	Read(ID string) (*entity.User, error)
	Update(ID string, body *entity.User) (*entity.User, error)
	Delete(ID string) error
}

func NewRepository(db *gorm.DB) RepositoryInterface {
	return &Repository{
		db: db,
	}
}

func (r *Repository) List() ([]entity.User, error) {
	var users []entity.User
	err := r.db.Find(&users).Error
	return users, err
}

func (r *Repository) Create(user *entity.User) error {
	hash, _ := password.HashPassword(user.Password)
	user.Password = string(hash)
	user.RoleID = 2

	return r.db.Create(user).Error
}

func (r *Repository) Read(ID string) (*entity.User, error) {
	var user entity.User
	err := r.db.First(&user, ID).Error
	return &user, err
}

func (r *Repository) Update(ID string, user *entity.User) (*entity.User, error) {
	err := r.db.Save(&user).Error
	return user, err
}

func (r *Repository) Delete(ID string) error {
	return r.db.Delete(&entity.User{}, ID).Error
}
