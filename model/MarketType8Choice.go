package model

// Choice of format for the market type.
type MarketType8Choice struct {

	// Market type expressed as an ISO 20022 code.
	Code *MarketType2Code `xml:"Cd"`

	// Market type expressed as a proprietary code.
	Proprietary *GenericIdentification30 `xml:"Prtry"`
}

func (m *MarketType8Choice) SetCode(value string) {
	m.Code = (*MarketType2Code)(&value)
}

func (m *MarketType8Choice) AddProprietary() *GenericIdentification30 {
	m.Proprietary = new(GenericIdentification30)
	return m.Proprietary
}
