package model

// Choice of format for the market type.
type MarketType3Choice struct {

	// Market type expressed as an ISO 20022 code.
	Code *MarketType2Code `xml:"Cd"`

	// Market type expressed as a proprietary code.
	Proprietary *GenericIdentification20 `xml:"Prtry"`
}

func (m *MarketType3Choice) SetCode(value string) {
	m.Code = (*MarketType2Code)(&value)
}

func (m *MarketType3Choice) AddProprietary() *GenericIdentification20 {
	m.Proprietary = new(GenericIdentification20)
	return m.Proprietary
}
