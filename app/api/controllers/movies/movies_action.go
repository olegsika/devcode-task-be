package movies

import (
	moviesForm "devcodeIdentity/app/api/forms/movies"
	"devcodeIdentity/app/helpers/resp"
	"devcodeIdentity/app/helpers/rest"
	"encoding/json"
	"net/http"
)

var ActionMovies = rest.Action(func(w http.ResponseWriter, r *http.Request) {
	form := moviesForm.FormMovies{}

	var err error

	err = form.GetMovies()

	if err != nil {
		resp.BadRequest(w, err.Error())
		return
	}

	err = form.GetMovieActors()

	if err != nil {
		resp.BadRequest(w, err.Error())
		return
	}

	_ = json.NewEncoder(w).Encode(form.Movies)
})
