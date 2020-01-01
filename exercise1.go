//package beryl_leanGo
package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	//var user_name  string
	userName := bufio.NewReader(os.Stdin)
	fmt.Print("Enter User Name: ")
	text, _ := userName.ReadString('\n')
	fmt.Println("Hello ", text)

}
