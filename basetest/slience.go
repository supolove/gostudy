package basetest

import "fmt"

func main6() {
	data := []int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}

	s1 := data[8:]
	s2 := data[:2]

	copy(s2, s1)

	fmt.Println(s2)

}
