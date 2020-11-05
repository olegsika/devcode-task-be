package movies

import (
	"devcodeIdentity/app/api/forms/movies"
	"devcodeIdentity/app/helpers/resp"
	"devcodeIdentity/app/helpers/rest"
	"encoding/json"
	"github.com/gorilla/mux"
	"net/http"
	"strconv"
)

var ActionDelete = rest.Action(func(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)

	form := movies.DeleteForm{}

	var err error

	form.Model.Id, _ = strconv.Atoi(params["movie_id"])

	err = form.Validate()

	if err != nil {
		resp.BadRequest(w, err.Error())
		return
	}

	err = form.Delete()

	if err != nil {
		resp.BadRequest(w, err.Error())
		return
	}

	_ = json.NewEncoder(w).Encode("Deleted!")
})
