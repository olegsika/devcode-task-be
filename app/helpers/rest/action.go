package rest

import "net/http"

type Action func(w http.ResponseWriter, r *http.Request)

