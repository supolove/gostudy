package basetest

import (
	"fmt"
	"reflect"
	"strings"
	"testing"
	"unsafe"
)

/*
unsafe.Pointer
这个类型比较重要，它是实现定位欲读写的内存的基础。官方文档对该类型有四个重要描述：
（1）任何类型的指针都可以被转化为Pointer
（2）Pointer可以被转化为任何类型的指针
（3）uintptr可以被转化为Pointer
（4）Pointer可以被转化为uintptr

一个普通的T类型指针可以被转化为unsafe.Pointer类型指针，
并且一个unsafe.Pointer类型指针也可以被转回普通的指针，
被转回普通的指针类型并不需要和原始的T类型相同。

*/

/** unsafe start */
func pointer1() {
	a := [4]int{0, 1, 2, 3}
	p1 := unsafe.Pointer(&a[1])
	fmt.Println(uintptr(p1))
	fmt.Println(unsafe.Sizeof(a[0]))
	p3 := unsafe.Pointer(uintptr(p1) + 2*unsafe.Sizeof(a[0]))
	*(*int)(p3) = 6
	fmt.Println("a = ", a)

	fmt.Println(strings.Repeat("~", 20))

	type Person struct {
		name   string
		age    int
		gender bool
	}

	who := Person{"John", 30, true}
	pp := unsafe.Pointer(&who)
	pname := (*string)(unsafe.Pointer(uintptr(pp) + unsafe.Offsetof(who.name)))
	page := (*int)(unsafe.Pointer(uintptr(pp) + unsafe.Offsetof(who.age)))
	pgender := (*bool)(unsafe.Pointer(uintptr(pp) + unsafe.Offsetof(who.gender)))

	*pname = "Alice"
	*page = 28
	*pgender = false

	fmt.Println(who)
}

func float64bits(f float64) uint64 {
	fmt.Println(reflect.TypeOf(unsafe.Pointer(&f)))            //unsafe.Pointer
	fmt.Println(reflect.TypeOf((*uint64)(unsafe.Pointer(&f)))) //*uint64
	return *(*uint64)(unsafe.Pointer(&f))
}

func pointer2() {
	fmt.Printf("%#016x\n", float64bits(1.0)) // "0x3ff0000000000000"
}

func pointer3() {
	v1 := uint(12)
	v2 := int(12)

	fmt.Println(reflect.TypeOf(v1)) //uint
	fmt.Println(reflect.TypeOf(v2)) //int

	fmt.Println(reflect.TypeOf(&v1)) //*uint
	fmt.Println(reflect.TypeOf(&v2)) //*int

	p := &v1

	//两个变量的类型不同,不能赋值
	//p = &v2 //cannot use &v2 (type *int) as type *uint in assignment
	p = (*uint)(unsafe.Pointer(&v2)) //使用unsafe.Pointer进行类型的转换

	fmt.Println(reflect.TypeOf(p)) // *unit
}

/** unsafe end */

func TestPointer(t *testing.T) {
	//pointer1()
	//pointer2()
	pointer3()
}
