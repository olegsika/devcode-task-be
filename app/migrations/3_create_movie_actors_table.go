package main

import "github.com/go-pg/migrations"

func init() {
	var tableName = "movie_actors"

	tableColumns := []map[string]string{
		{"movie_id": "int not null"},
		{"actor_id": "int not null"},
		{"PRIMARY KEY(movie_id, actor_id)": ""},
	}

	foreignKeys := map[string]map[string]string{
		"movie_id": {
			"fk_name":    tableName + "_movie_id_fkey",
			"table_name": "movies",
			"column":     "id",
			"on_delete":  "CASCADE",
			"on_update":  "",
		},
		"actor_id": {
			"fk_name":    tableName + "_actor_id_fkey",
			"table_name": "actors",
			"column":     "id",
			"on_delete":  "CASCADE",
			"on_update":  "",
		},
	}

	_ = migrations.Register(func(db migrations.DB) error {
		var err = CreateTable(db, tableName, tableColumns)

		if err != nil {
			return err
		}

		return CreateForeignKeys(db, tableName, foreignKeys)

	}, func(db migrations.DB) error {
		var fkNames = GetFkNamesFromMap(foreignKeys)

		var err = DropForeignKeys(db, tableName, fkNames)

		if err != nil {
			return err
		}

		return DropTable(db, tableName)
	})
}
