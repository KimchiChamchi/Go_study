###goroutines

goroutines 은 동시작업을 할 수 있게 해주어 작업속도가 훨씬 빠르게 만들어 준다.

실행할 함수 앞에 "go " 만 추가로 붙여주면 다른작업과 동시에 실행된다.

단, goroutines으로 실행되고 있는 함수가 아직 실행중이어도  
main 함수가 종료되면 goroutines도 종료된다.

강의를 보며 함수가 어느시점에 동작하는지 확인해보기위해  
goroutines 과 다음 함수 사이에 fmt.println을 잔뜩 끼워넣어봤으나 fmt.println가 전부 먼저 실행되었다.  
아마 goroutines 다음에 있는 함수가 실행될 때 같이 실행되는것 같다.  
또 goroutines을 사용중인 함수가 종료되면 goroutines도 종료되는게 아닐까 해서 main함수 안에 또다른 A함수를 만들고 그 안에서 goroutines을 사용해보았으나 A함수가 종료되어도 main함수가 종료되지 않은 상태면 goroutines은 계속 동작했다.
