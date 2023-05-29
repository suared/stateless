package Test

import (
	"log"
	"testing"
)

func TestMarketModel(t *testing.T) {
	davidsAccount := stateless.model.stocks.MakeAccount("David")
	davidsAccount.BuyStock("QQQ", 3)

	log.Printf("davids account is: %v", davidsAccount)
	//Tests:  Get; Patch (update, move, add, delete)
	//log.Printf("Running API Lifecyce Test E2E with uri: %v", appURI)
	//t.Errorf("Could not marshall action object: %v", err)
}
