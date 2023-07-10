package db

import (
	"database/sql"
	"fmt"
	"os"
)

func CreateDatabaseTable() {
	var (
		host   = os.Getenv("DB_HOST")
		port   = os.Getenv("DB_PORT")
		user   = os.Getenv("DB_USER")
		pass   = os.Getenv("DB_PASS")
		dbName = os.Getenv("DB_NAME")
	)

	dbConnectionString := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s", user, pass, host, port, dbName)

	db, err := sql.Open("mysql", dbConnectionString)
	if err != nil {
		panic(err.Error())
	}
	defer db.Close()

	query := `CREATE TABLE IF NOT EXISTS papers (
		id INT AUTO_INCREMENT PRIMARY KEY,
		url VARCHAR(200) NOT NULL,
		title VARCHAR(200) NOT NULL,
		github VARCHAR(200) NOT NULL
	);`

	_, err = db.Exec(query)
	if err != nil {
		panic(err.Error())
	}

}
