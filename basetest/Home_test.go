package basetest

import (
	"fmt"
	"github.com/pkg/errors"
	"unsafe"
)

var age int

// ------------------- 数据类型 -------------------//

/**
数字类型
uinit8
uint16
uint32
uint64
int8
int16
int32
int64

float32
float64
complex64		// 32位实数和虚数
complex128		// 64位实数和虚数

其他数字类型
byte		类似 uint8
rune		类似 int32
uint		32 或 64 位
int			32 或 64 位
uintptr		无符号整型，用于存放一个指针


*/

/**
派生类型
(a) 指针类型（Pointer）
(b) 数组类型
(c) 结构化类型(struct)
(d) Channel 类型
(e) 函数类型
(f) 切片类型
(g) 接口类型（interface）
(h) Map 类型
*/

// ------------------- 变量 -------------------//

// 这种因式分解关键字的写法一般用于声明全局变量
var (
	vname1 int
	vname2 int
)

// ------------------- 常量 -------------------//

// 显式
const b string = "abc"

// 隐式
const d = "abc"

//枚举
const (
	UNKNOWN = 0
	FEMALE  = 1
	MALE    = 2
	m       = "abc"
	n       = len(m)
	l       = unsafe.Sizeof(m)
)

// 常量可以用len(), cap()容量, unsafe.Sizeof()函数计算表达式的值。常量表达式中，函数必须是内置函数，否则编译不过：

// itoa
// 特殊常量，可以认为是一个可以被编译器修改的常量。
const (
	aa = iota //0
	ab = iota //1
	ac = iota //2
)

const (
	ai = 1 << iota
	aj = 3 << iota
	ak
	al
)

// ------------------- 数组 -------------------//

// 声明数组
// var variable_name [SIZE] variable_type

// 初始化数组
var balance = [5]float32{1000.0, 2.0, 3.4, 7.0, 50.0}

var balance2 = []float32{1000.0, 2.0, 3.4, 7.0, 50.0}

// ------------------- 指针 -------------------//
// & 取地址符
// 指针 ： 一个指针变量指向了一个值的内存地址。

// 指针的声明格式
// var var_name *var-type

// 指针的使用流程
// 定义指针
// 为指针变量赋值
// 访问指针变量中指向地址的值
//df

//eg：
func test() {
	var a int = 20
	var ip *int

	ip = &a
	fmt.Printf("a 变量的地址是: %x\n", &a)

	/* 指针变量的存储地址 */
	fmt.Printf("ip 变量储存的指针地址: %x\n", ip)

	/* 使用指针访问值 */
	fmt.Printf("*ip 变量的值: %d\n", *ip)
}

// ------------------- 结构体 -------------------//

// 定义结构体
/*
type struct_variable_type struct {
	member definition;
	member definition;
	...
	member definition;
}
*/

// 结构体赋值
/*
variable_name := structure_variable_type {value1, value2...valuen}
或
variable_name := structure_variable_type { key1: value1, key2: value2..., keyn: valuen}
*/

type Books struct {
	title   string
	author  string
	subject string
	book_id int
}

// ------------------- 切片 -------------------//
// Go 语言切片是对数组的抽象。

// 定义切片：你可以声明一个未指定大小的数组来定义切片
// var identifier []type

// 使用make()函数来创建切片:
/*
	var slice1 []type = make([]type, len)

	也可以简写为

	slice1 := make([]type, len)
*/

// ------------------- 范围 -------------------//
/*
Go 语言中 range 关键字用于 for 循环中迭代数组(array)、切片(slice)、通道(channel)或集合(map)的元素。
在数组和切片中它返回元素的索引和索引对应的值，在集合中返回 key-value 对的 key 值。
*/

// ------------------- 递归函数 -------------------//

// ------------------- 类型转换 -------------------//

// ------------------- 接口 -------------------//
/*
Go 语言提供了另外一种数据类型即接口，它把所有的具有共性的方法定义在一起，任何其他类型只要实现了这些方法就是实现了这个接口。
*/
// eg:
type Phone interface {
	call()
}

type NokiaPhone struct {
}

func (nokiaPhone NokiaPhone) call() {
	fmt.Println("I am Nokia, I can call you!")
}

type Iphone struct {
}

func (iphone Iphone) call() {
	fmt.Println("I am iPhone, I can call you!")
}

func test2() {
	var phone Phone
	phone = new(NokiaPhone)
	phone.call()

	phone = new(Iphone)
	phone.call()
}

// ------------------- 错误处理 -------------------//
// error类型是一个接口类型，这是它的定义：
/*
	type error interface {
		Error() string
	}
*/

func Sqrt(f float64) (float64, error) {
	if f < 0 {
		return 0, errors.New("math: square root of negative number")
	}
	return 0, nil
}

// ------------------- 并发 -------------------//
/*
Go 语言支持并发，我们只需要通过 go 关键字来开启 goroutine 即可。
goroutine 是轻量级线程，goroutine 的调度是由 Golang 运行时进行管理的。

语法格式：
go 函数名( 参数列表

例如：
go f(x,y,z)

开启一个新的 goroutine:
f(x,y,z)
*/

// ------------------- 通道（channel）-------------------//
/*
通道（channel）是用来传递数据的一个数据结构。
通道可用于两个 goroutine 之间通过传递一个指定类型的值来同步运行和通讯。
操作符 <- 用于指定通道的方向，发送或接收。
如果未指定方向，则为双向通道。
ch <- v    // 把 v 发送到通道 ch
v := <-ch  // 从 ch 接收数据
           // 并把值赋给 v
*/

/*
声明一个通道
ch := make(chan int)
*/
func sum(s []int, c chan int) {
	sum := 0
	for _, v := range s {
		sum += v
	}
	c <- sum // 把 sum 发送到通道 c
}

func main60() {
	s := []int{7, 2, 8, -9, 4, 0}

	c := make(chan int)
	go sum(s[:len(s)/2], c)
	go sum(s[len(s)/2:], c)
	x, y := <-c, <-c // 从通道 c 中接收

	fmt.Println(x, y, x+y)
}

/*
通道可以设置缓冲区，通过 make 的第二个参数指定缓冲区大小：
ch := make(chan int, 100)
*/

// ------------------- 遍历通道与关闭通道 -------------------//
/*
// 关闭通道
close(c)
*/

func main7() {

	println("Hello World")

	// 忽略的字段为 0 或 空
	fmt.Println(Books{title: "Go 语言", author: "www.runoob.com"})
}

func PrintTest() {
	fmt.Println("hello test")
}
