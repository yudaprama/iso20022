package model

// Choice of format for the market type.
type MarketType9Choice struct {

	// Market type expressed as an ISO 20022 code.
	Code *MarketType5Code `xml:"Cd"`

	// Market type expressed as a proprietary code.
	Proprietary *GenericIdentification30 `xml:"Prtry"`
}

func (m *MarketType9Choice) SetCode(value string) {
	m.Code = (*MarketType5Code)(&value)
}

func (m *MarketType9Choice) AddProprietary() *GenericIdentification30 {
	m.Proprietary = new(GenericIdentification30)
	return m.Proprietary
}
