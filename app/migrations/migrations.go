package main

import (
	database "devcodeIdentity/app/helpers/db"
	"flag"
	"fmt"
	"github.com/go-pg/migrations"
	"os"
)

func main() {
	db := database.NewDb()

	flag.Parse()

	defer db.Close()

	oldVersion, newVersion, err := migrations.Run(db, flag.Args()...)

	if err != nil {
		fmt.Fprintf(os.Stderr, err.Error() + "\n")
	}

	if newVersion != oldVersion {
		fmt.Printf("migrated from version %d to %d\n", oldVersion, newVersion)
	} else {
		fmt.Printf("version is %d\n", oldVersion)
	}
}

/**
 * Up migrations functions
 */

func CreateTable(db migrations.DB, tableName string, tableColumns []map[string]string) error {
	var separate = ""

	var sql = `CREATE TABLE IF NOT EXISTS ` + tableName + ` (`

	for _, newColumnInfo := range tableColumns {

		for columnName, columnDefinition := range newColumnInfo {
			sql += separate + columnName + " " + columnDefinition

			if separate == "" {
				separate = ", "
			}
		}
	}

	sql += `);`

	_, err := db.Exec(sql)

	return err
}

func CreateForeignKeys(db migrations.DB, tableName string, foreignKeys map[string]map[string]string) error {
	var sql = ""

	for fromColumn, fkDefinition := range foreignKeys {
		sql += "ALTER TABLE " + tableName
		sql += " ADD CONSTRAINT " + fkDefinition["fk_name"] + " FOREIGN KEY (" + fromColumn + ")"
		sql += " REFERENCES " + fkDefinition["table_name"] + "(" + fkDefinition["column"] + ") "

		if fkDefinition["on_delete"] != "" {
			sql += " ON DELETE " + fkDefinition["on_delete"]
		}
		if fkDefinition["on_update"] != "" {
			sql += " ON UPDATE " + fkDefinition["on_update"]
		}

		sql += ";"
	}

	if len(sql) <= 0 {
		return nil
	}
	_, err := db.Exec(sql)

	return err
}

func GetFkNamesFromMap(foreignKeys map[string]map[string]string) []string {
	var fkNames []string

	for _, fkDefinition := range foreignKeys {
		fkNames = append(fkNames, fkDefinition["fk_name"])
	}

	return fkNames
}

/**
 * Down migrations functions
 */

func DropTable(db migrations.DB, tableName string) error {
	var sql = `DROP TABLE ` + tableName + `;`

	_, err := db.Exec(sql)

	return err
}

func DropForeignKeys(db migrations.DB, tableName string, fkNames []string) error {
	var sql = ``

	for _, fkName := range fkNames {
		sql += `ALTER TABLE ` + tableName + ` DROP CONSTRAINT "` + fkName + `";`
	}

	_, err := db.Exec(sql)

	return err
}