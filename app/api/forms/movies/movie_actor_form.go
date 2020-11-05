package movies

import (
	"devcodeIdentity/app/models"
	"errors"
)

type MovieActorForm struct {
	Model  models.MovieActors
	Action models.MoveAction
}

func (m *MovieActorForm) Validate() error {
	if m.Model.MovieId <= 0 {
		return errors.New("Movie Id can not be blank. ")
	}

	if m.Model.ActorId <= 0 {
		return errors.New("Actor Id can not be blank. ")
	}

	return nil
}

func (m *MovieActorForm) Move() error {
	switch m.Action {
	case models.ActionAdd:
		return m.add()
	case models.ActionRemove:
		return m.remove()
	default:
		return errors.New("Unsupported action. ")
	}
}

func (m *MovieActorForm) add() error {
	ex, err := models.DB().
		Model((*models.MovieActors)(nil)).
		Where("actor_id = ? AND movie_id = ?", m.Model.ActorId, m.Model.MovieId).
		Exists()

	if err != nil {
		return err
	}

	if ex {
		return errors.New("The actor already in movie. ")
	}

	return models.DB().Insert(&m.Model)
}

func (m *MovieActorForm) remove() error {
	_, err := models.DB().
		Model(&m.Model).
		Where("movie_id = ? and actor_id = ?", m.Model.MovieId, m.Model.ActorId).
		Delete()

	if err != nil {
		return err
	}

	return nil
}
