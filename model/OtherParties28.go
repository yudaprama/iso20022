package model

// Other parties information.
type OtherParties28 struct {

	// Party, either an individual or organisation, whose assets are being invested.
	Investor *PartyIdentification99 `xml:"Invstr,omitempty"`

	// Foreign Financial Institution which has been authorised by local authorities to act as account management institution in the country.
	QualifiedForeignIntermediary *PartyIdentification100 `xml:"QlfdFrgnIntrmy,omitempty"`

	// Identification of the stock exchange to which transaction reporting will be done.
	StockExchange *PartyIdentification100 `xml:"StockXchg,omitempty"`

	// Institution to which a trade must be reported.
	TradeRegulator *PartyIdentification100 `xml:"TradRgltr,omitempty"`

	// Party responsible for the administration of a tri-party collateral transaction including collateral allocation, marking to market and substitution of collateral.
	TripartyAgent *PartyIdentification100 `xml:"TrptyAgt,omitempty"`
}

func (o *OtherParties28) AddInvestor() *PartyIdentification99 {
	o.Investor = new(PartyIdentification99)
	return o.Investor
}

func (o *OtherParties28) AddQualifiedForeignIntermediary() *PartyIdentification100 {
	o.QualifiedForeignIntermediary = new(PartyIdentification100)
	return o.QualifiedForeignIntermediary
}

func (o *OtherParties28) AddStockExchange() *PartyIdentification100 {
	o.StockExchange = new(PartyIdentification100)
	return o.StockExchange
}

func (o *OtherParties28) AddTradeRegulator() *PartyIdentification100 {
	o.TradeRegulator = new(PartyIdentification100)
	return o.TradeRegulator
}

func (o *OtherParties28) AddTripartyAgent() *PartyIdentification100 {
	o.TripartyAgent = new(PartyIdentification100)
	return o.TripartyAgent
}
