package main

import (
	"fmt"
	"slices"
	"strings"
)

func main() {
	//Binart Search
	nums := []int{2, 4, 6, 8}
	indexer, err := slices.BinarySearch(nums, 2)

	if err != true {
		fmt.Println("Error:", err)
		return
	}

	fmt.Println(indexer) //0

	// Clip
	sliceCliper := make([]int, 5, 10)
	fmt.Println(cap(sliceCliper))
	sliceCliper = slices.Clip(sliceCliper)
	fmt.Println(cap(sliceCliper))

	// clone
	newSlice := []int{1, 2, 3, 4, 5}
	copySlice := slices.Clone(newSlice)
	fmt.Println(copySlice)

	// Compact
	nameSlice := []string{"Terry", "Terry", "Leon", "Terry", "LEON", "TERRY", "marry"} // Duplicates
	nameSlice = slices.Compact(nameSlice)                                              // Remove Duplicates
	fmt.Println(nameSlice)                                                             //["Terry", "Leon", "Terry"]

	funcNameSlice := []string{"Terry", "Terry", "Leon", "Terry", "LEON", "TERRY", "marry"}
	funcNameSlice = slices.CompactFunc(funcNameSlice, func(a, b string) bool {
		return strings.EqualFold(a, b)
	})
	fmt.Println(funcNameSlice) //["Terry", "Leon", "marry"]

	// Compare & Equal
	sliceA := []int{1, 2, 3}
	sliceB := []int{1, 2, 3}
	sliceC := []int{1, 2, 4}
	//areEqual := slices.Compare(sliceA, sliceB)  //0
	//areEqual2 := slices.Compare(sliceA, sliceC) //-1
	areEqual := slices.Equal(sliceA, sliceB)  //true
	areEqual2 := slices.Equal(sliceA, sliceC) //false
	//fmt.Println(areEqual)                       // 0
	//fmt.Println(areEqual2)                      // -1
	fmt.Println(areEqual)  // true
	fmt.Println(areEqual2) // false

	// Contains
	numbers := []int{1, 2, 3, 4, 5}
	contains := slices.Contains(numbers, 3)
	contains2 := slices.Contains(numbers, 6)
	fmt.Println(contains2)
	fmt.Println(contains)

	// Delete
	numSlice := []int{1, 2, 3, 4, 5}
	numSlice = slices.Delete(numSlice, 0, 4)
	fmt.Println(numSlice) // [5]

	//	index
	numSlice2 := []int{10, 20, 30, 40, 50}
	index := slices.Index(numSlice2, 30)
	fmt.Println(index) //2

	//	insert
	numSlice3 := []int{1, 2, 3, 4, 5}
	numSlice3 = slices.Insert(numSlice3, 2, 99, 100, 200)
	fmt.Println(numSlice3) // [1 2 99 100 3 4 5]

	//	isSorted
	numSlice4 := []int{1, 2, 3, 4, 5}
	stringSlice := []string{"apple", "banana", "cherry"} // base on a b c
	isSorted := slices.IsSorted(numSlice4)
	isSorted2 := slices.IsSorted(stringSlice)
	fmt.Println(isSorted)  // true
	fmt.Println(isSorted2) // true

	//	Max and Min
	numSlice5 := []int{10, 20, 5, 40, 15}
	maxValue := slices.Max(numSlice5)
	minValue := slices.Min(numSlice5)
	fmt.Println("Max:", maxValue) // Max: 40
	fmt.Println("Min:", minValue) // Min: 5

	// Replace & Sort
	numSlice6 := []int{5, 3, 8, 1, 2}
	slices.Replace(numSlice6, 1, 4, 99, 100)
	fmt.Println("After Replace:", numSlice6) // After Replace: [5 99 100 1 2]

	slices.Sort(numSlice6)
	fmt.Println("After Sort:", numSlice6) // After Sort: [1 2 5 99 100]ထ
}
