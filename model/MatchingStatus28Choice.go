package model

// Choice of format for the matching status.
type MatchingStatus28Choice struct {

	// Provides the matching status of the instruction.
	Code *MatchingStatus1Code `xml:"Cd"`

	// Provides the matching status of the instruction.
	Proprietary *GenericIdentification47 `xml:"Prtry"`
}

func (m *MatchingStatus28Choice) SetCode(value string) {
	m.Code = (*MatchingStatus1Code)(&value)
}

func (m *MatchingStatus28Choice) AddProprietary() *GenericIdentification47 {
	m.Proprietary = new(GenericIdentification47)
	return m.Proprietary
}
