package main

import (
	"fmt"
	"slices"
)

func printArray() {
	fmt.Println("Print array")
	var arr [3]string
	arr = [3]string{"Aravind", "Sreekanth", "Arun"}
	fmt.Println("array : ", arr)
}
func intArray() {
	fmt.Println("Integer Array")
	arr2 := [3]int{12, 3, 5}
	fmt.Println(arr2)
}
func slicesPrinting() {
	fmt.Println("Slices example printing")
	s := []string{"a", "b", "c"}
	fmt.Println(s)
	//appending value
	s = append(s, "d")
	fmt.Println(s)
	s = slices.Delete(s, 1, 2)
	fmt.Println(s)
}
