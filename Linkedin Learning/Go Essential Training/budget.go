package main

import (
	"fmt"
	"time"
)

type Budget struct {
	CampaignID string
	Balance    float64
	Expires    time.Time
}

func NewBudget(campaignID string, balance float64, expires time.Time) (*Budget, error) {

	if campaignID == "" {
		return nil, fmt.Errorf("empty campaign ID")
	}

	if balance <= 0 {
		return nil, fmt.Errorf("balance must be bigger than zero")
	}

	if expires.Before(time.Now()) {
		return nil, fmt.Errorf("bad expiration date")
	}

	b := Budget{
		CampaignID: campaignID,
		Balance:    balance,
		Expires:    expires,
	}

	return &b, nil

}

func main() {
	expires := time.Now().Add(7 * 24 * time.Hour)

	b1, err := NewBudget("puppies", 32.2, expires)
	if err != nil {
		fmt.Println("ERROR: ", err)
	} else {
		fmt.Printf("%#v\n", b1)
	}

	b2, err := NewBudget("kittens", -3.2, expires)
	if err != nil {
		fmt.Println("ERROR: ", err)
	} else {
		fmt.Printf("%#v\n", b2)
	}

}
