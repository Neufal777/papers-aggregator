package pkg

import (
	"io"
	"net/http"
	"strconv"
)

// GetSourceHtml returns the HTML source of the given URL
func GetSourceHtml(url string, pages int) (string, error) {
	completeHtml := ""

	for i := 0; i < pages; i++ {
		response, err := http.Get(url + strconv.Itoa(i))

		if err != nil {
			return "", err
		}
		defer response.Body.Close()
		html, _ := io.ReadAll(response.Body)
		completeHtml += string(html)
	}

	return completeHtml, nil
}
