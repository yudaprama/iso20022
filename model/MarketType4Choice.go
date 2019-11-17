package model

// Choice of format for the market type.
type MarketType4Choice struct {

	// Market type expressed as an ISO 20022 code.
	Code *MarketType4Code `xml:"Cd"`

	// Market type expressed as a proprietary code.
	Proprietary *GenericIdentification20 `xml:"Prtry"`
}

func (m *MarketType4Choice) SetCode(value string) {
	m.Code = (*MarketType4Code)(&value)
}

func (m *MarketType4Choice) AddProprietary() *GenericIdentification20 {
	m.Proprietary = new(GenericIdentification20)
	return m.Proprietary
}
