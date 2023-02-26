package main

import "fmt"

// 切片
// 与 python 中的类似

func main() {
	slice()
	append_()
	inPlace()
}

func slice() {
	// 切片有三个属性：指针，长度和容量
	// 指针：指向数组中切片的第一个元素
	// 长度：切片中的元素个数，不能超过切片的容量
	// 容量：切片的起始元素到底层数组的最后一个元素间的元素个数

	// 此数组成为切片的底层数组，改变此数组，所属切片中对应的内容也会改变
	// 反过来改变切片的内容也会改变底层数组的内容，切片相当于一个指针
	// python 中的切片并不会
	months := [...]string{1: "January", 2: "February", 3: "March", 4: "April",
		5: "May", 6: "June", 7: "July", 8: "August",
		9: "September", 10: "October", 11: "November", 12: "December"}

	// low:high:max
	// low: 切片第一个元素在数组中的索引
	// high: 切片最后一个元素在数组中的索引
	// max: 切片最大容量在数组中的索引，此项可以省略，cap = max - low
	summer := months[6:9:11]

	fmt.Println(summer)
	months[6] = "六月"
	fmt.Println(summer)

	// reslice 重切片，即在切片上创建一个切片，
	// 可以对原切片进行长度的收缩和拓展，新的切片与原切片共用同一个底层数组
	// 在 summer 切片的容量内拓展了切片的长度，拓展的长度不能超过容量
	endlessSummer := summer[:5]
	fmt.Println(endlessSummer)
	// 对切片的长度进行收缩
	finateSummer := summer[:3]
	fmt.Println(finateSummer)

	// 反转数组，因为是切片，所以是一个指针，
	// 传入函数时，只会复制一个指针，不会复制整个数组
	// 所以在函数内修改数组可以改变原来的数组
	reverse(months[:])

	// 声明一个有固定长度的数组
	a := [...]int{0, 1, 2, 3, 4, 5}
	// 创建一个切片，容量为整个数组
	b := []int{0, 1, 2, 3, 4, 5}
	fmt.Println(a, b)
	// 切片不能使用 == 或 != 来进行比较
	// 但切片可以与 nil 进行比较
	fmt.Println(summer == nil)

	c := make([]int, 2)   // 创建一个长度和容量都为 2 的切片
	c = make([]int, 2, 4) // 创建一个长度为 2，容量为 4 的切片
	fmt.Println(c)

}

func append_() {
	// 可以使用 append 函数将元素追加到切片后面

	var runes []rune
	for _, r := range "Hello, 世界" {
		runes = append(runes, r)
	}

	// fmt.Println([]rune("Hello, 世界") )
	fmt.Printf("%q\n", runes)

	a := [...]int{0, 1, 2, 3, 4}
	// appendInt(a[:3], 5)
	// 使用 append 函数会用添加的数覆盖原来在底层数组的数
	fmt.Println(append(a[:3], 5))
	fmt.Println(a)

	// 因为在容量不够时会创建一个更大的底层数组，并进行数据复制，且判断的策略较为复杂
	// 所以不能保证原始的切片和调用 append 后的结果切片指向
	// 同一个底层数组，也不能确保指向不同的底层数组，
	// 所以需要对原始切片的变量进行重新赋值，赋值时 len 和 cap 会随着改变
	b := a[:3]
	b = append(b, 5)
	b = append(b, 6, 7)
	b = append(b, a[:2]...)
}

func inPlace() {
	// 切片就地修改，修改后底层数组不变
	data := []string{"one", "", "three"}
	data = nonempty(data)
	fmt.Println(data)

	// 栈操作
	var stack []int
	stack = append(stack, 1)   // push
	top := stack[len(stack)-1] // 栈顶
	fmt.Println(top)
	stack = stack[:len(stack)-1] // pop

	// 利用 copy 函数来删除某一个元素
	s := []int{5, 6, 7, 8, 9}
	fmt.Println(remove(s, 2))
	
}

func reverse(s []string) {
	for i, j := 0, len(s)-1; i < j; i, j = i+1, j-1 {
		s[i], s[j] = s[j], s[i]
	}
}

func appendInt(x []int, y int) []int {
	var z []int
	zlen := len(x) + 1
	if zlen <= cap(x) {
		// 将切片向后拓展一位
		z = x[:zlen]
	} else {
		zcap := zlen
		if zcap < 2*len(x) {
			zcap = 2 * len(x)
		}
		z = make([]int, zlen, zcap)
		copy(z, x)
		// 此时 z 与 x 的底层数组并不相同
	}
	z[len(x)] = y
	return z
}

func nonempty(strings []string) []string {
	out := strings[:0] // 与原始切片相同底层数组的切片，且长度为 0
	for _, s := range strings {
		if s != "" {
			out = append(out, s)
		}
	}

	return out
}

func remove(slice []int, i int) []int {
	// 删除切片中的第 i 个元素
	copy(slice[i:], slice[i+1:])
	return slice[:len(slice)-1]
}
