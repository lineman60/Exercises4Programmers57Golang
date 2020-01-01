package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	//var user_name  string
	userName := bufio.NewReader(os.Stdin)
	fmt.Print("Enter a string: ")
	text, _ := userName.ReadString('\n')
	strText := strings.TrimSpace(text)
	fmt.Println("The String ", strText, " is ", len(strText), " in Length")

}
