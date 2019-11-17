package model

// Choice of format for the market type.
type MarketType16Choice struct {

	// Market type expressed as an ISO 20022 code.
	Code *MarketType2Code `xml:"Cd"`

	// Market type expressed as a proprietary code.
	Proprietary *GenericIdentification47 `xml:"Prtry"`
}

func (m *MarketType16Choice) SetCode(value string) {
	m.Code = (*MarketType2Code)(&value)
}

func (m *MarketType16Choice) AddProprietary() *GenericIdentification47 {
	m.Proprietary = new(GenericIdentification47)
	return m.Proprietary
}
