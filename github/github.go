package main

import (
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
)

func main() {
	resp, err := http.Get("https://api.github.com/users/m-maina")
	if err != nil {
		log.Fatalf("error: %s", err)
	}
	if resp.StatusCode != http.StatusOK {
		log.Fatalf("error: %s", resp.Status)
	}
	fmt.Printf("Content-Type: %s\n", resp.Header.Get("Content_Type"))
	io.Copy(os.Stdout, resp.Body)
}
