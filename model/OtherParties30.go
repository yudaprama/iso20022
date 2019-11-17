package model

// Other parties information.
type OtherParties30 struct {

	// Party, either an individual or organisation, whose assets are being invested.
	Investor *PartyIdentification110 `xml:"Invstr,omitempty"`

	// Identification of the stock exchange to which transaction reporting will be done.
	StockExchange *PartyIdentification111 `xml:"StockXchg,omitempty"`

	// Institution to which a trade must be reported.
	TradeRegulator *PartyIdentification111 `xml:"TradRgltr,omitempty"`
}

func (o *OtherParties30) AddInvestor() *PartyIdentification110 {
	o.Investor = new(PartyIdentification110)
	return o.Investor
}

func (o *OtherParties30) AddStockExchange() *PartyIdentification111 {
	o.StockExchange = new(PartyIdentification111)
	return o.StockExchange
}

func (o *OtherParties30) AddTradeRegulator() *PartyIdentification111 {
	o.TradeRegulator = new(PartyIdentification111)
	return o.TradeRegulator
}
