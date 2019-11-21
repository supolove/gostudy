package basetest

import (
	"fmt"
	"reflect"
	"sort"
	"testing"
	"unsafe"
)

// 数组
func shuzu1()  {
	// 定义
	var a [3]int                    // 定义长度为3的int型数组, 元素全部为0
	var b = [...]int{1, 2, 3}       // 定义长度为3的int型数组, 元素为 1, 2, 3
	var c = [...]int{2: 3, 1: 2}    // 定义长度为3的int型数组, 元素为 0, 2, 3
	var d = [...]int{1, 2, 4: 5, 6} // 定义长度为6的int型数组, 元素为 1, 2, 0, 0, 5, 6

	fmt.Println(a)
	fmt.Println(b)
	fmt.Println(c)
	fmt.Println(d)

	//第一种方式是定义一个数组变量的最基本的方式，数组的长度明确指定，数组中的每个元素都以零值初始化。
	//第二种方式定义数组，可以在定义的时候顺序指定全部元素的初始化值，数组的长度根据初始化元素的数目自动计算。
	//第三种方式是以索引的方式来初始化数组的元素，因此元素的初始化值出现顺序比较随意。
	// 这种初始化方式和map[int]Type类型的初始化语法类似。数组的长度以出现的最大的索引为准，没有明确初始化的元素依然用0值初始化。
	//第四种方式是混合了第二种和第三种的初始化方式，前面两个元素采用顺序初始化，
	// 第三第四个元素零值初始化，第五个元素通过索引初始化，最后一个元素跟在前面的第五个元素之后采用顺序初始化。
}

// 切片
func qiepian1(){
	var (
		a []int               // nil切片, 和 nil 相等, 一般用来表示一个不存在的切片
		b = []int{}           // 空切片, 和 nil 不相等, 一般用来表示一个空的集合
		c = []int{1, 2, 3}    // 有3个元素的切片, len和cap都为3
		d = c[:2]             // 有2个元素的切片, len为2, cap为3
		e = c[0:2:cap(c)]     // 有2个元素的切片, len为2, cap为3
		f = c[:0]             // 有0个元素的切片, len为0, cap为3
		g = make([]int, 3)    // 有3个元素的切片, len和cap都为3
		h = make([]int, 2, 3) // 有2个元素的切片, len为2, cap为3
		i = make([]int, 0, 3) // 有0个元素的切片, len为0, cap为3
	)
	fmt.Println(a)
	fmt.Println(b)
	fmt.Println(c)
	fmt.Println(d)
	fmt.Println(e)
	fmt.Println(f)
	fmt.Println(g)
	fmt.Println(h)
	fmt.Println(i)

	a = append(a[:1], a[1+1:]...) // 删除中间1个元素


	// sort

	//SortFloat64FastV1
	var aa = []float64{4, 2, 5, 7, 2, 1, 88, 1}
	// 强制类型转换
	var bb []int = ((*[1 << 20]int)(unsafe.Pointer(&aa[0])))[:len(aa):cap(aa)]

	// 以int方式给float64排序
	sort.Ints(bb)

	//SortFloat64FastV2
	// 通过 reflect.SliceHeader 更新切片头部信息实现转换
	var cc []int
	aHdr := (*reflect.SliceHeader)(unsafe.Pointer(&aa))
	cHdr := (*reflect.SliceHeader)(unsafe.Pointer(&cc))
	*cHdr = *aHdr

	// 以int方式给float64排序
	sort.Ints(cc)


}

func TestSlience(t *testing.T) {
	
	shuzu1()
	/*
	data := []int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}

	s1 := data[9:]
	// 创建长度容量为2的切片
	s2 := data[:2]

	fmt.Println(s2)
	// s1 会替换 s2的元素
	copy(s2, s1)

	fmt.Println(s2)

	 */
}
