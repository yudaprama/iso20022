package model

// Choice of format for the market/client side information.
type MarketClientSide3Choice struct {

	// Market side or a client side information expressed as an ISO 20022 code.
	Code *MarketClientSideCode `xml:"Cd"`

	// Market side or a client side information expressed as a proprietary code.
	Proprietary *GenericIdentification38 `xml:"Prtry"`
}

func (m *MarketClientSide3Choice) SetCode(value string) {
	m.Code = (*MarketClientSideCode)(&value)
}

func (m *MarketClientSide3Choice) AddProprietary() *GenericIdentification38 {
	m.Proprietary = new(GenericIdentification38)
	return m.Proprietary
}
