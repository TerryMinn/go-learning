package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main(){
	fmt.Println("Can you tell me your age?")
	reader := bufio.NewReader(os.Stdin)
	input,err := reader.ReadString('\n');

	if err != nil {
		fmt.Printf("Error is %v",err);
		return;
	}

	input = strings.TrimSpace(input)
	number,_ := strconv.Atoi(input)

	birthYear := 2025 - number;
	fmt.Printf("Your birthyear is %d",birthYear); 
}