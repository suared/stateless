package model

import (
	"fmt"
)

const (
	NSF = iota
	NoPrice
)

type CustomError struct {
	ID      int
	Message string
}

func (e *CustomError) Code() int {
	return e.ID
}

func (e *CustomError) Error() string {
	return e.Message
}

var mrMarket *market

func init() {
	mrMarket = new(market)
	mrMarket.quotes = make(map[string]quote)
}

type account struct {
	accountID string
	cash      int64
	stocks    map[string]stock
}

func MakeAccount(id string) *account {
	newAccount := new(account)
	newAccount.accountID = id
	newAccount.stocks = make(map[string]stock)
	return newAccount
}

type stock struct {
	Symbol   string
	Quantity int64
}

type market struct {
	quotes map[string]quote
}

type quote struct {
	Symbol string
	Price  int64
}

func (a *account) AddCash(value int64) {
	a.cash += value
}

func (a *account) BuyStock(symbol string, quantity int64) error {
	currentPrice := mrMarket.quotes[symbol].Price
	if currentPrice <= 0 {
		err := &CustomError{}
		err.ID = NoPrice
		err.Message = symbol + ": No Market price available"
		return err
	}

	purchaseValue := currentPrice * quantity
	if a.cash <= purchaseValue {
		err := &CustomError{}
		err.ID = NSF
		err.Message = fmt.Sprintf("Insufficient cash: %v, to purchase: quantity: %v, symbol: %v, at current price: %v",
			a.cash, quantity, symbol, currentPrice)
		return err
	}
	a.cash -= purchaseValue

	currentStock := a.stocks[symbol]
	if currentStock.Symbol == symbol {
		currentStock.Quantity += quantity
	} else {
		currentStock.Symbol = symbol
		currentStock.Quantity = quantity
	}

	a.stocks[symbol] = currentStock

	return nil
}

func (a *account) GetTotalValue() int64 {
	var totalValue int64
	for _, stock := range a.stocks {
		marketPrice := mrMarket.quotes[stock.Symbol].Price
		if marketPrice <= 0 {
			panic(stock.Symbol + "Market price can't be zero")
		}
		totalValue += int64(stock.Quantity) * marketPrice
	}
	totalValue += a.cash
	return totalValue
}

func (a *account) SellStock(symbol string, number int64) {
	if stock, ok := a.stocks[symbol]; ok {
		if stock.Quantity < number {
			panic("Not enough shares")
		}
		stock.Quantity -= number
	} else {
		panic("Stock not found")
	}
}

func UpdateQuote(symbol string, price int64) {
	mrMarket.quotes[symbol] = quote{Symbol: symbol, Price: price}
}
