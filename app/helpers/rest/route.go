package rest

type Route struct {
	Path   string
	Method string
	Action Action
}
