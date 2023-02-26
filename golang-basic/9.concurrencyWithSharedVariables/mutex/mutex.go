package bank

import "sync"

// 互斥锁
// 使用 sync 包的 Mutex 类型
var (
	// 按照惯例，被互斥量保护的变量声明应当紧贴在互斥量的声明之后
	mu 		sync.Mutex
	balance int
)

func Withdraw(amount int) bool {
	mu.Lock()
	defer mu.Unlock()
	// 在这里若是直接调用原来的 Deposit() 函数会导致运行两次 mu.Lock()，导致阻塞。
	// 但若在这个函数里不加上互斥锁，会导致再进行超额提款时，在某个瞬间余额会降到 0 以下，
	// 还没有恢复原始余额前若，有一个小额的提款，就会被拒绝。
	deposit(-amount)
	if balance < 0 {
		deposit(amount)
		return false // 余额不足
	}
	return true
}

func Deposit(amount int) {
	mu.Lock()
	// 在比较复杂的情况使用 defer 延迟执行，可以确保 Unlock 一定会执行
	defer mu.Unlock()
	// Lock 与 Unlock 之间称为临界区域，可以任意读取和修改共享变量
	deposit(amount)
}

// 封装即通过在程序中减少对数据结构的非预期交互，来保证数据结构中的不变量。
// 因为类似的原因，封装也可以用来保持并发中的不变性。
func deposit(amount int) { balance += amount }

// func Deposit(amount int) {
// 	mu.Lock()
// 	defer mu.Unlock()
// 	balance = balance + amount
// }

func Balance() int {
	b := balance
	return b
}
