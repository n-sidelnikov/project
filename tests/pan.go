package main

import "fmt"

var global string = "global"

func main() {

	var local string = "local"

	fmt.Println(global)
	fmt.Println(local)
}
