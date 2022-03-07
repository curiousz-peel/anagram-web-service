package main

import (
	"fmt"
	"os"
	"path/filepath"
	"sort"
	"strings"
)

func buildPathToTXT() string {
	dir_path, err := os.Getwd()
	if err != nil {
		fmt.Printf(err.Error())
		os.Exit(1)
	}
	fmt.Printf(filepath.Join(dir_path, "words.txt") + "\n")
	return filepath.Join(dir_path, "words.txt")
}

func orderWord(word string) string {
	word_letters_strings := strings.Split(strings.ToLower(word), "")
	sort.Strings(word_letters_strings)
	ordered_word := strings.Join(word_letters_strings, "")
	return ordered_word
}
