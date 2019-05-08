package exsoll1

import "fmt"

var agrData []string
var argdata1 []string

// Ex6 -Exported function to main
func Ex6() {
	agrData = append(agrData, "zero", "one")

	argdata1 = append(argdata1, "0", "1", "2", "3")
	// fmt.Println(len(agrData))
	// fmt.Println(agrData[0])
	// fmt.Println(&agrData)

	agrData = append(agrData, argdata1...)
	fmt.Println(agrData)
	for index, element := range agrData {

		if index%2 == 0 {
			fmt.Println(index, element)
		} else {
			low := index
			high := index
			if len(agrData) != index {
				high = index + 1
			}
			agrData = append(agrData[:low], agrData[high:]...)

		}
	}
	fmt.Println(agrData)
}
