package model

// Choice of format for the market/client side information.
type MarketClientSide4Choice struct {

	// Market side or a client side information expressed as an ISO 20022 code.
	Code *MarketClientSideCode `xml:"Cd"`

	// Market side or a client side information expressed as a proprietary code.
	Proprietary *GenericIdentification30 `xml:"Prtry"`
}

func (m *MarketClientSide4Choice) SetCode(value string) {
	m.Code = (*MarketClientSideCode)(&value)
}

func (m *MarketClientSide4Choice) AddProprietary() *GenericIdentification30 {
	m.Proprietary = new(GenericIdentification30)
	return m.Proprietary
}
