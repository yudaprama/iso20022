package model

// Other parties information.
type OtherParties10 struct {

	// Party, either an individual or organisation, whose assets are being invested.
	Investor *PartyIdentification37Choice `xml:"Invstr,omitempty"`

	// Identification of the stock exchange to which transaction reporting will be done.
	StockExchange *PartyIdentification49Choice `xml:"StockXchg,omitempty"`

	// Institution to which a trade must be reported.
	TradeRegulator *PartyIdentification49Choice `xml:"TradRgltr,omitempty"`
}

func (o *OtherParties10) AddInvestor() *PartyIdentification37Choice {
	o.Investor = new(PartyIdentification37Choice)
	return o.Investor
}

func (o *OtherParties10) AddStockExchange() *PartyIdentification49Choice {
	o.StockExchange = new(PartyIdentification49Choice)
	return o.StockExchange
}

func (o *OtherParties10) AddTradeRegulator() *PartyIdentification49Choice {
	o.TradeRegulator = new(PartyIdentification49Choice)
	return o.TradeRegulator
}
