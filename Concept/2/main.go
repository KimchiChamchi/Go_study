/////////////////////////////// if 조건문
// package main

// import "fmt"

// func CanIDrink(age int) bool {
// 	// if 의 조건 안에서 바로 변수선언을 하고 사용할 수 있다.
// 	// 한국나이 := 만나이+2; 한국나이가 19세 미만이면...
// 	if koreanAge := age + 2; koreanAge < 19 {
// 		return false
// 	}
// 	//else 생략가능
// 	//if 후 "return" 하면 "else {return}" 과 같다
// 	return true

// }

// func main() {
// 	if CanIDrink(18) {
// 		fmt.Println("넌 술 마셔도 됑")
// 	} else {
// 		fmt.Println("넌 술 마시면 안됑")
// 	}
// }
////////////////////////////////////////////////////////////////////
/////////////////////////////////////////////// switch
// package main

// import "fmt"

// func CanIDrink(age int) bool {
// 	// switch 에서도 변수 바로 선언하여 사용 가능
// 	switch koreanAge := age + 2; {
// 	case koreanAge > 18:
// 		return true
// 	case koreanAge < 18:
// 		return false
// 	case koreanAge == 18:
// 		return false
// 	}
// 	return false
// }
// func main() {
// 	fmt.Println(CanIDrink(19))
// }
/////////////////////////////////////////////////////////////////////
////////////////////////////////////////////// pointer
// package main

// import "fmt"

// func main() {
// 	a := 2
// 	b := a
// 	fmt.Println(a, b)
// 	//&을 붙이면 해당값의 메모리 주소를 불러온다
// 	fmt.Println(&a, &b)
// 	a1 := 2
// 	b1 := &a1
// 	fmt.Println(&a1, b1) // &a1과 b1은 같다 (b1에 a1의 메모리 주소를 넣었기때문)
// 	// *b1 은 b1이 가리키고 있는 메모리 주소에 들어있는 값을 의미함
// 	fmt.Println(a1, *b1) // a1과 *b1 또한 같다 (b1은 a1의 메모리 주소이고)
// 	//                                (a1의 메모리 주소 안에 2가 있기때문)
// 	a1 = 5
// 	fmt.Println(a1, *b1) // b1은 a1의 메모리주소를 가지고 있기 때문에
// 	//                      a1의 값이 바뀌면 *b1도 똑같이 a1의 값으로 바뀐다
// 	*b1 = 10             //             거꾸로 b1을 이용하여 값을 바꾸는것도 가능하다
// 	fmt.Println(a1, *b1) // b1에 들어있는 메모리주소에 있는 값을 변경한것이기 때문에
// 	//                      이는 곧 a1의 메모리주소에 있는 값이니 a1이 변경되는것이다
// }
/////////////////////////////////////////////////////////////////////
////////////////////////////////////////////// 배열, slice, append
// package main

// import "fmt"

// func main() {
// 	// go에서는 배열의 길이를 제한해서 사용하기 때문에
// 	names := [5]string{"ㅁ", "ㄴ", "ㅇ", "ㄹ"}
// 	fmt.Println(names)

// 	// names[5] = "ㅎ2" // 여섯번째 배열은 넣을 수 없음

// 	// [] : slice라고 함. 배열에 제한 없이 쓸 수 있다.
// 	names2 := []string{"ㅂ", "ㅈ", "ㄷ", "ㄱ"}
// 	// slice에 배열을 추가하려면 append를 사용한다.
// 	// append는 원본을 수정하지 않고 새로운 slice를 반환함
// 	names2 = append(names2, "ㅅ")
// 	fmt.Println(names2)
// }
/////////////////////////////////////////////////////////////////////
////////////////////////////////////////////// map
// package main

// import "fmt"

// func main() {
// 	// key 와 value 타입을 지정하고 그에 맞는 객체를 넣음
// 	//변수   는 맵  [key]  value { key  : value,  key  :value}
// 	sangmin := map[string]string{"name": "쌍민", "나이": "99"}
// 	fmt.Println(sangmin)
// 	for key, value := range sangmin {
// 		fmt.Println(key, value)
// 	}
// }
/////////////////////////////////////////////////////////////////////
////////////////////////////////////////////// struct 구조체
// go에는 메서드가 없다
package main

import "fmt"

// 구조체 선언을 해두면 그에 맞게 내용을 담을 수 있다
type person struct {
	name    string
	age     int
	favFood []string
}

func main() {
	favFood := []string{"kimchi", "chamchi"}

	// 순서대로 넣으면 순서에 맞게 들어간다.
	// 그런데 이건 struct가 어떻게 생겼는지 한눈에 보기 어렵다
	// sangmin := person{"sangmin", 99, favFood}

	// 그래서 아래처럼 어떤게 어떤값인지 같이 적어줄 수 있다.
	sangmin := person{name: "sangmin", age: 99, favFood: favFood}
	fmt.Println(sangmin)
	fmt.Println(sangmin.name)
	fmt.Println(sangmin.age)
	fmt.Println(sangmin.favFood[1])
}
