package model

// Choice of format for the market type.
type MarketType15Choice struct {

	// Market type expressed as an ISO 20022 code.
	Code *MarketType4Code `xml:"Cd"`

	// Market type expressed as a proprietary code.
	Proprietary *GenericIdentification30 `xml:"Prtry"`
}

func (m *MarketType15Choice) SetCode(value string) {
	m.Code = (*MarketType4Code)(&value)
}

func (m *MarketType15Choice) AddProprietary() *GenericIdentification30 {
	m.Proprietary = new(GenericIdentification30)
	return m.Proprietary
}
