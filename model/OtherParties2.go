package model

// Other parties information.
type OtherParties2 struct {

	// Party, either an individual or organisation, whose assets are being invested.
	Investor []*PartyIdentificationAndAccount19 `xml:"Invstr,omitempty"`

	// Foreign Financial Institution which has been authorised by local authorities to act as account management institution in the country.
	QualifiedForeignIntermediary *PartyIdentificationAndAccount21 `xml:"QlfdFrgnIntrmy,omitempty"`

	// Identification of the stock exchange to which transaction reporting will be done.
	StockExchange *PartyIdentificationAndAccount21 `xml:"StockXchg,omitempty"`

	// Institution to which a trade must be reported.
	TradeRegulator *PartyIdentificationAndAccount21 `xml:"TradRgltr,omitempty"`

	// Party responsible for the administration of a tri-party collateral transaction including collateral allocation, marking to market and substitution of collateral.
	TripartyAgent *PartyIdentificationAndAccount21 `xml:"TrptyAgt,omitempty"`
}

func (o *OtherParties2) AddInvestor() *PartyIdentificationAndAccount19 {
	newValue := new(PartyIdentificationAndAccount19)
	o.Investor = append(o.Investor, newValue)
	return newValue
}

func (o *OtherParties2) AddQualifiedForeignIntermediary() *PartyIdentificationAndAccount21 {
	o.QualifiedForeignIntermediary = new(PartyIdentificationAndAccount21)
	return o.QualifiedForeignIntermediary
}

func (o *OtherParties2) AddStockExchange() *PartyIdentificationAndAccount21 {
	o.StockExchange = new(PartyIdentificationAndAccount21)
	return o.StockExchange
}

func (o *OtherParties2) AddTradeRegulator() *PartyIdentificationAndAccount21 {
	o.TradeRegulator = new(PartyIdentificationAndAccount21)
	return o.TradeRegulator
}

func (o *OtherParties2) AddTripartyAgent() *PartyIdentificationAndAccount21 {
	o.TripartyAgent = new(PartyIdentificationAndAccount21)
	return o.TripartyAgent
}
