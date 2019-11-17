package model

// Provides specification for the underlying product of the margin.
type MarginProductType1Choice struct {

	// Specifies the underlying product of the margin using a code.
	Code *MarginProduct1Code `xml:"Cd"`

	// Specifies the underlying product of the margin using a proprietary format.
	Proprietary *GenericIdentification30 `xml:"Prtry"`
}

func (m *MarginProductType1Choice) SetCode(value string) {
	m.Code = (*MarginProduct1Code)(&value)
}

func (m *MarginProductType1Choice) AddProprietary() *GenericIdentification30 {
	m.Proprietary = new(GenericIdentification30)
	return m.Proprietary
}
