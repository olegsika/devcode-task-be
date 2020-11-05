package models

type Movie struct {
	Id     int     `sql:"id" json:"id,omitempty"`
	Name   string  `sql:"name" json:"name,omitempty"`
	Year   int     `sql:"year" json:"year,omitempty"`
	Actors []Actor `sql:"-" json:"actors"`
}

func (a *Movie) isUpdate() bool {
	return a.Id > 0
}
