package model

// Choice of format for the trading date code.
type TradingDateCode1Choice struct {

	// Trading date expressed as a ISO20022 code.
	Code *TradingDate1Code `xml:"Cd"`

	// Trading date expressed as a proprietary code.
	Proprietary *GenericIdentification38 `xml:"Prtry"`
}

func (t *TradingDateCode1Choice) SetCode(value string) {
	t.Code = (*TradingDate1Code)(&value)
}

func (t *TradingDateCode1Choice) AddProprietary() *GenericIdentification38 {
	t.Proprietary = new(GenericIdentification38)
	return t.Proprietary
}
