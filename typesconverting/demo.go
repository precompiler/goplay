package main

import (
	"fmt"
	"reflect"
	"strconv"
)

func main() {
	integer := 100
	fmt.Println(reflect.TypeOf(integer))
	fmt.Printf("%T\n", integer)

	fmt.Println(float64(integer))
	fmt.Printf("%T\n", float64(integer))
	// not gonna work
	// number := 10
	// str := "String concat " + number

	PI := 3.14
	var x int64 = int64(PI)
	fmt.Println(x)

	number := 49
	str := "String concat " + string(number) // 49 is ascii code for char '1'
	fmt.Println(str)

	str = "String concat " + strconv.Itoa(number) // int to ascii
	fmt.Println(str)

}
