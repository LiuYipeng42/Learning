package main

import "sync"

// 延迟初始化
// 使用 sync 包的 Once 类型
var loadDataOnce sync.Once
var data map[string]string

func loadData() {
	data = make(map[string]string)
	data["a"] = "sdawdeaef..."
	data["b"] = "sdawdeaef..."
	data["c"] = "sdawdeaef..."
	data["d"] = "sdawdeaef..."
}

// 可能会发生数据竞态
// 一个 goroutine 刚运行完 make 使 data 由 nil 变为非空，map 还没有数据，
// 但另一个 goroutine 开始运行，发现 data 非空，访问数据后出错
// func Data(key string) string {
// 	if data == nil {
// 		loadData()
// 	}
// 	return data[key]
// }

func Data(key string) string {
	// Once 包含有一个 bool 变量和一个互斥量，
	// bool 变量记录函数是否执行，互斥量阻止多个 goroutine 同时初始化，
	// 执行过一次后 bool 变量变为 false，后续的调用相当于空操作
	loadDataOnce.Do(loadData)
	return data[key]
}

func main() {
	go Data("a")
	go Data("b")
}
