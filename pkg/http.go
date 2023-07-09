package pkg

import (
	"io"
	"net/http"
)

// GetSourceHtml returns the HTML source of the given URL
func GetSourceHtml(url string) (string, error) {
	response, err := http.Get(url)
	if err != nil {
		return "", err
	}
	defer response.Body.Close()
	html, _ := io.ReadAll(response.Body)
	return string(html), nil
}
