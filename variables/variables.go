package variables

import "fmt"

//know all variables declared forms
func Variables() {
	var variable1 string = "variable 1"
	variable2 := "variable 2"

	var (
		variable3 string = "12ada3"
		variable4 string
	)
	variable5, variable6 := "five", "six"

	const constant1 string = "constant"

	variable5, variable6 = variable6, variable5

	fmt.Println(variable1,variable2, variable3,variable4, variable5, variable6, constant1)
}