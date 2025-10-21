package loop_and_statements

import "fmt"

//A defer statement defers the execution of a function until the surrounding function returns.

func ToDefer() {
	defer fmt.Println("world")

	fmt.Println("hello")
}
