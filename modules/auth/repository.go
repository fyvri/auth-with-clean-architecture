package auth

import (
	user_entity "auth-with-clean-architecture/modules/user/entity"

	"gorm.io/gorm"
)

type Repository struct {
	db *gorm.DB
}
type RepositoryInterface interface {
	FindByUsername(username string) (*user_entity.User, error)
}

func NewRepository(db *gorm.DB) RepositoryInterface {
	return &Repository{
		db: db,
	}
}

func (r *Repository) FindByUsername(username string) (*user_entity.User, error) {
	var user *user_entity.User
	res := r.db.Where("username = ?", username).First(&user).Error

	return user, res
}
