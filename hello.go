package main

import "fmt"

func main() {
    var slice1 [][]int = make([][]int, 3)
	fmt.Println("Hello Gophers!")
	fmt.Println(slice1)
    // #=> [[] [] []]

    slice2 := [][]int{{}, {}, {}}
    fmt.Println(slice2)
    // #=> [[] [] []]
}