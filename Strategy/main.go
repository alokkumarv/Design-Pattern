package main

import "log"

type PaymentStrategy interface {
	Pay(amount float64)
}

type CredittCardStrategy struct {
	name, cardnumber string
}

func (c *CredittCardStrategy) Pay(amount float64) {
	log.Printf("Amount payed using credit card %v by %v", c.cardnumber, c.name)
}

type PaypalStrategy struct {
	email string
}

func (p *PaypalStrategy) Pay(amount float64) {
	log.Printf("Amount payed using paypal email %v", p.email)
}

type CrypotoStrategy struct {
	walletId string
}

func (c *CrypotoStrategy) Pay(amount float64) {
	log.Printf("Amount payed using crypto %v", c.walletId)
}

// Context

type PaymenProcessor struct {
	strategy PaymentStrategy
}

func (p *PaymenProcessor) SetStrategy(strategy PaymentStrategy) {
	p.strategy = strategy
}

func (p *PaymenProcessor) ExecutePayment(amount float64) {
	p.strategy.Pay(amount)
}

func main() {

	creditCard := &CredittCardStrategy{name: "Alok", cardnumber: "3456 4567 56433 6465"}
	paymentprocessor := &PaymenProcessor{}
	paymentprocessor.SetStrategy(creditCard)
	paymentprocessor.ExecutePayment(45.55)

	cryptoStrategy := &CrypotoStrategy{walletId: "3456345"}
	paymentprocessor.SetStrategy(cryptoStrategy)
	paymentprocessor.ExecutePayment(78.78)

}
