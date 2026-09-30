package main

import (
	"fmt"
	"time"
)

type order struct {
	id        string
	amount    float32
	status    string
	createdAt time.Time
}

func (o *order) changeStatus(status string) {
    o.status = status
}

func main() {
    myOrder := order{}

    myOrder.id = "1"
    myOrder.amount = 2
    myOrder.status = "added to cart"
    myOrder.createdAt = time.Now()

    myOrder2 := order {
        id: "2",
        amount: 100,
        status: "order placed",
        createdAt: time.Now(),
    }
    myOrder3 := order{}
    myOrder3.changeStatus("ready to deliver")

    fmt.Println(myOrder)
    fmt.Println(myOrder2)
    fmt.Println(myOrder3.status)
}