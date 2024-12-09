package user

import (
	"auth-with-clean-architecture/modules/user/entity"
	"fmt"

	"gorm.io/gorm"
)

type Controller struct {
	uc UseCaseInterface
}

type ControllerInterface interface {
	List() (*[]Item, error)
	Create(body *CreateRequest) (*Item, error)
	Read(ID string) (*Item, error)
	Update(ID string, body *UpdateRequest) (*Item, error)
	Delete(ID string) error
}

func NewController(uc UseCaseInterface) ControllerInterface {
	return &Controller{
		uc: uc,
	}
}

type Item struct {
	ID       uint   `json:"id"`
	FullName string `json:"first_name"`
	Username string `json:"username"`
}

func (c *Controller) List() (*[]Item, error) {
	users, err := c.uc.List()
	if err != nil {
		return nil, err
	}

	res := &[]Item{}
	for _, user := range users {
		c := Item{
			ID:       user.ID,
			FullName: user.FullName,
			Username: user.Username,
		}
		*res = append(*res, c)
	}

	return res, nil
}

func (c *Controller) Create(body *CreateRequest) (*Item, error) {
	user := entity.User{
		Model:    gorm.Model{},
		FullName: body.FullName,
		Username: body.Username,
		Password: body.Password,
	}
	err := c.uc.Create(&user)
	if err != nil {
		return nil, err
	}

	res := &Item{
		ID:       user.ID,
		FullName: body.FullName,
		Username: body.Username,
	}

	return res, nil
}

func (c *Controller) Read(ID string) (*Item, error) {
	user, err := c.uc.Read(ID)
	if err != nil {
		return nil, err
	}

	if user == nil {
		return nil, fmt.Errorf("user not found")
	}

	res := &Item{
		ID:       user.ID,
		FullName: user.FullName,
		Username: user.Username,
	}

	return res, nil
}

func (c *Controller) Update(ID string, body *UpdateRequest) (*Item, error) {
	req := entity.User{
		FullName: body.FullName,
	}

	user, err := c.uc.Update(ID, &req)
	if err != nil {
		return nil, err
	}

	res := &Item{
		ID:       user.ID,
		FullName: user.FullName,
		Username: user.Username,
	}

	return res, nil
}

func (c *Controller) Delete(ID string) error {
	return c.uc.Delete(ID)
}
