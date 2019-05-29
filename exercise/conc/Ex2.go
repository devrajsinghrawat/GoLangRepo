package conc

import (
	"fmt"
	"sync"
)

// Ex2 - Concrancy - race condition
func Ex2() {
	var wg sync.WaitGroup
	counter := 0

	gr := 100
	wg.Add(gr)
	for index := 0; index < gr; index++ {
		go func() {
			defer wg.Done()
			v := counter
			v++
			counter = v

		}()
	}
	wg.Wait()
	fmt.Printf("Counter %v \n", counter)
}
