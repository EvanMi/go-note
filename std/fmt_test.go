package fmttest

import (
	"fmt"
	"testing"
)

type User struct {
	Id int64
}

func TestPrintf1(t *testing.T) {
	user := &User{Id: 1}

	//原始值类型
	fmt.Printf("%v\n", user)  //值的默认格式表示
	fmt.Printf("%+v\n", user) //类似%v，但输出结构体时会添加字段名
	fmt.Printf("%#v\n", user) //值的Go语法表示
	fmt.Printf("%T\n", user)  //打印值的类型
	fmt.Print("%%\n")         //打印百分号
}

func TestPrintf2(t *testing.T) {
	//bool型
	fmt.Printf("%t\n", true)
}

//整数相关
// %b 表示为二进制
// %c 表示为unicode码所对应的字符
// %d 表示为十进制
// %o 表示为八进制
// %x 表示为十六进制小写(a~f)
// %X 表示为十六进制大写(A~F)
// %U 表示为Unicode格式
// %q 表示单引号括起来的go语法字符字面值，必要时会采用安全的转义表示

func TestPrintf3(t *testing.T) {
	n := 180
	fmt.Printf("%b\n", n)
	fmt.Printf("%c\n", n)
	fmt.Printf("%d\n", n)
	fmt.Printf("%o\n", n)
	fmt.Printf("%x\n", n)
	fmt.Printf("%X\n", n)
	fmt.Printf("%U\n", n)

	a := 96
	fmt.Printf("%q\n", a)
	fmt.Printf("%q\n", 0x4E2D)
}

//浮点数与复数
// %b 无小数部分、二进制指数的科学计数法
// %e 科学计数法
// %E 科学计数法
// %f 有小数部分但无指数部分
// %F 等价于%f
// %g 根据实际情况采用%e或%f格式，来获取到更简洁、准确的输出
// %G 根据实际情况采用%E或%F格式，来获取到更简洁、准确的输出

func TestPrintf4(t *testing.T) {
	f := 18.54
	fmt.Printf("%b\n", f)
	fmt.Printf("%e\n", f)
	fmt.Printf("%E\n", f)
	fmt.Printf("%f\n", f)
	fmt.Printf("%F\n", f)
	fmt.Printf("%g\n", f)
	fmt.Printf("%G\n", f)
}

//字符串相关
// %s直接输出字符串或者[]byte
// %q该值对应的双引号括起来的go语法字符串字面值，必要时会采用安全的转义表示
// %x每个字节用两字符十六进制数表示(a~f)
// %X每个字节用两字符十六进制数表示(A~F)

func TestPrintf5(t *testing.T) {
	s := "我是字符串"
	b := []byte{65, 66, 67}

	fmt.Printf("%s\n", s)
	fmt.Printf("%s\n", b)
	fmt.Printf("%q\n", s)
	fmt.Printf("%x\n", s)
	fmt.Printf("%X\n", s)
}

//指针相关
//%p 表示为十六进制，并加上前导的0x

//宽度标识符
// 宽度 通过一个紧跟在百分号后面的十进制数指定，如果未指定宽度，则表示值时除必须之外不做填充。
// 精度 通过（可选的）宽度后跟点号再跟十进制数指定。如果为指定精度，会使用默认精度，如果点号后没有跟数字，表示精度为0.

func TestPrintf6(t *testing.T) {
	n := 13.14

	fmt.Printf("%f\n", n)
	fmt.Printf("%10f\n", n)
	fmt.Printf("%10s\n", "我是字符串")
	fmt.Printf("%.2f\n", n)
	fmt.Printf("%10.2f\n", n)
	fmt.Printf("%10.f\n", n)
}

//其他flag
// + 总是输出数值的正负号；对%q(%+q)会生成全部是ASCII字符的输出（通过转义）。
// 空格 对数值，正数前加空格而负数前加负号；对字符串采用%x或%X时会给各打印的字节之间加空格
// - 在输出右边填充空白而不是默认的左边
// # 八进制数前加0，十六进制前加0x或0X,指针去掉前面的0x，对%q %U 会输出空格和单引号括起来的go字面量
// 0 使用0而不是空格填充，对于数值类型会把填充的0放在正负号后面

func TestPrintf7(t *testing.T) {
	s := "我是字符串"
	fmt.Printf("% d\n", -10)
	fmt.Printf("%+d\n", 10)
	fmt.Printf("%10s\n", s)
	fmt.Printf("%-10s\n", s)
	fmt.Printf("%-10.2f\n", 10.14)
	fmt.Printf("%010s\n", s)
}
