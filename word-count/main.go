package main

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
	"sort"
	"strings"
)

func main() {
	words, err := scanWords("./hp_fs.txt")
	if err != nil {
		panic(err)
	}

	wordCount := map[string]int{}

	for _, word := range words {
		wordCount[word]++
	}

	type pair struct {
		key   string
		value int
	}

	var pairs []pair
	for w, c := range wordCount {
		pairs = append(pairs, pair{w, c})
	}

	sort.Slice(pairs, func(i, j int) bool {
		return pairs[i].value > pairs[j].value
	})

	for _, kv := range pairs[:50] {
		fmt.Printf("%s\t\t%d\n", kv.key, kv.value)
	}
}

func scanWords(path string) ([]string, error) {
	file, err := os.Open(path)
	if err != nil {
		return nil, err
	}

	defer file.Close()

	scanner := bufio.NewScanner(file)

	scanner.Split(bufio.ScanWords)

	var words []string

	re := regexp.MustCompile(`[\.!,\?]`)

	for scanner.Scan() {
		w := scanner.Text()
		words = append(words, strings.ToLower(re.ReplaceAllString(w, "")))
	}

	return words, nil
}
