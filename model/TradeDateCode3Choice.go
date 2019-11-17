package model

// Choice of format for the trade date code.
type TradeDateCode3Choice struct {

	// Trade date expressed as an ISO 20022 code.
	Code *DateType3Code `xml:"Cd"`

	// Trade date expressed as an proprietary code.
	Proprietary *GenericIdentification30 `xml:"Prtry"`
}

func (t *TradeDateCode3Choice) SetCode(value string) {
	t.Code = (*DateType3Code)(&value)
}

func (t *TradeDateCode3Choice) AddProprietary() *GenericIdentification30 {
	t.Proprietary = new(GenericIdentification30)
	return t.Proprietary
}
