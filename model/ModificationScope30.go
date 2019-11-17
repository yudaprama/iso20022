package model

// Scope of the modification to be applied on an identified set of information.
type ModificationScope30 struct {

	// Specifies the type of modification to be applied.
	ModificationScopeIndication *DataModification1Code `xml:"ModScpIndctn"`

	// Additional information concerning limitations and restrictions on the account.
	AdditionalInformation []*AccountRestrictions1 `xml:"AddtlInf"`
}

func (m *ModificationScope30) SetModificationScopeIndication(value string) {
	m.ModificationScopeIndication = (*DataModification1Code)(&value)
}

func (m *ModificationScope30) AddAdditionalInformation() *AccountRestrictions1 {
	newValue := new(AccountRestrictions1)
	m.AdditionalInformation = append(m.AdditionalInformation, newValue)
	return newValue
}
