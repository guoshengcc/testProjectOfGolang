package main

import (
	"fmt"
	"testproject/cmd"
)

func main() {
	// fmt.Println("Hi")
	err := cmd.Run("C:\\GitHub\\gomock", "./HeadContent.xml")
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println("Ok")
	}
}
