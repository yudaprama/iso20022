package model

// Scope of the modification to be applied on an identified set of information.
type ModificationScope2 struct {

	// Specifies the type of modification to be applied on a set of information.
	ModificationScopeIndication *DataModification2Code `xml:"ModScpIndctn"`

	// Alternative identification, for example, national registration identification number, passport number, or an account number used to further identify the beneficial owner, for example, a Central Provident Fund (CFP) account as required for Singapore.
	OtherIdentification *GenericIdentification9 `xml:"OthrId"`
}

func (m *ModificationScope2) SetModificationScopeIndication(value string) {
	m.ModificationScopeIndication = (*DataModification2Code)(&value)
}

func (m *ModificationScope2) AddOtherIdentification() *GenericIdentification9 {
	m.OtherIdentification = new(GenericIdentification9)
	return m.OtherIdentification
}
