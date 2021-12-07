package main

import (
	"fmt"

	"github.com/LeessangMin/gosu/dictionary/mydict"
)

func main() {
	// dictionary := mydict.Dictionary{}
	// dictionary["kimchi"] = "김치"
	// fmt.Println(dictionary)

	// // 단어 검색
	// dictionary := mydict.Dictionary{"first": "First word"}
	// definition, err := dictionary.Search("second")
	// if err != nil {
	// 	fmt.Println(err)
	// } else {
	// 	fmt.Println(definition)
	// }

	// // 단어 추가
	// dictionary := mydict.Dictionary{}
	// word := "안녕"
	// definition := "인사"
	// err := dictionary.Add(word, definition)
	// if err != nil {
	// 	fmt.Println(err)
	// }
	// hello, _ := dictionary.Search(word)
	// fmt.Println("찾는 단어 : ", word, " 단어뜻 : ", hello)
	// err2 := dictionary.Add(word, definition)
	// if err2 != nil {
	// 	fmt.Println(err2)
	// }

	// // 단어 업데이트
	// dictionary := mydict.Dictionary{}
	// baseWord := "안녕"
	// dictionary.Add(baseWord, "인사")
	// err := dictionary.Update(baseWord, "반가운 인사")
	// if err != nil {
	// 	fmt.Println(err)
	// }
	// word, _ := (dictionary.Search(baseWord))
	// fmt.Println(word)

	// 단어 삭제
	dictionary := mydict.Dictionary{}
	baseWord := "안녕"
	dictionary.Add(baseWord, "인사")
	dictionary.Search(baseWord)
	dictionary.Delete(baseWord)
	word, err := dictionary.Search(baseWord)
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(word)
	}
}
