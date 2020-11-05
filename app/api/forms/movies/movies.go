package movies

import (
	"devcodeIdentity/app/models"
	"github.com/go-pg/pg"
)

type FormMovies struct {
	Movies []models.Movie
}

func (f *FormMovies) GetMovies() error {
	return models.DB().Model(&f.Movies).Select()
}

func (f *FormMovies) GetMovieActors() error {
	var err error
	for i, movie := range f.Movies {
		var actorIds []int
		err = models.DB().
			Model((*models.MovieActors)(nil)).
			Where("movie_id = ?", movie.Id).
			Column("actor_id").
			Select(&actorIds)

		if len(actorIds) <= 0 {
			continue
		}

		if err != nil {
			return err
		}

		err = models.DB().
			Model(&movie.Actors).
			Where("id in (?)", pg.In(actorIds)).
			Select()

		if err != nil {
			return err
		}

		f.Movies[i].Actors = movie.Actors
	}

	return nil
}
