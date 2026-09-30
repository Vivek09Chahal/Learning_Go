package main

import "fmt"

type paymenter interface {
    pay(amount float32)
}

type payment struct{
    // gateway stripe
    // gateway razorPay
    gateway paymenter
}

func (p payment) makePayment(amount float32) {
    // razorPayPaymentGW := razorPay{} 
    // stripePaymentGW := stripe{}

    // razorPayPaymentGW.pay(amount)
    // stripePaymentGW.pay(amount)
    p.gateway.pay(amount)
}

// type stripe struct{}
// func (s stripe) pay(amount float32) {
//     fmt.Println("payment using stripe", amount)
// }

type razorPay struct{}
func (r razorPay) pay(amount float32){
    fmt.Println("payment using razorpay", amount)
}

func main() {
    // stripePaymentGW := stripe{}
    razorPaymentGW := razorPay{}
    newPayment := payment{
        // gateway: stripePaymentGW,
        gateway: razorPaymentGW,
    }
    newPayment.makePayment(100)
}
