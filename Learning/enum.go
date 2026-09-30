package main

import "fmt"

// enum

type orderStatus int

const (
    Received orderStatus = iota
    Confirmed
    Prepared
    Delivered
)

func changeOrderStatus(status orderStatus) {
    fmt.Println("changing statis", status)
}

func printintDetail() {
    changeOrderStatus(Confirmed)
}