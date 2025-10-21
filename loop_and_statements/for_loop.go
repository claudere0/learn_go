package loop_and_statements

import "fmt"

/*
the init statement: executed before the first iteration
the condition expression: evaluated before every iteration
the post statement: executed at the end of every iteration

The loop will stop iterating once the boolean condition evaluates to false.

*/

func SumFor() {
	sum := 0
	for i := 0; i < 10; i++ {
		sum += i
	}
	fmt.Println(sum)
}

//The init and post statements are optional.
//At that point you can drop the semicolons: C's while is spelled for in Go.

func ForNowIsWhile() {
	sum := 1

	for sum < 1000 {
		sum += sum
	}
	fmt.Println(sum)
}

/*

If you omit the loop condition it loops forever, so an infinite loop is compactly expressed.

for {
}

*/
