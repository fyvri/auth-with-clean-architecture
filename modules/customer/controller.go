package customer

import (
	"auth-with-clean-architecture/modules/customer/entity"

	"gorm.io/gorm"
)

type Controller struct {
	uc UseCaseInterface
}

type ControllerInterface interface {
	List() (*[]Item, error)
	Create(body *CreateRequest) (*Item, error)
	Read(ID string) (*Item, error)
	Update(ID string, body *CreateRequest) (*Item, error)
	Delete(ID string) error
}

func NewController(uc UseCaseInterface) ControllerInterface {
	return &Controller{
		uc: uc,
	}
}

type Item struct {
	ID        uint   `json:"id"`
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
	Email     string `json:"email"`
	Avatar    string `json:"avatar"`
}

func (c *Controller) List() (*[]Item, error) {
	items, err := c.uc.List()
	if err != nil {
		return nil, err
	}

	res := &[]Item{}
	for _, item := range items {
		c := Item{
			ID:        item.ID,
			FirstName: item.FirstName,
			LastName:  item.LastName,
			Email:     item.Email,
			Avatar:    item.Avatar,
		}
		*res = append(*res, c)
	}

	return res, nil
}

func (c *Controller) Create(body *CreateRequest) (*Item, error) {
	item := entity.Customer{
		Model:     gorm.Model{},
		FirstName: body.FirstName,
		LastName:  body.LastName,
		Email:     body.Email,
	}
	err := c.uc.Create(&item)
	if err != nil {
		return nil, err
	}

	res := &Item{
		ID:        item.ID,
		FirstName: item.FirstName,
		LastName:  item.LastName,
		Email:     item.Email,
		Avatar:    item.Avatar,
	}

	return res, nil
}

func (c *Controller) Read(ID string) (*Item, error) {
	item, err := c.uc.Read(ID)
	if err != nil {
		return nil, err
	}

	res := &Item{
		ID:        item.ID,
		FirstName: item.FirstName,
		LastName:  item.LastName,
		Email:     item.Email,
		Avatar:    item.Avatar,
	}

	return res, nil
}

func (c *Controller) Update(ID string, body *CreateRequest) (*Item, error) {
	req := entity.Customer{
		FirstName: body.FirstName,
		LastName:  body.LastName,
		Email:     body.Email,
	}

	item, err := c.uc.Update(ID, &req)
	if err != nil {
		return nil, err
	}

	res := &Item{
		ID:        item.ID,
		FirstName: item.FirstName,
		LastName:  item.LastName,
		Email:     item.Email,
		Avatar:    item.Avatar,
	}

	return res, nil
}

func (c *Controller) Delete(ID string) error {
	return c.uc.Delete(ID)
}
