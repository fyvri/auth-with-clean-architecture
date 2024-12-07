package user

import (
	"auth-with-clean-architecture/modules/user/entity"
	"fmt"
)

type UseCase struct {
	r RepositoryInterface
}

type UseCaseInterface interface {
	List() ([]entity.User, error)
	Create(item *entity.User) error
	Read(ID string) (*entity.User, error)
	Update(ID string, item *entity.User) (*entity.User, error)
	Delete(ID string) error
}

func NewUseCase(r RepositoryInterface) UseCaseInterface {
	return &UseCase{
		r: r,
	}
}

func (uc *UseCase) List() ([]entity.User, error) {
	return uc.r.List()
}

func (uc *UseCase) Create(item *entity.User) error {
	return uc.r.Create(item)
}

func (uc *UseCase) Read(ID string) (*entity.User, error) {
	user, err := uc.r.Read(ID)
	if user == nil {
		return nil, fmt.Errorf("user not found")
	}

	return user, err
}

func (uc *UseCase) Update(ID string, item *entity.User) (*entity.User, error) {
	user, err := uc.r.Read(ID)
	if err != nil {
		return nil, err
	}

	user.FullName = item.FullName

	return uc.r.Update(ID, user)
}

func (uc *UseCase) Delete(ID string) error {
	return uc.r.Delete(ID)
}
