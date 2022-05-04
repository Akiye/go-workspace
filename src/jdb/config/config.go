package config

import "database/sql"

func GetDB() (db *sql.DB, err error) {

	dbDriver := "mysql"
	dbUser := "root"
	dbPassword := "lmi123"
	dbName := "task_group"
	db, err = sql.Open(dbDriver, dbUser+":"+dbPassword+"@/"+dbName)
	return db, err
}
