package movies

import "C"
import (
	"devcodeIdentity/app/models"
	"errors"
	"time"
)

type UpdateForm struct {
	Model models.Movie
}

func (u *UpdateForm) Validate() error {
	if u.Model.Id <= 0 {
		return errors.New("Movie not found. ")
	}

	if len(u.Model.Name) <= 0 {
		return errors.New("Name can nod be blank. ")
	}

	if u.Model.Year <= 0 {
		return errors.New("Year can not be blank. ")
	}

	tm := time.Now()

	if u.Model.Year > tm.Year() {
		return errors.New("The film can not be grater than current year. ")
	}

	return nil
}

func (u *UpdateForm) Update() error {
	return models.Save(&u.Model)
}
