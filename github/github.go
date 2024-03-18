package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
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
	/*if _, err := io.Copy(os.Stdout, resp.Body); err != nil {
		log.Fatalf("error: can't copy - %s", err)
	}
	*/
	// []byte, []int, []uint32, []string
	// ["hi", 7, null]
	// any (interface)
	// []any
	var r Reply
	dec := json.NewDecoder(resp.Body)
	if err := dec.Decode(&r); err != nil {
		log.Fatalf("error: can't decode - %s", err)
	}
	fmt.Println(r)
}

type Reply struct {
	Name         string
	Public_Repos int
}
