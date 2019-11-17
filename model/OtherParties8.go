package model

// Other parties information.
type OtherParties8 struct {

	// Party, either an individual or organisation, whose assets are being invested.
	Investor []*PartyIdentificationAndAccount40 `xml:"Invstr,omitempty"`

	// Foreign Financial Institution which has been authorised by local authorities to act as account management institution in the country.
	QualifiedForeignIntermediary *PartyIdentificationAndAccount41 `xml:"QlfdFrgnIntrmy,omitempty"`

	// Identification of the stock exchange to which transaction reporting will be done.
	StockExchange *PartyIdentificationAndAccount41 `xml:"StockXchg,omitempty"`

	// Institution to which a trade must be reported.
	TradeRegulator *PartyIdentificationAndAccount41 `xml:"TradRgltr,omitempty"`

	// Party responsible for the administration of a tri-party collateral transaction including collateral allocation, marking to market and substitution of collateral.
	TripartyAgent *PartyIdentificationAndAccount41 `xml:"TrptyAgt,omitempty"`

	// Party that identifies a broker when required (for example, authorised broker, prime broker, etc).
	Broker *PartyIdentificationAndAccount41 `xml:"Brkr,omitempty"`
}

func (o *OtherParties8) AddInvestor() *PartyIdentificationAndAccount40 {
	newValue := new(PartyIdentificationAndAccount40)
	o.Investor = append(o.Investor, newValue)
	return newValue
}

func (o *OtherParties8) AddQualifiedForeignIntermediary() *PartyIdentificationAndAccount41 {
	o.QualifiedForeignIntermediary = new(PartyIdentificationAndAccount41)
	return o.QualifiedForeignIntermediary
}

func (o *OtherParties8) AddStockExchange() *PartyIdentificationAndAccount41 {
	o.StockExchange = new(PartyIdentificationAndAccount41)
	return o.StockExchange
}

func (o *OtherParties8) AddTradeRegulator() *PartyIdentificationAndAccount41 {
	o.TradeRegulator = new(PartyIdentificationAndAccount41)
	return o.TradeRegulator
}

func (o *OtherParties8) AddTripartyAgent() *PartyIdentificationAndAccount41 {
	o.TripartyAgent = new(PartyIdentificationAndAccount41)
	return o.TripartyAgent
}

func (o *OtherParties8) AddBroker() *PartyIdentificationAndAccount41 {
	o.Broker = new(PartyIdentificationAndAccount41)
	return o.Broker
}
