package model

// Other parties information.
type OtherParties31 struct {

	// Party, either an individual or organisation, whose assets are being invested.
	Investor *PartyIdentification110 `xml:"Invstr,omitempty"`

	// Foreign Financial Institution which has been authorised by local authorities to act as account management institution in the country.
	QualifiedForeignIntermediary *PartyIdentification111 `xml:"QlfdFrgnIntrmy,omitempty"`

	// Identification of the stock exchange to which transaction reporting will be done.
	StockExchange *PartyIdentification111 `xml:"StockXchg,omitempty"`

	// Institution to which a trade must be reported.
	TradeRegulator *PartyIdentification111 `xml:"TradRgltr,omitempty"`

	// Party responsible for the administration of a tri-party collateral transaction including collateral allocation, marking to market and substitution of collateral.
	TripartyAgent *PartyIdentification111 `xml:"TrptyAgt,omitempty"`
}

func (o *OtherParties31) AddInvestor() *PartyIdentification110 {
	o.Investor = new(PartyIdentification110)
	return o.Investor
}

func (o *OtherParties31) AddQualifiedForeignIntermediary() *PartyIdentification111 {
	o.QualifiedForeignIntermediary = new(PartyIdentification111)
	return o.QualifiedForeignIntermediary
}

func (o *OtherParties31) AddStockExchange() *PartyIdentification111 {
	o.StockExchange = new(PartyIdentification111)
	return o.StockExchange
}

func (o *OtherParties31) AddTradeRegulator() *PartyIdentification111 {
	o.TradeRegulator = new(PartyIdentification111)
	return o.TradeRegulator
}

func (o *OtherParties31) AddTripartyAgent() *PartyIdentification111 {
	o.TripartyAgent = new(PartyIdentification111)
	return o.TripartyAgent
}
