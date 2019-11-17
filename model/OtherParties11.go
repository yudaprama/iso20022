package model

// Other parties information.
type OtherParties11 struct {

	// Party, either an individual or organisation, whose assets are being invested.
	Investor *PartyIdentification37Choice `xml:"Invstr,omitempty"`

	// Foreign Financial Institution which has been authorised by local authorities to act as account management institution in the country.
	QualifiedForeignIntermediary *PartyIdentification43Choice `xml:"QlfdFrgnIntrmy,omitempty"`

	// Identification of the stock exchange to which transaction reporting will be done.
	StockExchange *PartyIdentification43Choice `xml:"StockXchg,omitempty"`

	// Institution to which a trade must be reported.
	TradeRegulator *PartyIdentification43Choice `xml:"TradRgltr,omitempty"`

	// Party responsible for the administration of a tri-party collateral transaction including collateral allocation, marking to market and substitution of collateral.
	TripartyAgent *PartyIdentification43Choice `xml:"TrptyAgt,omitempty"`
}

func (o *OtherParties11) AddInvestor() *PartyIdentification37Choice {
	o.Investor = new(PartyIdentification37Choice)
	return o.Investor
}

func (o *OtherParties11) AddQualifiedForeignIntermediary() *PartyIdentification43Choice {
	o.QualifiedForeignIntermediary = new(PartyIdentification43Choice)
	return o.QualifiedForeignIntermediary
}

func (o *OtherParties11) AddStockExchange() *PartyIdentification43Choice {
	o.StockExchange = new(PartyIdentification43Choice)
	return o.StockExchange
}

func (o *OtherParties11) AddTradeRegulator() *PartyIdentification43Choice {
	o.TradeRegulator = new(PartyIdentification43Choice)
	return o.TradeRegulator
}

func (o *OtherParties11) AddTripartyAgent() *PartyIdentification43Choice {
	o.TripartyAgent = new(PartyIdentification43Choice)
	return o.TripartyAgent
}
