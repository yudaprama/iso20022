package model

// Choice between a standard code or proprietary code to specify the type of market.
type MarketTypeFormat1Choice struct {

	// Standard code to specify the type of market in which transactions take place, for example, primary or secondary.
	Code *MarketType3Code `xml:"Cd"`

	// Proprietary identification of the type of market in which transactions take place.
	Proprietary *GenericIdentification20 `xml:"Prtry"`
}

func (m *MarketTypeFormat1Choice) SetCode(value string) {
	m.Code = (*MarketType3Code)(&value)
}

func (m *MarketTypeFormat1Choice) AddProprietary() *GenericIdentification20 {
	m.Proprietary = new(GenericIdentification20)
	return m.Proprietary
}
