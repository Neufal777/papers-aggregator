package main

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"net/http"
	"os"

	"github.com/GHW/pkg"
	_ "github.com/go-sql-driver/mysql"
)

func main() {
	http.HandleFunc("/data", getDataHandler)
	http.ListenAndServe(":8080", nil)
}

func getDataHandler(w http.ResponseWriter, r *http.Request) {
	query := r.URL.Query().Get("q")

	results := DatabaseSearch(query)
	resultsMarshalldata, err := json.Marshal(results)

	if err != nil {
		fmt.Println(err)
	}

	w.Header().Set("Content-Type", "application/json")
	w.Write([]byte(resultsMarshalldata))
}

func Execute() {
	url := "https://paperswithcode.com/?page="
	regex := `<h1><a href=.([^"']*).>([^<]*)[^']*github.com/([^"]*)`

	downloader := pkg.New()
	downloader.SetSource(url).
		SetRegex(regex).
		SetLimit(13).
		DownloadPapers().
		Save()

}

func DatabaseSearch(query string) []pkg.Paper {
	//Save into a database
	var (
		host   = os.Getenv("DB_HOST")
		port   = os.Getenv("DB_PORT")
		user   = os.Getenv("DB_USER")
		pass   = os.Getenv("DB_PASS")
		dbName = os.Getenv("DB_NAME")
	)
	collection := []pkg.Paper{}
	dbConnectionString := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s", user, pass, host, port, dbName)

	db, err := sql.Open("mysql", dbConnectionString)

	if err != nil {
		panic(err.Error())
	}

	defer db.Close()

	// Query the database
	rows, err := db.Query("SELECT * FROM papers WHERE title LIKE CONCAT('%', ?, '%')", query)

	if err != nil {
		panic(err.Error())
	}

	defer rows.Close()

	for rows.Next() {
		var id int
		var url string
		var title string
		var github string

		err = rows.Scan(&id, &url, &title, &github)

		if err != nil {
			panic(err.Error())
		}

		collection = append(collection, pkg.Paper{
			ID:     id,
			URL:    url,
			Title:  title,
			Github: github,
		})
	}

	return collection
}
