package model

// Scope of the modification to be applied on an identified set of information.
type ModificationScope35 struct {

	// Type of modification to be applied.
	ModificationScopeIndication *DataModification1Code `xml:"ModScpIndctn"`

	// Additional information such as remarks or notes that must be conveyed about the party and or  limitations and restrictions.
	AdditionalInformation []*AdditiononalInformation12 `xml:"AddtlInf"`
}

func (m *ModificationScope35) SetModificationScopeIndication(value string) {
	m.ModificationScopeIndication = (*DataModification1Code)(&value)
}

func (m *ModificationScope35) AddAdditionalInformation() *AdditiononalInformation12 {
	newValue := new(AdditiononalInformation12)
	m.AdditionalInformation = append(m.AdditionalInformation, newValue)
	return newValue
}
