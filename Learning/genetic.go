package main

import "fmt"

func printSlice[T any](items []T) {
    for _, item := range items {
        fmt.Println(item)
    }
}

func main() {
    printSlice([]int{1, 2, 3, 4, 5})
}