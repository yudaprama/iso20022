package model

// Choice of format for the matching status.
type MatchingStatus1Choice struct {

	// Provides the matching status of the instruction.
	Code *MatchingStatus1Code `xml:"Cd"`

	// Provides the matching status of the instruction.
	Proprietary *GenericIdentification20 `xml:"Prtry"`
}

func (m *MatchingStatus1Choice) SetCode(value string) {
	m.Code = (*MatchingStatus1Code)(&value)
}

func (m *MatchingStatus1Choice) AddProprietary() *GenericIdentification20 {
	m.Proprietary = new(GenericIdentification20)
	return m.Proprietary
}
