package model

// Other parties information.
type OtherParties9 struct {

	// Party, either an individual or organisation, whose assets are being invested.
	Investor []*PartyIdentificationAndAccount36 `xml:"Invstr,omitempty"`

	// Foreign Financial Institution which has been authorised by local authorities to act as account management institution in the country.
	QualifiedForeignIntermediary *PartyIdentificationAndAccount37 `xml:"QlfdFrgnIntrmy,omitempty"`

	// Identification of the stock exchange to which transaction reporting will be done.
	StockExchange *PartyIdentificationAndAccount37 `xml:"StockXchg,omitempty"`

	// Institution to which a trade must be reported.
	TradeRegulator *PartyIdentificationAndAccount37 `xml:"TradRgltr,omitempty"`

	// Party responsible for the administration of a tri-party collateral transaction including collateral allocation, marking to market and substitution of collateral.
	TripartyAgent *PartyIdentificationAndAccount37 `xml:"TrptyAgt,omitempty"`
}

func (o *OtherParties9) AddInvestor() *PartyIdentificationAndAccount36 {
	newValue := new(PartyIdentificationAndAccount36)
	o.Investor = append(o.Investor, newValue)
	return newValue
}

func (o *OtherParties9) AddQualifiedForeignIntermediary() *PartyIdentificationAndAccount37 {
	o.QualifiedForeignIntermediary = new(PartyIdentificationAndAccount37)
	return o.QualifiedForeignIntermediary
}

func (o *OtherParties9) AddStockExchange() *PartyIdentificationAndAccount37 {
	o.StockExchange = new(PartyIdentificationAndAccount37)
	return o.StockExchange
}

func (o *OtherParties9) AddTradeRegulator() *PartyIdentificationAndAccount37 {
	o.TradeRegulator = new(PartyIdentificationAndAccount37)
	return o.TradeRegulator
}

func (o *OtherParties9) AddTripartyAgent() *PartyIdentificationAndAccount37 {
	o.TripartyAgent = new(PartyIdentificationAndAccount37)
	return o.TripartyAgent
}
