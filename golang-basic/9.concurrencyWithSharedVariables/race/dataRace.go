package main

import (
	"fmt"
)

var balance int

func Deposit(amount int) { balance = balance + amount }

func Balance() int { return balance }

func main() {

	// 当 Alice 在运行 Deposit 时，会先运行 balance + amount， 再给 balance 赋值
	// 若在运行 balance + amount 后，还没有赋值，调度器就切换到 Bob 线程，并执行完了 Deposit(100)，
	// balance 的值变为 100，之后回到 Alice 的 Deposit 操作中，将 balance + amount 的值赋值给 balance
	// balance 的值从 100 变为 200，出现错误

	// Alice
	go func() {
		Deposit(200)
		fmt.Println("=", Balance())
	}()

	// Bob
	go Deposit(100)

	// 此种情况叫做 数据竞态
	// 定义：
	//     多个线程对于同一个变量，同时地、进行读/写操作的现象并且至少有一个线程进行写操作。
	//     （也就是说，如果所有线程都是只进行读操作，那么将不构成数据争用）
	// 后果：
	//     如果发生了数据争用，读取该变量时得到的值将变得不可知，使得该多线程程序的运行结果将完全不可预测，可能直接崩溃。
	
	// e.g.
	// a := 42
	// if a == 42 {
	// 	fmt.Printf("answer = %d", a)
	// }
	// 若一个线程发现 a 符合条件，进行输出，再输出前另一个线程改变了 a 的值，最后就会输出错误的结果
}
