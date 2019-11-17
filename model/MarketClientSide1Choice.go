package model

// Choice of format for the market/client side information.
type MarketClientSide1Choice struct {

	// Market side or a client side information expressed as an ISO 20022 code.
	Code *MarketClientSideCode `xml:"Cd"`

	// Market side or a client side information expressed as a proprietary code.
	Proprietary *GenericIdentification20 `xml:"Prtry"`
}

func (m *MarketClientSide1Choice) SetCode(value string) {
	m.Code = (*MarketClientSideCode)(&value)
}

func (m *MarketClientSide1Choice) AddProprietary() *GenericIdentification20 {
	m.Proprietary = new(GenericIdentification20)
	return m.Proprietary
}
