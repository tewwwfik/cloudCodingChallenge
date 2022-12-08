package main

import (
	"bufio"
	"fmt"
	"golang.org/x/exp/maps"
	"os"
	"strings"
	"time"
)

func main() {
	begin := time.Now().UTC().UnixMilli()
	//a Map with unique keys for anagrams that stores a map for value to avoid duplicate values.
	m := make(map[[26]int32]map[string]bool)

	if len(os.Args) <= 1 {
		panic("You should give a file as argument.")
	}
	//get input file as argument.
	filePath := os.Args[1]
	readFile, err := os.Open(filePath)
	defer readFile.Close()
	if err != nil {
		panic("Couldn't read file")
	}
	fileScanner := bufio.NewScanner(readFile)
	fileScanner.Split(bufio.ScanLines)
	var data []string
	for fileScanner.Scan() {
		//writing every line to slice as new element.
		data = append(data, fileScanner.Text())
	}

	processData(data, m)

	end := time.Now().UTC().UnixMilli()
	printMap(m)
	println(begin, end, end-begin)
}

// processData groups words according to anagrams. If a word has no anagram it has group with 1 elements.
func processData(data []string, m map[[26]int32]map[string]bool) {
	for _, line := range data {
		key := createHashKey(line)
		//checks key is exist if not creates new element with the key that includes number of letters on word.
		if _, exists := m[key]; exists {
			value := m[key]
			value[line] = true
		} else {
			value := map[string]bool{line: true}
			m[key] = value
		}
	}
}

// createHashKey creates a key that includes letter count for each string
// Ex. abc 111000000... aac 20100000...
func createHashKey(input string) [26]int32 {
	var key [26]int32
	for _, runeValue := range strings.ToLower(input) {
		if runeValue-97 <= 25 {
			key[runeValue-97]++
		} else {
			panic("Input Error")
		}
	}
	return key
}

// printMap prints all anagrams that stored as key of inner map
func printMap(m map[[26]int32]map[string]bool) {
	for _, anagrams := range m {
		if len(anagrams) > 1 {
			fmt.Println(maps.Keys(anagrams))
		}
	}
}
