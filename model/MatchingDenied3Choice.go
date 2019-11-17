package model

// Specifies the matching processing change requested.
type MatchingDenied3Choice struct {

	// Specifies the execution of a matching denial process.
	Code *MatchingProcess1Code `xml:"Cd"`

	// Specifies the execution of a matching denial process.
	Proprietary *GenericIdentification30 `xml:"Prtry"`
}

func (m *MatchingDenied3Choice) SetCode(value string) {
	m.Code = (*MatchingProcess1Code)(&value)
}

func (m *MatchingDenied3Choice) AddProprietary() *GenericIdentification30 {
	m.Proprietary = new(GenericIdentification30)
	return m.Proprietary
}
