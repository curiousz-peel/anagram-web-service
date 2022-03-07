package main

import (
	"bufio"
	"fmt"
	"io/ioutil"
	"os"
	"strings"
	"time"
)

func initWordsFromTXT(file_path string, words_map map[string][]string) {
	start := time.Now()
	f, err := os.Open(file_path)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error on opening words file: %v\n", err)
		os.Exit(2)
	}
	defer f.Close()
	input := bufio.NewScanner(f)
	for input.Scan() {
		ordered_word := orderWord(input.Text())
		words_map[ordered_word] = append(words_map[ordered_word], input.Text())
	}
	fmt.Printf("initWordsFromTXT time is: %s", time.Since(start))
}

func initWordsTXTInMemory(file_path string, words_map map[string][]string) {
	start := time.Now()
	content, err := ioutil.ReadFile(file_path)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error on opening words file: %v\n", err)
		os.Exit(2)
	}

	split := strings.Split(string(content), "\n")

	for _, word := range split {
		ordered_word := orderWord(word)
		words_map[ordered_word] = append(words_map[ordered_word], word)
	}
	fmt.Printf("initWordsTXTInMemory time is: %s", time.Since(start))
}
