package test

import (
	"log"
	"testing"

	stocks "github.com/suared/stateless/model/stocks"
)

func TestMarketModel(t *testing.T) {
	stocks.UpdateQuote("QQQ", 12323)
	davidsAccount := stocks.MakeAccount("David")
	davidsAccount.AddCash(999999999)
	davidsAccount.BuyStock("QQQ", 3)

	log.Printf("davids account is: %v", davidsAccount)
	//Tests:  Get; Patch (update, move, add, delete)
	//log.Printf("Running API Lifecyce Test E2E with uri: %v", appURI)
	//t.Errorf("Could not marshall action object: %v", err)
}
