package tempconv

import "fmt"


// type 声明定义一个新的命名类型，它和某个已有类型使用同样的底层类型
// 类似于 c 语言中的 typedef
type Celsius float64
type Fahrenheit float64

const (
	// 标识符以大写字母开头，表明可以在包外被访问
	AbsoluteZeroC Celsius = -273.15
	FreezingC     Celsius = 0
	BoilingC      Celsius = 100
)


// 名字为 String 的方法关联到 Celsius 类型，返回 c 变量的数字值
func (c Celsius) Srtring() string {
	return fmt.Sprintf("%g C", c)
}

func (c Fahrenheit) Srtring() string {
	return fmt.Sprintf("%g F", c)
}

