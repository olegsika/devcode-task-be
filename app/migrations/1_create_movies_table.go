package main

import "github.com/go-pg/migrations"

func init() {
	var tableName = "movies"

	tableColumns := []map[string]string{
		{"id": "SERIAL PRIMARY KEY"},
		{"name": "varchar(255) not null"},
		{"year": "int not null"},
	}

	_ = migrations.Register(func(db migrations.DB) error {

		return CreateTable(db, tableName, tableColumns)

	}, func(db migrations.DB) error {
		return DropTable(db, tableName)
	})
}
