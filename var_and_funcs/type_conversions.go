package var_and_funcs

import (
	"fmt"
	"math"
)

func TypeConv() {
	var x, y int = 3, 4
	var f float64 = math.Sqrt(float64(x*x + y*y))
	var z uint = uint(f)
	var r rune = rune(x)
	t := rune(y)
	fmt.Println(x, y, z, r, t)
}
