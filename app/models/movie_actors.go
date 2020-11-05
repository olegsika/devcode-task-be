package models

type MoveAction string

const (
	ActionAdd    MoveAction = "add"
	ActionRemove MoveAction = "remove"
)

type MovieActors struct {
	MovieId int `sql:"movie_id" json:"movie_id"`
	ActorId int `sql:"actor_id" json:"actor_id"`
}
