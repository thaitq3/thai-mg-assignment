package handler

import (
	"io/ioutil"
	"net/http"
	"sort"
	"strings"
	"word-counting/utils"
)

type WordHandler struct {
}

type WordCount struct {
	Word  string `json:"word"`
	Count int    `json:"count"`
}

func (h *WordHandler) Handle(w http.ResponseWriter, r *http.Request) {
	d, err := ioutil.ReadAll(r.Body)
	if err != nil {
		utils.OutputBadRequest(w, err)
		return
	}

	// get words in string
	s := strings.Map(removePunctuation, string(d))
	words := strings.Fields(s)

	// build counting word data
	wordCountDict := getWordCountDict(words)

	topPopular := getTopPopularWords(wordCountDict, 10)

	utils.OutputData(w, 200, topPopular)
}

func getTopPopularWords(wordCountDict map[string]int, top int) []*WordCount {
	wordCounts := make([]*WordCount, 0, len(wordCountDict))
	for key, val := range wordCountDict {
		wordCounts = append(wordCounts, &WordCount{Word: key, Count: val})
	}

	sort.Slice(wordCounts, func(i, j int) bool {
		return wordCounts[i].Count > wordCounts[j].Count
	})

	if top > len(wordCounts) {
		return wordCounts
	} else {
		return wordCounts[:top]
	}
}

func removePunctuation(r rune) rune {
	if strings.ContainsRune(".,:;\"'`[]{}()", r) {
		return -1
	} else {
		return r
	}
}

func getWordCountDict(words []string) map[string]int {
	wm := make(map[string]int)
	for _, word := range words {
		if _, ok := wm[word]; ok {
			wm[word]++
		} else {
			wm[word] = 1
		}
	}

	return wm
}
