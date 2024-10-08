package main

import (
	"fmt"
	"oop-in-go/payment"
)

func main() {

	credit := payment.CreateCreditAccount(
		"Debra Ingram",
		"1111-2222-3333-4444",
		4,
		2021,
		123)

	fmt.Printf("Owner name: %v\n", credit.OwnerName())
	fmt.Printf("Card number: %v\n", credit.CardNumber())
	fmt.Println("Trying to change card number")
	err := credit.SetCardNumber("invalid")
	if err != nil {
		fmt.Printf("That didn't work: %v\n", err)
	}
	fmt.Printf("Available credit: %v\n", credit.AvailableCredit())

}
