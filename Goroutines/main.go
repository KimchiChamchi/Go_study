package main

import (
	"fmt"
	"time"
)

func main() {
	asdf()

	// goroutines
	// 함수 앞에 go 를 붙이면 다음 함수와 동시에 실행되게 해줌
	// 작업을 순서대로만 하는것보다 속도가 올라가는 장점이 있음
	// 단, 작업도중 main함수가 종료되면 그냥 종료됨.
	// goroutines이 완료될때까지 기다려주지 않음.
	go sexyCount("main1")
	sexyCount2("main2")
}

func asdf() {
	fmt.Println(1)
	go sexyCount("asdf1")
	fmt.Println(2)
	fmt.Println(2)
	fmt.Println(2)
	fmt.Println(2)
	fmt.Println(2)
	fmt.Println(2)
	fmt.Println(2)

	sexyCount2("asdf2?")

}
func sexyCount(person string) {
	for i := 0; i < 10; i++ {
		fmt.Println(person, "은(는) 섹시함", i)
		time.Sleep(time.Second)
	}
}
func sexyCount2(person string) {
	for i := 0; i < 3; i++ {
		fmt.Println(person, "은(는) 섹시함", i)
		time.Sleep(time.Second)
	}
}
