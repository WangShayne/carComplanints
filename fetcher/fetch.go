package fetcher

import (
	"fmt"
	"io"
	"net/http"
)

func Fetch(url string) (contents io.ReadCloser) {
	resp, err := http.Get(url)
	if err != nil {
		fmt.Println(err)
	}

	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		fmt.Println("error code is :", resp.StatusCode)
	}

	return resp.Body
}
