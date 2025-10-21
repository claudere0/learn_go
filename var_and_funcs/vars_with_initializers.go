package var_and_funcs

import "fmt"

var i, j int = 1, 2

func Var() {
	var c, python, java = true, false, "no!"
	fmt.Println(i, j, c, python, java)
}
