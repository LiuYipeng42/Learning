package main

import (
	"bytes"
	"fmt"
	"strconv"
	"strings"
	"unicode/utf8"
)

func main() {
	string_()
	escapeChar()
	unicode()
}


func string_() {
	// 字符串
	s := "hello, world"
	fmt.Println(len(s))
	fmt.Println(s[0], s[7])
	fmt.Println(s[:5], s[7:], s[:])
	fmt.Println("goodbye" + s[5:])

	// 可以通过比较运算符进行字符串的比较
	// 比较运算按字节进行，结果服从本身的字典排序
	fmt.Println("abcd" > s)

	// 字符串变量可以重新赋值，但字符串不可变
	s1 := "hello"
	// 将 hello 变量的地址赋值给 t 变量，表示原字符串不变
	t := s1
	s1 += ", world"
	fmt.Println(t)
	fmt.Println(s1)
	// 编译出错
	// t[0] = 'a'

	// 字符串不可变使得复制任何长度字符串的开销都很低，
	// 字符串 s 及其子串 (s[i:j]) 可以共享数据

	// 关于字符串的操作都存在在 strings 包中
	fmt.Println(strings.Contains(s1, t))
}


func escapeChar() {
	// 转义字符
	// \ 后可以加上一个数字，用于表示一个字节大小的字符
	// 16 进制转义字符：\xhh，2 位 16 进制数
	// 8 进制转义字符：\ooo，3 位 8 进制数
	// 字符串支持转义字符
	t := "string\n"
	// 原生字符串不支持转义字符
	t1 := `string\n`
	fmt.Println(t, t1)

	// Unicode 16 位码点值转义：\uhhhh，4 位 16 进制数
	// Unicode 32 位码点值转义：\Uhhhhhhhh，8 位 16 进制数
	fmt.Println("世界")
	// 每次都将 \x 后的数字转义成一个字节数据，每三个字节数据正好转义成一个汉字
	// 因为 Go 的原文件总是以 UTF-8 编码，所以被转换成汉字
	fmt.Println("\xe4\xb8\x96\xe7\x95\x8c")
	fmt.Println("\u4e16\u754c")
	fmt.Println("\U00004e16\U0000754c")
}


func unicode() {
	s := "hello, 世界"
	fmt.Println(len(s))
	// go 将字符称为 rune
	// rune 类型是和 int32 等价的类型，通常用于表示一个 Unicode 码点
	fmt.Println(utf8.RuneCountInString(s))

	// 按照每个 Unicode 码点 进行遍历
	for i := 0; i < len(s); {
		r, size := utf8.DecodeRuneInString(s[i:])
		fmt.Println("%d\t %c\n", i, r)
		i += size
	}

	// 使用 range 也可以遍历
	for i, r := range s {
		fmt.Printf("%d\t%q\t%d\n", i, r, r)
	}

	s1 := "プログラム"
	// 以十六进制形式输出，并在每两个数位间插入空格
	fmt.Printf("% x\n", s1)
	// 将字符串强制转换成 rune 类型的数组
	// 此操作会分配新的字节数组，拷贝填入字符串含有的字节，并生成一个数组引用
	// 所以改变此字节数组的值不会改变原字符串的值
	r := []rune(s1)
	fmt.Printf("%x\n", r)
	// rune 类型的数组强制转换成字符串
	// 字节数组转换成字符串时也会产生一份副本，保证改变原数组时不会改变字符串
	fmt.Println(string(r))
}


func packages() {
	// 处理字符串的包：strings
	// 处理字节数组的包：bytes
	// 字符串与其他类型的转换：strconv
	// 判别文字符号值的类型（若判断字符串是否为数字）：unicode
}

func intsToString(values []int) string {
	// fmt.Sprint() 有类似的功能
	// 因为字符串是不可变的，所以每次在字符串后面加上一个字符时，会创建一个新的字符串，会浪费内存，
	// bytes.Buffer 类型可以避免此类问题，此类型起始为空，其大小会随着各种类型数据的写入而增长
	var buf bytes.Buffer
	buf.WriteByte('[')
	for i, v := range values {
		if i > 0 {
			buf.WriteString(", ")
		}
		fmt.Fprintf(&buf, "%d", v)
	}
	buf.WriteByte(']')
	return buf.String()
}

func stringAndNum() {

	// 字符串与数字之间的转换
	// 数字转字符串
	x := 123
	x1 := fmt.Sprintf("%d", x)
	x2 := strconv.Itoa(x)
	// 转换成其他进制的字符串
	x3 := strconv.FormatInt(int64(x), 2)
	fmt.Println(x1, x2, x3)

	// 字符串转数字
	if x, err := strconv.Atoi("123"); err==nil {
		fmt.Println(x)
	}
	// 转换成 10 进制 64 位的整数
	if x, err := strconv.ParseInt("123", 10, 64); err==nil {
		fmt.Println(x)
	}


}