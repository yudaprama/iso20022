package model

// Choice of format for the trading capacity.
type TradingPartyCapacity1Choice struct {

	// Trading capacity expressed as an ISO 20022 code.
	Code *TradingCapacity4Code `xml:"Cd"`

	// Trading capacity expressed as a proprietary code.
	Proprietary *GenericIdentification38 `xml:"Prtry"`
}

func (t *TradingPartyCapacity1Choice) SetCode(value string) {
	t.Code = (*TradingCapacity4Code)(&value)
}

func (t *TradingPartyCapacity1Choice) AddProprietary() *GenericIdentification38 {
	t.Proprietary = new(GenericIdentification38)
	return t.Proprietary
}
