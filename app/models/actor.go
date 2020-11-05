package models

type Actor struct {
	Id        int    `sql:"id" json:"id,omitempty"`
	Name      string `sql:"name" json:"name,omitempty"`
}

func (a *Actor) isUpdate() bool {
	return a.Id > 0
}