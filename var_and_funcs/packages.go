package var_and_funcs

import (
	"fmt"
	"math/rand"
)

func Packages() {
	fmt.Println("My favorite number is", rand.Intn(10))
	fmt.Println("But i also like", rand.Intn(100))
}
