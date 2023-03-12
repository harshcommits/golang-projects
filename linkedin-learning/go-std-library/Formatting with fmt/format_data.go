package main

import (
	"fmt"
)

func main() {

	f := 123.4567

	//control decimal places
	fmt.Printf("%.2f\n", f)

	//control no. of values printed; here it prints 123.456700, which occupies 10 spaces
	fmt.Printf("%10f\n", f)

	//print with padding and precision
	fmt.Printf("%10.2f\n", f)

	//always use a + sign (can be used with -)
	fmt.Printf("%+10.2f\n", f)

	//pad with zeros instead of spaces
	fmt.Printf("%010.2f\n", f)

}
