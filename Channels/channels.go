package main

import (
	"fmt"
	"math/rand"
	"time"
)

func processNum(numChan chan int) {
    for num := range numChan {
        fmt.Println("processing number", num)
        time.Sleep(time.Second * 1)
    }
}

func main() {

    numChan := make(chan int)

    go processNum(numChan)

    // numChan <- 5
    for {
        numChan <- rand.Intn(100)
    }
    // time.Sleep(time.Second * 1)
    
//     messageChannel := make(chan string)
// 
//     messageChannel <- "ping"
// 
//     message := <- messageChannel
// 
//     fmt.Println(message)
}