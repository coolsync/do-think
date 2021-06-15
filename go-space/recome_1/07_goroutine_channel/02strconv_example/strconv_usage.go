package main

import (
	"fmt"
	"os"
	"strconv"
)

func main() {
	// parse string to int64 type
	intStr := "10000"
	res1, err := strconv.ParseInt(intStr, 10, 64)
	if err != nil {
		fmt.Fprintf(os.Stdin, "parse string failed, err: %v\n", err)
		os.Exit(1)
	}
	fmt.Printf("%v, %T\n", res1, res1)

	res2, _ := strconv.Atoi(intStr)
	fmt.Printf("%v, %T\n", res2, res2)

	// parse string to float64
	floatStr := "1.123"
	floatVal, _ := strconv.ParseFloat(floatStr, 64)
	fmt.Printf("%v, %T\n", floatVal, floatVal)

	// parse string to bool
	boolStr := "true"
	boolVal, _ := strconv.ParseBool(boolStr)
	fmt.Printf("%v, %T\n", boolVal, boolVal)

	// parse int to string
	n := 97
	// resInt := fmt.Sprintf("%d", n)
	str1 := strconv.FormatInt(int64(n), 10)
	fmt.Printf("%v, %T\n", str1, str1)

	resInt := strconv.Itoa(n)
	fmt.Printf("%q, %T\n", resInt, resInt)

}
