package model

// Specifies the matching processing change requested.
type MatchingDenied4Choice struct {

	// Specifies the execution of a matching denial process.
	Code *MatchingProcess1Code `xml:"Cd"`

	// Specifies the execution of a matching denial process.
	Proprietary *GenericIdentification47 `xml:"Prtry"`
}

func (m *MatchingDenied4Choice) SetCode(value string) {
	m.Code = (*MatchingProcess1Code)(&value)
}

func (m *MatchingDenied4Choice) AddProprietary() *GenericIdentification47 {
	m.Proprietary = new(GenericIdentification47)
	return m.Proprietary
}
