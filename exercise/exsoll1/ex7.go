package exsoll1

import "fmt"

// Ex7 -Exported function to main
func Ex7() {
	a := make([]int, 10, 12)

	fmt.Println(a)
	fmt.Println(len(a))
	fmt.Println(cap(a))
	for index := 0; index < 10; index++ {
		a[index] = index
	}
	fmt.Println(a)

	a = append(a, 10, 11)
	fmt.Println(a)
	fmt.Println(len(a))
	fmt.Println(cap(a))

	a = append(a, 12)
	fmt.Println(a)
	fmt.Println(len(a))
	fmt.Println(cap(a))
}
