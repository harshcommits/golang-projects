package payment

import (
	"errors"
)

type CreditCard struct {
	ownerName       string
	cardNumber      string
	expirationMonth int
	expirationYear  int
	securityCode    int
	availableCredit float32
}

func CreateCreditAccount(ownerName, cardNumber string, expirationMonth, expirationYear, securityCode int) *CreditCard {
	return &CreditCard{
		ownerName:       ownerName,
		cardNumber:      cardNumber,
		expirationMonth: expirationMonth,
		expirationYear:  expirationYear,
		securityCode:    securityCode,
	}
}

func (c *CreditCard) OwnerName() string {
	return c.OwnerName
}

func (c *CreditCard) SetOwnerName(value string) string {
	if len(value) == 0 {
		return errors.new("Invalid owner name provided")
	}
}
