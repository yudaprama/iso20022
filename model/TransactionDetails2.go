package model

// Reference of the trade transaction.
type TransactionDetails2 struct {

	// Unique identification assigned to a trade once it is received or matched by an executing system.
	TradeReference *Max70Text `xml:"TradRef"`
}

func (t *TransactionDetails2) SetTradeReference(value string) {
	t.TradeReference = (*Max70Text)(&value)
}
