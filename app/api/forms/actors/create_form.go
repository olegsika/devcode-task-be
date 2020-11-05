package actors

import (
	"devcodeIdentity/app/models"
	"errors"
)

type CreateForm struct {
	Model models.Actor
}

func (c *CreateForm) Validate() error {
	if len(c.Model.Name) <= 0 {
		return errors.New("Name can not be blank. ")
	}

	return nil
}

func (c *CreateForm) Save() error {
	return models.Save(&c.Model)
}
