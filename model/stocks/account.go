package model

var mrMarket *market

func init() {
	mrMarket = new(market)
}

type account struct {
	accountID string
	cash      int64
	stocks    map[string]stock
}

func MakeAccount(id string) *account {
	newAccount := new(account)
	newAccount.accountID = id
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

func (a *account) BuyStock(symbol string, quantity int64) {
	currentStock := a.stocks[symbol]
	if currentStock.Symbol == symbol {
		currentStock.Quantity += quantity
	} else {
		currentStock.Symbol = symbol
		currentStock.Quantity = quantity
	}

	a.stocks[symbol] = currentStock
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
