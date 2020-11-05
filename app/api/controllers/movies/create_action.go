package movies

import (
	moviesForm "devcodeIdentity/app/api/forms/movies"
	"devcodeIdentity/app/helpers/req"
	"devcodeIdentity/app/helpers/resp"
	"devcodeIdentity/app/helpers/rest"
	"encoding/json"
	"net/http"
)

var ActionCreate = rest.Action(func(w http.ResponseWriter, r *http.Request) {
	form := moviesForm.CreateForm{}

	_ = req.NewRequest(r, &form.Model)

	var err error

	err = form.Validate()

	if err != nil {
		resp.BadRequest(w, err.Error())
		return
	}

	err = form.Save()

	if err != nil {
		resp.BadRequest(w, err.Error())
		return
	}

	_ = json.NewEncoder(w).Encode(form.Model)
})