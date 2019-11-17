package model

// Other parties information.
type OtherParties4 struct {

	// Party, either an individual or organisation, whose assets are being invested.
	Investor *PartyIdentification14Choice `xml:"Invstr,omitempty"`

	// Identification of the stock exchange to which transaction reporting will be done.
	StockExchange *PartyIdentification10Choice `xml:"StockXchg,omitempty"`

	// Institution to which a trade must be reported.
	TradeRegulator *PartyIdentification10Choice `xml:"TradRgltr,omitempty"`
}

func (o *OtherParties4) AddInvestor() *PartyIdentification14Choice {
	o.Investor = new(PartyIdentification14Choice)
	return o.Investor
}

func (o *OtherParties4) AddStockExchange() *PartyIdentification10Choice {
	o.StockExchange = new(PartyIdentification10Choice)
	return o.StockExchange
}

func (o *OtherParties4) AddTradeRegulator() *PartyIdentification10Choice {
	o.TradeRegulator = new(PartyIdentification10Choice)
	return o.TradeRegulator
}
