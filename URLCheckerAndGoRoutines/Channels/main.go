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
	// 변수c를 채널 string타입으로 만듦
	c := make(chan string)
	people := []string{"상민", "길동", "둘리", "도우너", "또치"}
	for _, person := range people {
		// 함수에 사람 이름이랑 채널 넣으면 isSexy함수에서
		// 5초뒤 채널을 통해 (person + "is sexy")를 돌려줌
		go isSexy(person, c)
	}

	fmt.Println("메시지 기다리는중")
	for i := 1; i < len(people)+1; i++ {
		time.Sleep(time.Millisecond * 500)
		fmt.Println(i, "번째 메시지:", <-c)
	}

}

//                    인자를 채널 string타입으로 받음
func isSexy(person string, channel chan string) {
	time.Sleep(time.Second * 5)
	// fmt.Println(person)
	// channel에 (person + "is sexy") 보냄
	channel <- person + "is sexy"
}
