package main

import "fmt"

var deposits = make(chan int) // 发送存款余额
var balances = make(chan int) // 接收余额


// 使用通道来向 监控goroutine 发送更新数据或查询数据的请求
func Deposit(amount int) { deposits <- amount }
func Balance() int       { return <-balances }

// 通过禁止多个 goroutine 访问同一个变量，来避免数据竞态
// 此 goroutine 称为 监控goroutine
func teller() {
	// 将 balance 变量限制在一个 goroutine 中
	var balance int
	for {
		select {
		// 每次循环只会从通道中接收一个数据，对 balance 进行更新
		case amount := <-deposits:
			balance += amount
		// 若 balances 通道为空，就将 balance 的值放入通道
		case balances <- balance:
		}
	}
}

func init() {
	go teller()
}

func main() {
	done := make(chan struct{})

	// Alice
	go func() {
		Deposit(200)
		fmt.Println("=", Balance())
		done <- struct{}{}
	}()

	// Bob
	go func() {
		Deposit(100)
		done <- struct{}{}
	}()

	// Wait for both transactions.
	<-done
	<-done

	if got, want := Balance(), 300; got != want {
		fmt.Printf("Balance = %d, want %d\n", got, want)
	}
}