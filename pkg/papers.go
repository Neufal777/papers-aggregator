package pkg

type Paper struct {
	ID     int    `json:"id"`     // ID of the paper
	URL    string `json:"url"`    // URL to the paper
	Title  string `json:"title"`  // Title of the paper
	Github string `json:"github"` // Github repo of the paper
}

// func GetPapersFromSource(source string, regex string) ([]Paper, error) {
// 	paperslist := []Paper{}
// 	paperrhtml, err := GetSourceHtml(source)

// 	regexcompile := regexp.MustCompile(regex)

// 	//find matches
// 	matches := regexcompile.FindAllStringSubmatch(paperrhtml, -1)

// 	for _, match := range matches {
// 		paperslist = append(paperslist, Paper{
// 			URL:   "https://paperswithcode.com" + match[1],
// 			Title: match[2],
// 		})
// 	}

// 	if err != nil {
// 		return []Paper{}, err
// 	}

// 	return paperslist, nil
// }
