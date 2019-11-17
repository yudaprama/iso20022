package model

// Trade market identification using a externally defined code or proprietary identification.
type TradeMarket1Choice struct {

	// Standard trade market code.
	Code *ExternalTradeMarket1Code `xml:"Cd"`

	// Trade market expressed as proprietary identification.
	Proprietary *GenericIdentification20 `xml:"Prtry"`
}

func (t *TradeMarket1Choice) SetCode(value string) {
	t.Code = (*ExternalTradeMarket1Code)(&value)
}

func (t *TradeMarket1Choice) AddProprietary() *GenericIdentification20 {
	t.Proprietary = new(GenericIdentification20)
	return t.Proprietary
}
