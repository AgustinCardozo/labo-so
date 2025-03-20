package database

import (
	"database/sql"
	_ "github.com/denisenkom/go-mssqldb"
	"log"
)

const Driver = "sqlserver"

func getConnectionString() string {
	return "server=127.0.0.1,1433;user id=SA;password=Password123!;database=GD2015C1;TrustServerCertificate=True;"
}

func GetSQLConnection() (db *sql.DB, err error) {
	db, err = sql.Open(Driver, getConnectionString())
	if err != nil {
		log.Fatal(err.Error())
	}

	err = db.Ping()
	if err != nil {
		log.Printf(" Error en la conexión a la base de datos: %v", err)
		return nil, err
	}
	return db, nil
}
