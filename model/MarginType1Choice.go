package model

// Choice of formats to specify the type margin.
type MarginType1Choice struct {

	// Provides the margin type using a code.
	Code *MarginType1Code `xml:"Cd"`

	// Provides the margin type using a proprietary format.
	Proprietary *GenericIdentification30 `xml:"Prtry"`
}

func (m *MarginType1Choice) SetCode(value string) {
	m.Code = (*MarginType1Code)(&value)
}

func (m *MarginType1Choice) AddProprietary() *GenericIdentification30 {
	m.Proprietary = new(GenericIdentification30)
	return m.Proprietary
}
