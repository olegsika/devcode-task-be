package models

import (
	"devcodeIdentity/app/helpers/db"
	"github.com/go-pg/pg"
	"sync"
)

var conn *pg.DB
var connOnce sync.Once

type Base interface {
	isUpdate() bool
}

func AuthDB() *pg.DB {
	connOnce.Do(func() {
		conn = db.NewDb()
	})

	return conn.WithContext(conn.Context())
}

func DB() *pg.DB {
	return AuthDB()
}

func Save(model Base) error {
	if model.isUpdate() {

		var _, err = DB().Model(model).WherePK().Update()
		return err
	}

	return Insert(model)
}

func Insert(model Base) error {

	return DB().Insert(model)
}
