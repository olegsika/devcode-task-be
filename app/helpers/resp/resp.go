package resp

import (
	"encoding/json"
	"net/http"
)

type ErrResp struct {
	Err interface{} `json:"error"`
}

func BadRequest(w http.ResponseWriter, err interface{}) {
	w.WriteHeader(http.StatusBadRequest)
	_ = json.NewEncoder(w).Encode(&ErrResp{err})
}

