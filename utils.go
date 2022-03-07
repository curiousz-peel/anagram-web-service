package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"os"
	"path/filepath"
	"sort"
	"strings"
)

func buildJSON(writer http.ResponseWriter, wordsMap map[string][]string, word string) []byte {
	orderedWord := orderWord(word)
	responseBody := make(map[string][]string)
	responseBody[word] = WordsMap[orderedWord]
	jsonResponse, err := json.Marshal(responseBody)
	if err != nil {
		fmt.Fprintf(writer, "Error while marshaling JSON: ")
		fmt.Fprint(writer, err)
	}
	return jsonResponse
}

func buildPathToTXT() string {
	dirPath, err := os.Getwd()
	if err != nil {
		fmt.Printf(err.Error())
		os.Exit(1)
	}
	fmt.Printf(filepath.Join(dirPath, "words.txt") + "\n")
	return filepath.Join(dirPath, "words.txt")
}

func orderWord(word string) string {
	wordLettersStrings := strings.Split(strings.ToLower(word), "")
	sort.Strings(wordLettersStrings)
	orderedWord := strings.Join(wordLettersStrings, "")
	return orderedWord
}
