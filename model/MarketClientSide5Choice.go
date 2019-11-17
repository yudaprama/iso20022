package model

// Choice of format for the market/client side information.
type MarketClientSide5Choice struct {

	// Market side or a client side information expressed as an ISO 20022 code.
	Code *MarketClientSideCode `xml:"Cd"`

	// Market side or a client side information expressed as a proprietary code.
	Proprietary *GenericIdentification47 `xml:"Prtry"`
}

func (m *MarketClientSide5Choice) SetCode(value string) {
	m.Code = (*MarketClientSideCode)(&value)
}

func (m *MarketClientSide5Choice) AddProprietary() *GenericIdentification47 {
	m.Proprietary = new(GenericIdentification47)
	return m.Proprietary
}
