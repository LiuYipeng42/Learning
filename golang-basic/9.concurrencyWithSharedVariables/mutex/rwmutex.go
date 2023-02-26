package bank

import "sync"

// 互斥锁
// 使用 sync 包的 RWMutex 类型
var (
	mu      sync.RWMutex // 有读锁和写锁，即 RLock 与 Lock
	balance int
)

func Withdraw(amount int) bool {
	mu.Lock()
	defer mu.Unlock()
	deposit(-amount)
	if balance < 0 {
		deposit(amount)
		return false // 余额不足
	}
	return true
}

func Deposit(amount int) {
	mu.Lock()
	defer mu.Unlock()
	deposit(amount)
}

func deposit(amount int) { balance += amount }

// 此函数只是读取变量的状态，所以多个 Balance 请求其实可以安全地并发运行
// sync.RWMutex 提供一个 多读单写锁，可以提供此种功能
func Balance() int {
	// RLock 仅可用于在临界区域内对共享变量无写操作的情形
	// 1. 读写锁的读锁可以重入，在已经有读锁的情况下，可以任意加读锁。
	// 2. 在读锁没有全部解锁的情况下，写操作会阻塞直到所有读锁解锁。
	// 3. 写锁定的情况下，其他协程的读写都会被阻塞，直到写锁解锁。
	mu.RLock()
	defer mu.RUnlock()
	b := balance
	return b
}
