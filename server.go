package main

import (
	"fmt"
	"net/http"
	"os"
	"path/filepath"
)

func anagramsHandler(writer http.ResponseWriter, request *http.Request) {
	fmt.Fprintf(writer, "Hello World, %s!\n", request.URL.Path[1:])
	// h := request.Header
	// fmt.Fprintln(writer, h)

	request.ParseForm()
	fmt.Fprintln(writer, request.Form)
	// word := request.Form.Get("word")
	// fmt.Printf(word + "\n")

	word := request.FormValue("word")
	if word == "" {
		fmt.Fprint(writer, "Inexistent or no value for word as URL parameter!")
		return
	} else {
		fmt.Printf(word + "\n")
	}
}
func main() {
	dir_path, err := os.Getwd()
	if err != nil {
		fmt.Printf(err.Error())
		os.Exit(1)
	}
	fmt.Printf(filepath.Join(dir_path, "words.txt") + "\n")
	http.HandleFunc("/anagrams", anagramsHandler)
	http.ListenAndServe(":8080", nil)
}
