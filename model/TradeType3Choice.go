package model

// Choice of format for the trade type.
type TradeType3Choice struct {

	// Trade type information expressed as an ISO 20022 code.
	Code *TradeType3Code `xml:"Cd"`

	// Third party reporting information expressed as a proprietary code.
	Proprietary *GenericIdentification38 `xml:"Prtry"`
}

func (t *TradeType3Choice) SetCode(value string) {
	t.Code = (*TradeType3Code)(&value)
}

func (t *TradeType3Choice) AddProprietary() *GenericIdentification38 {
	t.Proprietary = new(GenericIdentification38)
	return t.Proprietary
}
