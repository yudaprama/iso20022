package model

// Other parties information.
type OtherParties29 struct {

	// Party, either an individual or organisation, whose assets are being invested.
	Investor []*PartyIdentificationAndAccount135 `xml:"Invstr,omitempty"`

	// Foreign Financial Institution which has been authorised by local authorities to act as account management institution in the country.
	QualifiedForeignIntermediary *PartyIdentificationAndAccount136 `xml:"QlfdFrgnIntrmy,omitempty"`

	// Identification of the stock exchange to which transaction reporting will be done.
	StockExchange *PartyIdentificationAndAccount137 `xml:"StockXchg,omitempty"`

	// Institution to which a trade must be reported.
	TradeRegulator *PartyIdentificationAndAccount137 `xml:"TradRgltr,omitempty"`

	// Party responsible for the administration of a tri-party collateral transaction including collateral allocation, marking to market and substitution of collateral.
	TripartyAgent *PartyIdentificationAndAccount136 `xml:"TrptyAgt,omitempty"`

	// Party that identifies a broker when required (for example, authorised broker, prime broker, etc).
	Broker *PartyIdentificationAndAccount136 `xml:"Brkr,omitempty"`
}

func (o *OtherParties29) AddInvestor() *PartyIdentificationAndAccount135 {
	newValue := new(PartyIdentificationAndAccount135)
	o.Investor = append(o.Investor, newValue)
	return newValue
}

func (o *OtherParties29) AddQualifiedForeignIntermediary() *PartyIdentificationAndAccount136 {
	o.QualifiedForeignIntermediary = new(PartyIdentificationAndAccount136)
	return o.QualifiedForeignIntermediary
}

func (o *OtherParties29) AddStockExchange() *PartyIdentificationAndAccount137 {
	o.StockExchange = new(PartyIdentificationAndAccount137)
	return o.StockExchange
}

func (o *OtherParties29) AddTradeRegulator() *PartyIdentificationAndAccount137 {
	o.TradeRegulator = new(PartyIdentificationAndAccount137)
	return o.TradeRegulator
}

func (o *OtherParties29) AddTripartyAgent() *PartyIdentificationAndAccount136 {
	o.TripartyAgent = new(PartyIdentificationAndAccount136)
	return o.TripartyAgent
}

func (o *OtherParties29) AddBroker() *PartyIdentificationAndAccount136 {
	o.Broker = new(PartyIdentificationAndAccount136)
	return o.Broker
}
