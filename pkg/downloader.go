package pkg

import (
	"database/sql"
	"fmt"
	"os"
	"regexp"

	_ "github.com/go-sql-driver/mysql"
)

type Downloader struct {
	Source string  // Source of the paper
	Regex  string  // Regex to match the paper
	Limit  int     // Limit the number of results
	Papers []Paper // List of papers
}

// New downloader
func New() Downloader {
	return Downloader{}
}

// SetSource sets the source of the results. ie: https://paperswithcode.com/?page=2
func (d *Downloader) SetSource(source string) *Downloader {
	d.Source = source
	return d
}

// SetRegex sets the regex to match the paper, ie: `<h1><a href=.([^"']*).>([^<]*)</a></h1>`
func (d *Downloader) SetRegex(regex string) *Downloader {
	d.Regex = regex
	return d
}

// SetLimit sets the limit of the pages, ie: 10
func (d *Downloader) SetLimit(limit int) *Downloader {
	d.Limit = limit
	return d
}

// GetLimit returns the limit of the results, ie: 10
func (d *Downloader) GetLimit() int {
	return d.Limit
}

// GetSource returns the source of the results. ie: https://paperswithcode.com/?page=2
func (d *Downloader) GetSource() string {
	return d.Source
}

// GetRegex returns the regex to match the paper, ie: `<h1><a href=.([^"']*).>([^<]*)</a></h1>`
func (d *Downloader) GetRegex() string {
	return d.Regex
}

// GetPapersList returns the list of papers
func (d *Downloader) GetPapers() []Paper {
	return d.Papers
}

func (d *Downloader) DownloadPapers() *Downloader {
	paperrhtml, err := GetSourceHtml(d.Source, d.Limit)
	regexcompile := regexp.MustCompile(d.Regex)

	//find matches
	matches := regexcompile.FindAllStringSubmatch(paperrhtml, -1)

	for _, match := range matches {
		d.Papers = append(d.Papers, Paper{
			URL:    "https://paperswithcode.com" + match[1],
			Title:  match[2],
			Github: match[3],
		})
	}

	if err != nil {
		return d
	}

	return d
}

func (d *Downloader) Save() *Downloader {
	//Save into a database
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

	for _, paper := range d.Papers {
		// check if the paper already exists
		query := fmt.Sprintf("SELECT COUNT(*) FROM papers WHERE url = '%s'", paper.URL)
		var count int
		err := db.QueryRow(query).Scan(&count)
		if err != nil {
			panic(err.Error())
		}

		if count == 0 {
			query := fmt.Sprintf("INSERT INTO papers (url, title, github) VALUES ('%s', '%s','%s')", paper.URL, paper.Title, paper.Github)

			_, err = db.Exec(query)
			if err != nil {
				panic(err.Error())
			}
			fmt.Printf("saving paper %s \n", paper.Title)

		} else {
			fmt.Printf("The paper %s already exists\n", paper.Title)
		}
	}

	fmt.Println("All the papers have been saved into the database (if they didn't exist)")
	return d
}
