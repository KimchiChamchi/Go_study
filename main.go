/////////////////////////////변수 선언, 할당, 타입에 대해
// package main

// // 축약형(:=)은 func 안에서만 사용 가능
// // asdf := "pk"
// func main() {
// 	// go 언어에서는 변수의 타입(string, int 등..)도 적어주어야 함
// 	// var name string = "sm"
// 	// 대신 아래와 같이 var와 타입을 생략할 수 있음
// 	// go가 알아서 타입을 찾아서 지정해줌
// 	name := "sm"
// 	name = "대홍단감자"
// 	println(name)
// }
///////////////////////////////////////////////////////
///////////////////////////////다른곳에 있는 패키지 불러오기
// package main

// import (
// 	"fmt"

// 	"github.com/LeessangMin/gosu/something"
// )

// func main() {
// 	fmt.Println("Hello world")
// 	something.SayHello()
// }
///////////////////////////////////////////////////////
///////////////////////////////함수 선언에 대해
// package main
// import "fmt"

// //a,b 둘다 타입이 같으면 아래처럼 통합 가능
// //                      반환될 값의 타입도 반드시 있어야 함
// func multiply(a, b int) int {
// 	return a * b
// }

// func main()  {
// fmt.Println(multiply(3, 9))
// }
//////////////////////////////////////////////////////////////////////
//////////////////////////// 함수 사용방법
// package main

// import (
// 	"fmt"
// 	"strings"
// )

// func lenAndUpper(name string) (int, string) {
// 	//     len() : ()안 변수값의 길이를 반환
// 	//                strings.ToUpper() : ()안 문자열을 대문자로 반환
// 	return len(name), strings.ToUpper(name)
// }

// //                  ...으로 여러 값을 받을 수 잇음
// func repeatMe(words ...string) {
// 	fmt.Println(words)
// }
// func main() {
// 	// _ 로 작성하면 해당 반환값은 무시됨
// 	// totalLenght, _ := lenAndUpper("kimchi")
// 	totalLenght, upperName := lenAndUpper("kimchi")
// 	fmt.Println(totalLenght, upperName)

// 	repeatMe("qwer", "asdf", "zxcv", "ggg")
// }
//////////////////////////////////////////////////////////////////
////////////////////////////////////// "naked" return
// package main

// import (
// 	"fmt"
// 	"strings"
// )

// //                            반환할 값들에 대해 미리 정의해두면
// func lenAndUpper(name string) (length int, uppercase string) {
// 	// length 와 uppercase 는 위에 정의한 이름과 같아야 함
// 	length = len(name)
// 	uppercase = strings.ToUpper(name)
// 	return // return 뒤에 반환할 값들을 생략할 수 있다
// }

// func main() {
// 	totalLenght, up := lenAndUpper("kimchi")
// 	fmt.Println(totalLenght, up)
// }
//////////////////////////////////////////////////////////////////////
////////////////////////////////////// "defer"
// package main

// import (
// 	"fmt"
// 	"strings"
// )

// func lenAndUpper(name string) (length int, uppercase string) {
// 	// defer 는 return 후에 실행된다
// 	defer fmt.Println("나 다 됐엉!")
// 	length = len(name)
// 	uppercase = strings.ToUpper(name)
// 	return
// }

// func main() {
// 	totalLenght, up := lenAndUpper("kimchi")
// 	fmt.Println(totalLenght, up)
// 	defer fmt.Println("나 다 됐엉!")
// }

//////////////////////////////////////////////////////////////////////
////////////////////////////////////// "range"
package main

import "fmt"

////////////// for 와 forr(range) 의 차이////////////
// for 는 i(인덱스) 초기값을 정하고 매회 i값을 증감시켜 특정 회수만큼 반복
// for i := 0; i < count; i++ {}
////////////////////////////////////////////////////
// forr 는 특정 함수나 배열 등을 넣어 그 안의 것들 개수만큼 반복
// 몇번 돌린건지 인덱스를 반환 받을 수도 있음
// for _, v := range v {}

// range 는 들어온 값의 개수만큼 반복문을 돌려준다.
// 아래의 경우 superAdd() 에 6개의 int 가 들어갔고
// 6번 돌려서 6번 반환한 것
func superAdd(numbers ...int) int {
	// // 기존 for문으로 작성 시
	// for i := 0; i < len(numbers); i++ {
	// 	fmt.Println(numbers[i])
	// }

	// // _, 가 있으면 인덱스를 무시
	// for _, number := range numbers {
	// 	fmt.Println(number)
	// }
	// // _, 가 없으면 range 는 인덱스값을 반환해줌
	// for number := range numbers {
	// 	fmt.Println(number)
	// }
	// return 1

	// range 사용하고 코드량 줄이고 합계 만들기
	total := 0
	for _, number := range numbers {
		total += number
	}
	return total
}
func main() {
	// superAdd(3, 5, 8, 7, 1, 9)
	result := superAdd(3, 5, 8, 7, 1, 9)
	fmt.Println(result)
}
