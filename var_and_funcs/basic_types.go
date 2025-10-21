package var_and_funcs

import (
	"fmt"
	"math/cmplx"
)

var (
	ToFloat float64    = 3.0034598938593
	MaxInt  uint64     = 1<<64 - 1
	z       complex128 = cmplx.Sqrt(-5 + 12i)
)

func BasicTypes() {
	fmt.Printf("Type: %T Value: %v\n", ToFloat, ToFloat)
	fmt.Printf("Type: %T Value: %v\n", MaxInt, MaxInt)
	fmt.Printf("Type: %T Value: %v\n", z, z)
}
