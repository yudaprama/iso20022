package model

// Choice of format for the trade date code.
type TradeDateCode1Choice struct {

	// Trade date expressed as an ISO 20022 code.
	Code *DateType3Code `xml:"Cd"`

	// Trade date expressed as an proprietary code.
	Proprietary *GenericIdentification20 `xml:"Prtry"`
}

func (t *TradeDateCode1Choice) SetCode(value string) {
	t.Code = (*DateType3Code)(&value)
}

func (t *TradeDateCode1Choice) AddProprietary() *GenericIdentification20 {
	t.Proprietary = new(GenericIdentification20)
	return t.Proprietary
}
