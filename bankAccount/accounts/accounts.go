package accounts

// 계좌 구조체
// Account 앞을 대문자로 작성해야 다른곳에서 import 할 수 있음
type Account struct {
	// struct 안의 내용(Owner, Balance)도 소문자면 private라서 다른곳에서 가져갈 수 없음
	// 소문자 : private / 대문자 : bublic
	owner   string
	balance int
}

// 새 계좌 만들기 함수
func NewAccount(owner string) *Account {
	account := Account{owner: owner, balance: 0}
	return &account
}

// 계좌에 예금하는 메서드
//   (receiver) : 다른언어에서 메서드랑 동등한 관계
func (a Account) Deposit(amount int) {
	a.balance += amount
}

// 계좌 잔액
func (a Account) Balance() {
	return a.balance
}
