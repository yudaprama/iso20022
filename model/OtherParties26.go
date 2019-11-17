package model

// Other parties information.
type OtherParties26 struct {

	// Party, either an individual or organisation, whose assets are being invested.
	Investor *PartyIdentification99 `xml:"Invstr,omitempty"`

	// Identification of the stock exchange to which transaction reporting will be done.
	StockExchange *PartyIdentification100 `xml:"StockXchg,omitempty"`

	// Institution to which a trade must be reported.
	TradeRegulator *PartyIdentification100 `xml:"TradRgltr,omitempty"`
}

func (o *OtherParties26) AddInvestor() *PartyIdentification99 {
	o.Investor = new(PartyIdentification99)
	return o.Investor
}

func (o *OtherParties26) AddStockExchange() *PartyIdentification100 {
	o.StockExchange = new(PartyIdentification100)
	return o.StockExchange
}

func (o *OtherParties26) AddTradeRegulator() *PartyIdentification100 {
	o.TradeRegulator = new(PartyIdentification100)
	return o.TradeRegulator
}
