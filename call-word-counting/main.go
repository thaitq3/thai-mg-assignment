package main

import (
	"bufio"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
)

func main() {
	d, _ := postDataFromFile("GoLang_Test.txt")

	fmt.Println(d)
}

func postDataFromFile(filePath string) (string, error) {
	url := "http://localhost:8080/words"

	f, err := os.Open(filePath)
	if err != nil {
		fmt.Println("can't opened this file")
		return "", err
	}
	defer f.Close()

	req, err := http.NewRequest("POST", url, bufio.NewReader(f))
	if err != nil {
		return "", err
	}

	res, err := http.DefaultClient.Do(req)
	if err != nil {
		return "", err
	}

	defer res.Body.Close()
	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return "", err
	}

	return string(body), nil
}
