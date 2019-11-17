package model

// Choice of format for the matching status.
type MatchingStatus27Choice struct {

	// Provides the matching status of the instruction.
	Code *MatchingStatus1Code `xml:"Cd"`

	// Provides the matching status of the instruction.
	Proprietary *GenericIdentification30 `xml:"Prtry"`
}

func (m *MatchingStatus27Choice) SetCode(value string) {
	m.Code = (*MatchingStatus1Code)(&value)
}

func (m *MatchingStatus27Choice) AddProprietary() *GenericIdentification30 {
	m.Proprietary = new(GenericIdentification30)
	return m.Proprietary
}
