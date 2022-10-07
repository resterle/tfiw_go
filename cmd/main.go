package main

import (
	"fmt"
	"io"
	"net/http"

	"github.com/resterle/tfiw_go/internal/web"
)

func main() {
	resp, err := http.Get("http://example.com/")
	if err != nil {
		// handle error
	}
	defer resp.Body.Close()
	body, err := io.ReadAll(resp.Body)

	fmt.Println(body)

	web.Run()
}
