package main

import (
	"bufio"
	"fmt"
	"os"

	"github.com/fatih/color"
)

func main(){
	fmt.Println("Hello, World!")
	color.Red("I am red")
	color.Blue("I am blue")
	fmt.Print("Press enter to exit.")

	reader := bufio.NewReader(os.Stdin);
	_,_ = reader.ReadString('\n');
}