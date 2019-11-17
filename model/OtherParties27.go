package model

// Other parties information.
type OtherParties27 struct {

	// Party, either an individual or organisation, whose assets are being invested.
	Investor []*PartyIdentificationAndAccount108 `xml:"Invstr,omitempty"`

	// Foreign Financial Institution which has been authorised by local authorities to act as account management institution in the country.
	QualifiedForeignIntermediary *PartyIdentificationAndAccount107 `xml:"QlfdFrgnIntrmy,omitempty"`

	// Identification of the stock exchange to which transaction reporting will be done.
	StockExchange *PartyIdentificationAndAccount109 `xml:"StockXchg,omitempty"`

	// Institution to which a trade must be reported.
	TradeRegulator *PartyIdentificationAndAccount109 `xml:"TradRgltr,omitempty"`

	// Party responsible for the administration of a tri-party collateral transaction including collateral allocation, marking to market and substitution of collateral.
	TripartyAgent *PartyIdentificationAndAccount107 `xml:"TrptyAgt,omitempty"`

	// Party that identifies a broker when required (for example, authorised broker, prime broker, etc).
	Broker *PartyIdentificationAndAccount107 `xml:"Brkr,omitempty"`
}

func (o *OtherParties27) AddInvestor() *PartyIdentificationAndAccount108 {
	newValue := new(PartyIdentificationAndAccount108)
	o.Investor = append(o.Investor, newValue)
	return newValue
}

func (o *OtherParties27) AddQualifiedForeignIntermediary() *PartyIdentificationAndAccount107 {
	o.QualifiedForeignIntermediary = new(PartyIdentificationAndAccount107)
	return o.QualifiedForeignIntermediary
}

func (o *OtherParties27) AddStockExchange() *PartyIdentificationAndAccount109 {
	o.StockExchange = new(PartyIdentificationAndAccount109)
	return o.StockExchange
}

func (o *OtherParties27) AddTradeRegulator() *PartyIdentificationAndAccount109 {
	o.TradeRegulator = new(PartyIdentificationAndAccount109)
	return o.TradeRegulator
}

func (o *OtherParties27) AddTripartyAgent() *PartyIdentificationAndAccount107 {
	o.TripartyAgent = new(PartyIdentificationAndAccount107)
	return o.TripartyAgent
}

func (o *OtherParties27) AddBroker() *PartyIdentificationAndAccount107 {
	o.Broker = new(PartyIdentificationAndAccount107)
	return o.Broker
}
