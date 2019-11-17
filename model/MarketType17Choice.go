package model

// Choice of format for the market type.
type MarketType17Choice struct {

	// Market type expressed as an ISO 20022 code.
	Code *MarketType4Code `xml:"Cd"`

	// Market type expressed as a proprietary code.
	Proprietary *GenericIdentification47 `xml:"Prtry"`
}

func (m *MarketType17Choice) SetCode(value string) {
	m.Code = (*MarketType4Code)(&value)
}

func (m *MarketType17Choice) AddProprietary() *GenericIdentification47 {
	m.Proprietary = new(GenericIdentification47)
	return m.Proprietary
}
