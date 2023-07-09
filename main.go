package main

import (
	"github.com/GHW/pkg"
)

func main() {
	url := "https://paperswithcode.com/?page="
	regex := `<h1><a href=.([^"']*).>([^<]*)</a></h1>`

	downloader := pkg.New()
	downloader.SetSource(url).
		SetRegex(regex).
		SetLimit(12).
		DownloadPapers().
		Save()

	// fmt.Println(
	// 	utils.PrettyPrintStruct(downloader),
	// )

}
