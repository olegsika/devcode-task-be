package movies

import (
	moviesForm "devcodeIdentity/app/api/forms/movies"
	"devcodeIdentity/app/helpers/req"
	"devcodeIdentity/app/helpers/resp"
	"devcodeIdentity/app/helpers/rest"
	"encoding/json"
	"github.com/gorilla/mux"
	"net/http"
	"strconv"
)

var ActionUpdate = rest.Action(func(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)

	form := moviesForm.UpdateForm{}

	form.Model.Id, _ = strconv.Atoi(params["movie_id"])

	req.NewRequest(r, &form.Model)

	var err error

	err = form.Validate()

	if err != nil {
		resp.BadRequest(w, err.Error())
		return
	}

	err = form.Update()

	if err != nil {
		resp.BadRequest(w, err.Error())
		return
	}

	_ = json.NewEncoder(w).Encode(form.Model)
})
