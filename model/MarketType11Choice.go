package model

// Choice of format for the market type.
type MarketType11Choice struct {

	// Market type expressed as an ISO 20022 code.
	Code *MarketType6Code `xml:"Cd"`

	// Market type expressed as a proprietary code.
	Proprietary *GenericIdentification38 `xml:"Prtry"`
}

func (m *MarketType11Choice) SetCode(value string) {
	m.Code = (*MarketType6Code)(&value)
}

func (m *MarketType11Choice) AddProprietary() *GenericIdentification38 {
	m.Proprietary = new(GenericIdentification38)
	return m.Proprietary
}
