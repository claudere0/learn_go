package struct_slice

import "fmt"

func Pointer() {
	var i int = 9
	fmt.Println("Initial value of i:", i)

	var p *int = &i
	fmt.Println("Value of pointer:", p)

	*p = 12
	fmt.Println("change value of i with p")
	fmt.Println("Now value of i:", i)

	fmt.Println("Adress of p: ", &p)

}
