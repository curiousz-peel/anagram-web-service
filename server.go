package main

import (
	"fmt"
	"net/http"
)

var WordsMap map[string][]string

func anagramsHandler(writer http.ResponseWriter, request *http.Request) {
	word := request.FormValue("word")
	if word == "" {
		fmt.Fprint(writer, "Inexistent or no value for 'word' as URL parameter!")
		return
	} else {
		ordered_word := orderWord(word)
		fmt.Println(WordsMap[ordered_word])
	}
}
func main() {

	WordsMap = make(map[string][]string)
	initWordsFromTXT(buildPathToTXT(), WordsMap)

	http.HandleFunc("/anagrams", anagramsHandler)
	http.ListenAndServe(":8080", nil)
}
