package actors

import (
	"devcodeIdentity/app/helpers/resp"
	"devcodeIdentity/app/helpers/rest"
	"devcodeIdentity/app/models"
	"encoding/json"
	"fmt"
	"net/http"
)

var ActionActors = rest.Action(func(w http.ResponseWriter, r *http.Request) {
	actors := []models.Actor{}

	err := models.DB().
		Model(&actors).
		Select()

	if err != nil {
		fmt.Println(err)
		resp.BadRequest(w, err.Error())
		return
	}

	_ = json.NewEncoder(w).Encode(actors)
})
