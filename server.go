package main

import (
	"fmt"
	"net/http"
)

var WordsMap map[string][]string

func anagramsHandler(writer http.ResponseWriter, request *http.Request) {
	writer.WriteHeader(http.StatusCreated)
	writer.Header().Set("Content-Type", "application/json")
	word := request.FormValue("word")
	if word == "" {
		fmt.Fprint(writer, "Inexistent or no value for 'word' as URL parameter!")
		return
	} else {
		writer.Write(buildJSON(writer, WordsMap, word))
	}
}
func main() {

	WordsMap = make(map[string][]string)
	initWordsFromTXT(buildPathToTXT(), WordsMap)

	http.HandleFunc("/anagrams", anagramsHandler)
	http.ListenAndServe(":8080", nil)
}
