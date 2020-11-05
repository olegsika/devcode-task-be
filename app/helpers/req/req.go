package req

import (
	"encoding/json"
	"net/http"
)

func NewRequest(r *http.Request, model interface{}) error {
	decoder := json.NewDecoder(r.Body)

	return decoder.Decode(&model)
}
