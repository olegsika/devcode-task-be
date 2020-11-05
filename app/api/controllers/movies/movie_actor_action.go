package movies

import (
	"devcodeIdentity/app/api/forms/movies"
	"devcodeIdentity/app/helpers/resp"
	"devcodeIdentity/app/helpers/rest"
	"devcodeIdentity/app/models"
	"encoding/json"
	"github.com/gorilla/mux"
	"net/http"
	"strconv"
)

var ActionMovieActor = rest.Action(func(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)

	form := movies.MovieActorForm{}

	form.Model.ActorId, _ = strconv.Atoi(params["actor_id"])
	form.Model.MovieId, _ = strconv.Atoi(params["movie_id"])
	form.Action = models.MoveAction(params["action"])

	var err error

	err = form.Validate()

	if err != nil {
		resp.BadRequest(w, err.Error())
		return
	}

	err = form.Move()

	if err != nil {
		resp.BadRequest(w, err.Error())
		return
	}

	_ = json.NewEncoder(w).Encode("Movie Updated")
})
