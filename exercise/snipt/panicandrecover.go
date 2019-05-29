package main

// import "fmt"

// func main() {
// 	f()
// 	fmt.Println("Returned normally from f.")
// }

// func f() {
// 	defer func() { // 4
// 		if r := recover(); r != nil {
// 			fmt.Println("Recovered in f", r)
// 		}
// 	}()
// 	fmt.Println("Calling g.") // 1
// 	g(0)
// 	fmt.Println("Returned normally from g.")
// }

// func g(i int) {
// 	if i > 3 {
// 		fmt.Println("Panicking!") // 4
// 		panic(fmt.Sprintf("%v", i))
// 	}
// 	defer fmt.Println("Defer in g", i) // 5
// 	fmt.Println("Printing in g", i)    // 2
// 	g(i + 1)                           // 3
// }
