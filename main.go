package main

import (
	"fmt"

	"github.com/GHW/pkg"
	"github.com/GHW/utils"
)

func main() {
	url := "https://paperswithcode.com/?page=2"
	regex := `<h1><a href=.([^"']*).>([^<]*)</a></h1>`

	papers, err := pkg.GetPapersFromSource(url, regex)

	if err != nil {
		panic(err)
	}

	fmt.Println(utils.PrettyPrintStruct(papers))
	fmt.Printf("You got %d results \n", len(papers))
}

// Aggregator of papers from different sources.

// Common interface for all sources.
