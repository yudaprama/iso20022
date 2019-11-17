package model

// Other parties information.
type OtherParties3 struct {

	// Party, either an individual or organisation, whose assets are being invested.
	Investor *PartyIdentification14Choice `xml:"Invstr,omitempty"`

	// Foreign Financial Institution which has been authorised by local authorities to act as account management institution in the country.
	QualifiedForeignIntermediary *PartyIdentification10Choice `xml:"QlfdFrgnIntrmy,omitempty"`

	// Identification of the stock exchange to which transaction reporting will be done.
	StockExchange *PartyIdentification10Choice `xml:"StockXchg,omitempty"`

	// Institution to which a trade must be reported.
	TradeRegulator *PartyIdentification10Choice `xml:"TradRgltr,omitempty"`

	// Party responsible for the administration of a tri-party collateral transaction including collateral allocation, marking to market and substitution of collateral.
	TripartyAgent *PartyIdentification10Choice `xml:"TrptyAgt,omitempty"`
}

func (o *OtherParties3) AddInvestor() *PartyIdentification14Choice {
	o.Investor = new(PartyIdentification14Choice)
	return o.Investor
}

func (o *OtherParties3) AddQualifiedForeignIntermediary() *PartyIdentification10Choice {
	o.QualifiedForeignIntermediary = new(PartyIdentification10Choice)
	return o.QualifiedForeignIntermediary
}

func (o *OtherParties3) AddStockExchange() *PartyIdentification10Choice {
	o.StockExchange = new(PartyIdentification10Choice)
	return o.StockExchange
}

func (o *OtherParties3) AddTradeRegulator() *PartyIdentification10Choice {
	o.TradeRegulator = new(PartyIdentification10Choice)
	return o.TradeRegulator
}

func (o *OtherParties3) AddTripartyAgent() *PartyIdentification10Choice {
	o.TripartyAgent = new(PartyIdentification10Choice)
	return o.TripartyAgent
}
