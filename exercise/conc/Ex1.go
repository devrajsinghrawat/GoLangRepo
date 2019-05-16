package conc

import (
	"fmt"
	"runtime"
	"sync"
)

var wg sync.WaitGroup

type httpPkg struct{}

func (httpPkg) Get(url string) {
	fmt.Printf("Get done .. %s \n", url)
}

var http httpPkg

// Ex1 - Concrancy
func Ex1() {

	fmt.Println("OS        - \t", runtime.GOOS)
	fmt.Println("Arch      - \t", runtime.GOARCH)
	fmt.Println("CPU's     - \t", runtime.NumCPU())
	fmt.Println("GORoutine - \t", runtime.NumGoroutine())

	wg.Add(1)

	go foo()
	fetch()
	bar()
	fmt.Println("CPU's     - \t", runtime.NumCPU())
	fmt.Println("GORoutine - \t", runtime.NumGoroutine())
	wg.Wait()

}

func foo() {
	for index := 0; index < 50; index++ {
		fmt.Printf("Foo - %v \n", index)

	}
	wg.Done()
}

func bar() {
	for index := 0; index < 50; index++ {
		fmt.Printf("Bar - %v \n", index)
	}
}

func fetch() {
	var urls = []string{
		"http://www.golang.org/",
		"http://www.google.com/",
		"http://www.somestupidname.com/",
	}
	for _, url := range urls {
		// Increment the WaitGroup counter.
		wg.Add(1)
		// Launch a goroutine to fetch the URL.
		go func(url string) {
			// Decrement the counter when the goroutine completes.
			defer wg.Done()
			// Fetch the URL.
			http.Get(url)
		}(url)
	}
	// Wait for all HTTP fetches to complete.
	wg.Wait()
}
