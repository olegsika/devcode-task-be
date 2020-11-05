package actors

import (
	"devcodeIdentity/app/models"
	"errors"
)

type UpdateForm struct {
	Model models.Actor
}

func (u *UpdateForm) Validate() error {
	if u.Model.Id <= 0 {
		return errors.New("Movie not found. ")
	}

	if len(u.Model.Name) <= 0 {
		return errors.New("Name can nod be blank. ")
	}

	return nil
}

func (u *UpdateForm) Update() error {
	return models.Save(&u.Model)
}