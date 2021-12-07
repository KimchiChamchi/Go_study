package main

import (
	"fmt"
	"time"
)

// func main() {
// 	channel := make(chan bool)
// 	people := [2]string{"상민", "길동"}
// 	for _, person := range people {
// 		go isSexy(person, channel)
// 	}
// 	// 아래 isSexy함수에서 보낸 true를 담은 channel을
// 	// // result에 담음
// 	// result := <-channel
// 	// fmt.Println(result)

//  // "<-" 는 blocking operation
// 	// 위에서 사용한 goroutines은 2개인데 <-channel 을 세개 부르면 deadlock 에러
// 	fmt.Println(<-channel)
// 	fmt.Println(<-channel)
// 	fmt.Println(<-channel)

// }

// func isSexy(person string, channel chan bool) {
// 	time.Sleep(time.Second * 5)
// 	// channel에 true 를 보냄
// 	channel <- true
// }

func main() {
	c := make(chan string)
	people := []string{"상민", "길동", "둘리", "도우너", "또치"}
	for _, person := range people {
		go isSexy(person, c)
	}
	fmt.Println("메시지 기다리는중")
	for i := 1; i < len(people)+1; i++ {
		fmt.Println(i, "번째 메시지:", <-c)
	}

}

func isSexy(person string, channel chan string) {
	time.Sleep(time.Second * 5)
	// fmt.Println(person)
	// channel에 (person+ "is sexy") 보냄
	channel <- person + "is sexy"
}
