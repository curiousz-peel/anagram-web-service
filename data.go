package main

import (
	"bufio"
	"fmt"
	"io/ioutil"
	"os"
	"strings"
	"time"
)

func initWordsFromTXT(filePath string, wordsMap map[string][]string) {
	start := time.Now()
	f, err := os.Open(filePath)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error on opening words file: %v\n", err)
		os.Exit(2)
	}
	defer f.Close()
	input := bufio.NewScanner(f)
	for input.Scan() {
		orderedWord := orderWord(input.Text())
		wordsMap[orderedWord] = append(wordsMap[orderedWord], input.Text())
	}
	fmt.Printf("initWordsFromTXT time is: %s", time.Since(start))
}

func initWordsTXTInMemory(filePath string, wordsMap map[string][]string) {
	start := time.Now()
	content, err := ioutil.ReadFile(filePath)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error on opening words file: %v\n", err)
		os.Exit(2)
	}

	split := strings.Split(string(content), "\n")

	for _, word := range split {
		orderedWord := orderWord(word)
		wordsMap[orderedWord] = append(wordsMap[orderedWord], word)
	}
	fmt.Printf("initWordsTXTInMemory time is: %s", time.Since(start))
}
