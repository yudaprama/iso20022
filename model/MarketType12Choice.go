package model

// Choice of format for the market type.
type MarketType12Choice struct {

	// Market type expressed as an ISO 20022 code.
	Code *MarketType2Code `xml:"Cd"`

	// Market type expressed as a proprietary code.
	Proprietary *GenericIdentification38 `xml:"Prtry"`
}

func (m *MarketType12Choice) SetCode(value string) {
	m.Code = (*MarketType2Code)(&value)
}

func (m *MarketType12Choice) AddProprietary() *GenericIdentification38 {
	m.Proprietary = new(GenericIdentification38)
	return m.Proprietary
}
