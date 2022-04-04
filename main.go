package main

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"strings"
)

type handleBody struct{}

func main() {
	http.HandleFunc("/", acceptMessage)
	http.ListenAndServe(":8080", nil)
}
func acceptMessage(w http.ResponseWriter, r *http.Request) {

	io.Copy(handleBody{}, r.Body)
}
func handelMessage(m []string) {
	message := strings.Join(m, " ")
	fmt.Println(message)
}
func (handleBody) Write(b []byte) (int, error) {

	dec := json.NewDecoder(strings.NewReader(string(b)))
	for dec.More() {
		var m message
		err := dec.Decode(&m)
		if err != nil {
			log.Fatal(err)
		}
		m.log()
	}

	return 1, nil
}
