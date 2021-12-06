package main

import (
	"fmt"

	"github.com/LeessangMin/gosu/bankAccount/accounts"
)

func main() {
	// public 이라 아무렇게나 가져다가 사용할 수 있으므로 private로 만들고
	// 또 다른 함수를 만들어 함수에 지정된 대로 제어
	// account := banking.Account{Owner: "sangmin", Balance: 1000}
	// account2 := banking.Account{Owner: "sangmin"}
	// fmt.Println(account)
	// fmt.Println(account2)

	account := accounts.NewAccount("sangmin")
	fmt.Println(account)

	account.Deposit(10)
	fmt.Println(account.Balance())
}
