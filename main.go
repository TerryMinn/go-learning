package main

import "fmt"

func main() {
	//var family = [3]string{"Terry", "Shin", "Minn"}
	//var home = [...]int{1, 2, 3}
	//school := [3]rune{'a', 'b', 'c'}
	//var office [2]bool
	//
	//secFamily := family
	//family[0] = "Shinn Thant Minn"
	//
	//secSchool := school[1:3]
	//thirdSchool := school[:3]
	//
	//secHome := home[:]
	//secHome[0] = 100
	//secHome = append(secHome, 200)
	//dummy := [3]string{}
	//
	//copy(dummy[:], family[:])
	//
	//matrix := [3][2]int{
	//	{1, 2},
	//	{3, 4},
	//	{5, 6},
	//}
	//secMatrix := [...][2]int{
	//	{1, 2},
	//	{3, 4},
	//	{5, 6},
	//}
	//
	//fmt.Println(family, home, school, office)
	//fmt.Println(secHome, secSchool, thirdSchool, secFamily, matrix)
	//fmt.Println(secMatrix)

	//b := []int{1, 2, 3, 4}
	//
	////b[4] = 5
	//b = append(b, 5)
	//
	//c := []int{6, 7, 8}
	//a := [2]int{9, 10}
	//
	//b = append(b, c...)
	//b = append(b, a[:]...)
	//
	//fmt.Println("b is", b)
	//fmt.Println("c is", c)
	//fmt.Println("c is", a)

	//a := []int{}
	//b := make([]int, 2, 5)
	//
	//fmt.Println("a is", a)
	//fmt.Println("b is", b)

	a := []int{1, 2, 4}
	c := []int{}

	b := a[:0]

	lener := copy(c, a)

	fmt.Println("c is", c, "lener is", lener)
	fmt.Println("b is", b)
}
