package movies

import (
	"devcodeIdentity/app/models"
	"errors"
	"time"
)

type CreateForm struct {
	Model models.Movie
}

func (c *CreateForm) Validate() error {
	if len(c.Model.Name) <= 0 {
		return errors.New("Name can not be blank. ")
	}

	if c.Model.Year <= 0 {
		return errors.New("Year can not be blank. ")
	}

	tm := time.Now()

	if c.Model.Year > tm.Year() {
		return errors.New("The film can not be grater than current year. ")
	}

	return nil
}

func (c *CreateForm) Save() error {
	return models.Save(&c.Model)
}
