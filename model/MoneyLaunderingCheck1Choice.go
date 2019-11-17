package model

// Choice of formats for the specification of the money laundering check.
type MoneyLaunderingCheck1Choice struct {

	// Money laundering status expressed as a code.
	Code *MoneyLaunderingCheck1Code `xml:"Cd"`

	// Money laundering status expressed as a proprietary code.
	Proprietary *GenericIdentification47 `xml:"Prtry"`
}

func (m *MoneyLaunderingCheck1Choice) SetCode(value string) {
	m.Code = (*MoneyLaunderingCheck1Code)(&value)
}

func (m *MoneyLaunderingCheck1Choice) AddProprietary() *GenericIdentification47 {
	m.Proprietary = new(GenericIdentification47)
	return m.Proprietary
}
