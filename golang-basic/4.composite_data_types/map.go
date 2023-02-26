package main

import (
	"fmt"
	"sort"
)

func main() {

	// var ages map[string]int
	// ages := make(map[string]int)
	ages := map[string]int{
		"charlie": 34,
	}

	// 可以使用 == 来进行比较的数据类型都可以作为键
	test := map[[2]int]int{
		{1, 2}: 1,
	}
	fmt.Println(test)

	ages["alice"] = 31
	ages["charlie"] = 42
	fmt.Println(ages["charlie"])
	// 可以使用内置函数 delete 来删除一个键值对
	delete(ages, "alice")

	// 可以删除和访问不存在的键值对，不存在的键值对的值默认是 0
	delete(ages, "nobody")
	fmt.Println(ages["nobody"])

	ages["alice"] += 1
	ages["nobody"]++
	fmt.Println(ages["nobody"])

	// map 中元素的迭代顺序是不固定的，不同的实现方法会使用不同的散列算法，
	// 得到不同的元素顺序，在写代码时，要将迭代顺序认为是随机的
	for name, age := range ages {
		fmt.Printf("%s\t%d\n", name, age)
	}

	// 按照字典顺序遍历 map
	// 首先将键进行排序，然后按照这个顺序输出
	names := make([]string, 0, len(ages))
	for name := range ages {
		names = append(names, name)
	}
	sort.Strings(names)
	for _, name := range names {
		fmt.Printf("%s\t%d\n", name, ages[name])
	}

	var none map[string]int
	fmt.Println(none == nil)
	fmt.Println(len(none) == 0)

	// 只是声明了一个变量，但并没有赋值，所以为 nil
	// 直接进行赋值操作时会报错，但可以进行访问
	// 所以设置元素之前，必须初始化 map
	fmt.Println(none["a"])
	// none["a"] = 42

	// 当 map 中的值可以为 0 时，
	// 可以用如下方法判断一个键值对是否在 map 中
	if age, ok := ages["bob"]; ok {
		fmt.Println(age)
	} else {
		fmt.Println("not exist")
	}

	// map 不可以直接使用 == 进行比较
	equal(map[string]int{"A": 0}, map[string]int{"B": 42})

	// 若想将不可使用 == 进行比较的数据结构作为键
	// 则可以将先将此数据结构的变量变为字符串后，再将字符串作为键即可
	// 比如切片
	Add([]int{1, 2})
	Add([]int{1, 2})
	fmt.Println(Count([]int{1, 2}))

	// map 的值也可以是复合数据类型，如 map 和 slice
	// 创建一个存储图的边的 map
	graph := make(map[string]map[string]bool)
	addEdge(graph, "a", "b")
	addEdge(graph, "b", "c")
	addEdge(graph, "c", "a")
	fmt.Println(hasEdge(graph, "a", "b"))

}

func addEdge(graph map[string]map[string]bool, from, to string){
	edges := graph[from]
	if edges == nil {
		edges = make(map[string]bool)
		graph[from] = edges
	}
	edges[to] = true
}

func hasEdge(graph map[string]map[string]bool, from, to string) bool {
	return graph[from][to]
}

var m = make(map[string]int)

func Add(list []int) {
	m[toString(list)] ++
}

func Count(list []int) int {
	return m[toString(list)]
}

func toString(list []int) string {
	return fmt.Sprintf("%q", list)
}


func equal(x, y map[string]int) bool {
	if len(x) != len(y) {
		return false
	}
	for k, xv := range x {
		// 先根据键获取值并判断是否存在，再比较是否相同
		// 不能直接进行 xv != y[k] 的比较，因为即使元素不存在也会被作为 0 值
		if yv, ok := y[k]; !ok || yv != xv {
			return false
		}
	}
	return true
}
