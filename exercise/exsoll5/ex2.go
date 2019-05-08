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

// Ex2 -Exported function to main - Struct, map and slice
func Ex2() {

	t := truck{vehicle: vehicle{doors: 2, color: "GRAY"},
		fourWheeler: true}
	s := sedan{vehicle: vehicle{doors: 4, color: "RED"},
		luxury: true}

	fmt.Print(t, s)

}
