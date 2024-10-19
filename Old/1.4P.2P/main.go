package main

import "fmt"

type str string

func (t str) log() {
	fmt.Println(t)
}

func main() {
	var name str = "MAX"
	name.log()

}
