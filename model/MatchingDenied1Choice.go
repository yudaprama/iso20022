package model

// Specifies the matching processing change requested.
type MatchingDenied1Choice struct {

	// Specifies the execution of a matching denial process.
	Code *MatchingProcess1Code `xml:"Cd"`

	// Specifies the execution of a matching denial process.
	Proprietary *GenericIdentification20 `xml:"Prtry"`
}

func (m *MatchingDenied1Choice) SetCode(value string) {
	m.Code = (*MatchingProcess1Code)(&value)
}

func (m *MatchingDenied1Choice) AddProprietary() *GenericIdentification20 {
	m.Proprietary = new(GenericIdentification20)
	return m.Proprietary
}
