package movies

import (
	"devcodeIdentity/app/models"
	"errors"
)

type DeleteForm struct {
	Model models.Movie
}

func (d *DeleteForm) Validate() error {
	if d.Model.Id <= 0 {
		return errors.New("Can not find the movie. ")
	}

	return nil
}

func (d *DeleteForm) Delete() error {
	_, err := models.DB().Model(&d.Model).WherePK().Delete()

	if err != nil {
		return err
	}

	return nil
}