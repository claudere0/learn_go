package var_and_funcs

import "fmt"

//Constants cannot be declared using the := syntax

const Pi = 3.14

func Constants() {
	const World = "世界"
	fmt.Println("Hello", World)
	fmt.Println("Happy", Pi, "Day")

	const Truth = true
	fmt.Println("Go rules?", Truth)
}
