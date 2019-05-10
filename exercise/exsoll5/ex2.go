package exsoll5

import "fmt"

type vehicle struct {
	doors uint
	color string
}

type truck struct {
	vehicle
	fourWheeler bool
}

type sedan struct {
	vehicle
	luxury bool
}

func (t truck) getDetails() {
	fmt.Printf("this is method of struct, color is - %v \n", t.color)
}

// Ex2 -Exported function to main - Struct, map and slice
func Ex2() {

	t := truck{vehicle: vehicle{doors: 2, color: "GRAY"},
		fourWheeler: true}
	s := sedan{vehicle: vehicle{doors: 4, color: "RED"},
		luxury: true}

	t.getDetails()

	fmt.Print(t, s)

}
