package customer

import (
	"auth-with-clean-architecture/modules/customer/entity"

	"gorm.io/gorm"
)

type Repository struct {
	db *gorm.DB
}

type RepositoryInterface interface {
	List() ([]entity.Customer, error)
	Create(body *entity.Customer) error
	Read(ID string) (*entity.Customer, error)
	Update(ID string, body *entity.Customer) (*entity.Customer, error)
	Delete(ID string) error
}

func NewRepository(db *gorm.DB) RepositoryInterface {
	return &Repository{
		db: db,
	}
}

func (r *Repository) List() ([]entity.Customer, error) {
	var customers []entity.Customer
	err := r.db.Find(&customers).Error
	return customers, err
}

func (r *Repository) Create(customer *entity.Customer) error {
	return r.db.Create(customer).Error
}

func (r *Repository) Read(ID string) (*entity.Customer, error) {
	var customer entity.Customer
	err := r.db.First(&customer, ID).Error
	return &customer, err
}

func (r *Repository) Update(ID string, customer *entity.Customer) (*entity.Customer, error) {
	err := r.db.Save(&customer).Error
	return customer, err
}

func (r *Repository) Delete(ID string) error {
	return r.db.Delete(&entity.Customer{}, ID).Error
}
