package main

import (
	"fmt"
	"strconv"
)

func main() {
	var age = 20
	var height = float64(age)
	var boolVar = "true"
	var name = "Terry"
	var char = rune(name[0]);
	var chars = []rune(name);

	var strnum = "20";
	var floatnum = "20.5";
	// num ,err := int(strnum);
	 num,err := strconv.Atoi(strnum);
	 boolVar2,err2 := strconv.ParseBool(boolVar)
	 strBoolVar2 := strconv.FormatBool(boolVar2)
	 floatNum ,err3 := strconv.ParseFloat(floatnum,64)
	 floatConv  := strconv.FormatFloat(floatNum,'f',10,64)

	fmt.Printf("error (Atoi) is %v\n", err)
	fmt.Printf("error (ParseBool) is %v\n", err2)
	fmt.Printf("error (ParseFloat) is %v\n", err3)
	fmt.Printf("age is %d\n", age)
	fmt.Printf("height is %.2f\n", height)
	fmt.Printf("name is %s\n", name)
	fmt.Printf("char is %c digit is %d\n", char, char)
	fmt.Printf("first char in chars is %c digit is %d\n", chars[0], chars[0])
	fmt.Printf("strnum is %s \n", strnum)
	fmt.Printf("num is %d \n", num)
	fmt.Printf("boolVar2 is %t \n", boolVar2)
	fmt.Printf("strBoolVar2 is %s \n", strBoolVar2)
	fmt.Printf("floatNum is %f \n", floatConv)
}