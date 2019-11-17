package model

// Choice of format for the market type.
type MarketType2Choice struct {

	// Market type expressed as an ISO 20022 code.
	Code *MarketType5Code `xml:"Cd"`

	// Market type expressed as a proprietary code.
	Proprietary *GenericIdentification20 `xml:"Prtry"`
}

func (m *MarketType2Choice) SetCode(value string) {
	m.Code = (*MarketType5Code)(&value)
}

func (m *MarketType2Choice) AddProprietary() *GenericIdentification20 {
	m.Proprietary = new(GenericIdentification20)
	return m.Proprietary
}
