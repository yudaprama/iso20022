package model

// Other parties information.
type OtherParties12 struct {

	// Party, either an individual or organisation, whose assets are being invested.
	Investor []*PartyIdentificationAndAccount46 `xml:"Invstr,omitempty"`

	// Foreign Financial Institution which has been authorised by local authorities to act as account management institution in the country.
	QualifiedForeignIntermediary *PartyIdentificationAndAccount47 `xml:"QlfdFrgnIntrmy,omitempty"`

	// Identification of the stock exchange to which transaction reporting will be done.
	StockExchange *PartyIdentificationAndAccount47 `xml:"StockXchg,omitempty"`

	// Institution to which a trade must be reported.
	TradeRegulator *PartyIdentificationAndAccount47 `xml:"TradRgltr,omitempty"`

	// Party responsible for the administration of a tri-party collateral transaction including collateral allocation, marking to market and substitution of collateral.
	TripartyAgent *PartyIdentificationAndAccount47 `xml:"TrptyAgt,omitempty"`

	// Party that identifies a broker when required (for example, authorised broker, prime broker, etc).
	Broker *PartyIdentificationAndAccount41 `xml:"Brkr,omitempty"`
}

func (o *OtherParties12) AddInvestor() *PartyIdentificationAndAccount46 {
	newValue := new(PartyIdentificationAndAccount46)
	o.Investor = append(o.Investor, newValue)
	return newValue
}

func (o *OtherParties12) AddQualifiedForeignIntermediary() *PartyIdentificationAndAccount47 {
	o.QualifiedForeignIntermediary = new(PartyIdentificationAndAccount47)
	return o.QualifiedForeignIntermediary
}

func (o *OtherParties12) AddStockExchange() *PartyIdentificationAndAccount47 {
	o.StockExchange = new(PartyIdentificationAndAccount47)
	return o.StockExchange
}

func (o *OtherParties12) AddTradeRegulator() *PartyIdentificationAndAccount47 {
	o.TradeRegulator = new(PartyIdentificationAndAccount47)
	return o.TradeRegulator
}

func (o *OtherParties12) AddTripartyAgent() *PartyIdentificationAndAccount47 {
	o.TripartyAgent = new(PartyIdentificationAndAccount47)
	return o.TripartyAgent
}

func (o *OtherParties12) AddBroker() *PartyIdentificationAndAccount41 {
	o.Broker = new(PartyIdentificationAndAccount41)
	return o.Broker
}
