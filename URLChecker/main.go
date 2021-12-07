package main

import (
	"errors"
	"fmt"
	"net/http"
)

var errRequestFailed = errors.New("요청실패")

func main() {
	//map 은 초기화 해야 사용가능하고 제일 뒤에 {}를 붙이거나
	//make() 함수로 감싸면 초기화 된다
	results := map[string]string{}
	// results2 := make(map[string]string)
	urls := []string{
		"https://www.naver.com/",
		"https://www.daum.net/",
		"https://www.facebook.com/",
		"https://www.instagram.com/",
		"https://www.airbnb.com/",
		"https://www.google.com/",
		"https://www.amazon.com/",
		"https://www.reddit.com/",
		"https://www.google.com/",
		"https://soundcloud.com/",
		"https://academy.nomadcoders.co/",
	}

	for _, url := range urls {
		result := "성공"
		err := hitURL(url)
		if err != nil {
			result = "실패!"
		}
		results[url] = result
	}
	for url, result := range results {
		fmt.Println(url, result)
	}
}

func hitURL(url string) error {
	fmt.Println("Checking : ", url)
	resp, err := http.Get(url)
	if err != nil || resp.StatusCode >= 400 {
		fmt.Println(err, resp.StatusCode)
		return errRequestFailed
	}
	return nil
}
