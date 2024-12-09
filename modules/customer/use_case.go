package customer

import (
	"auth-with-clean-architecture/modules/customer/entity"
	"fmt"
)

type UseCase struct {
	r RepositoryInterface
}

type UseCaseInterface interface {
	List() ([]entity.Customer, error)
	Create(item *entity.Customer) error
	Read(ID string) (*entity.Customer, error)
	Update(ID string, item *entity.Customer) (*entity.Customer, error)
	Delete(ID string) error
}

func NewUseCase(r RepositoryInterface) UseCaseInterface {
	return &UseCase{
		r: r,
	}
}

func (uc *UseCase) List() ([]entity.Customer, error) {
	return uc.r.List()
}

func (uc *UseCase) Create(item *entity.Customer) error {
	return uc.r.Create(item)
}

func (uc *UseCase) Read(ID string) (*entity.Customer, error) {
	customer, err := uc.r.Read(ID)
	if customer == nil {
		return nil, fmt.Errorf("customer not found")
	}

	return customer, err
}

func (uc *UseCase) Update(ID string, item *entity.Customer) (*entity.Customer, error) {
	customer, err := uc.r.Read(ID)
	if err != nil {
		return nil, err
	}

	customer.FirstName = item.FirstName
	customer.LastName = item.LastName
	customer.Email = item.Email

	return uc.r.Update(ID, customer)
}

func (uc *UseCase) Delete(ID string) error {
	return uc.r.Delete(ID)
}
