package model

// Provides details about business parties involved in the transaction.
type OtherParties18 struct {

	// Party that identifies the underlying investor.
	Investor []*PartyIdentificationAndAccount79 `xml:"Invstr,omitempty"`

	// Party that identifies the stock exchange.
	StockExchange *PartyIdentificationAndAccount87 `xml:"StockXchg,omitempty"`

	// Party that identifies the trade regulator.
	TradeRegulator *PartyIdentificationAndAccount87 `xml:"TradRgltr,omitempty"`

	// Party that handles tri-party transactions.
	TripartyAgent *PartyIdentificationAndAccount83 `xml:"TrptyAgt,omitempty"`

	// Foreign Financial Institution which has been authorised by local authorities to act as account management institution in the country.
	QualifiedForeignIntermediary *PartyIdentificationAndAccount77 `xml:"QlfdFrgnIntrmy,omitempty"`
}

func (o *OtherParties18) AddInvestor() *PartyIdentificationAndAccount79 {
	newValue := new(PartyIdentificationAndAccount79)
	o.Investor = append(o.Investor, newValue)
	return newValue
}

func (o *OtherParties18) AddStockExchange() *PartyIdentificationAndAccount87 {
	o.StockExchange = new(PartyIdentificationAndAccount87)
	return o.StockExchange
}

func (o *OtherParties18) AddTradeRegulator() *PartyIdentificationAndAccount87 {
	o.TradeRegulator = new(PartyIdentificationAndAccount87)
	return o.TradeRegulator
}

func (o *OtherParties18) AddTripartyAgent() *PartyIdentificationAndAccount83 {
	o.TripartyAgent = new(PartyIdentificationAndAccount83)
	return o.TripartyAgent
}

func (o *OtherParties18) AddQualifiedForeignIntermediary() *PartyIdentificationAndAccount77 {
	o.QualifiedForeignIntermediary = new(PartyIdentificationAndAccount77)
	return o.QualifiedForeignIntermediary
}
