package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"
)

var urls []string = []string{
	"https://golang.org/",
	"https://golang.org/doc/",
	"https://golang.org/pkg/compress/",
	"https://golang.org/pkg/compress/gzip/",
	"https://golang.org/pkg/crypto/md5/"}

func getBody(url string) (string, error) {
	resp, err := http.Get(url)
	if err != nil {
		return "", err
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	return string(body), nil
}

func main() {
	totalCount := 0
	for _, url := range urls {
		body, err := getBody(url)
		if err != nil {
			fmt.Println(url, ":", err)
			continue
		}
		count := strings.Count(body, "Go")
		fmt.Println(url, ":", count)
		totalCount += count
	}
	fmt.Println("Total count :", totalCount)
}
