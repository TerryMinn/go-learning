package main

func checkAge(age, secAge int) bool {
	return age > secAge
}

func main() {
	//Basic Looping Example
	// Looping statement to print numbers from 0 to 9
	for i := 0; i < 10; i++ {

		// Break and Continue Example
		if i == 5 {
			continue
		} else if i == 8 {
			break
		}

		println(i)
	}

	nums := []int{1, 2, 3}
	maps := map[string]int{"one": 1, "two": 2, "three": 3}
	// Range Loop Example for slice
	for index, value := range nums {
		println("Index:", index, "Value:", value)
	}
	//	Range Loop Example for map
	for key, value := range maps {
		println("Key:", key, "Value:", value)
	}

	//	While Loop Example
	count := 0
	for count < 5 {
		println("Count:", count)
		count++
	}
}
