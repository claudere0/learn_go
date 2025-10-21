package loop_and_statements

import (
	"fmt"
)

func Sqrt(x float64) float64 {
	z := 1.0
	fmt.Printf("Initial z for x=%v: %v\n", x, z)
	for i := 0; i < 10; i++ {
		z -= (z*z - x) / (2 * z)
		fmt.Printf("Iteration %d: z = %v\n", i+1, z)
	}

	return z
}

func LoopAndFuncs() {
	fmt.Println(Sqrt(2))
}
