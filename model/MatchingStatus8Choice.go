package model

// Choice of format for the matching status.
type MatchingStatus8Choice struct {

	// Provides the matching status of the instruction.
	Code *MatchingStatus1Code `xml:"Cd"`

	// Provides the matching status of the instruction.
	Proprietary *GenericIdentification38 `xml:"Prtry"`
}

func (m *MatchingStatus8Choice) SetCode(value string) {
	m.Code = (*MatchingStatus1Code)(&value)
}

func (m *MatchingStatus8Choice) AddProprietary() *GenericIdentification38 {
	m.Proprietary = new(GenericIdentification38)
	return m.Proprietary
}
