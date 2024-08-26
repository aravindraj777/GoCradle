package main

import "fmt"

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
