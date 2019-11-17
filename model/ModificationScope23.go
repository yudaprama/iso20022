package model

// Specifies the type of modification to be applied on a data set.
type ModificationScope23 struct {

	// Specifies the type of modification to be applied on a set of information.
	ModificationScopeIndication *DataModification2Code `xml:"ModScpIndctn"`

	// Alternative identification, for example, national registration identification number, passport number, or an account number used to further identify the beneficial owner, for example, a Central Provident Fund (CFP) account as required for Singapore.
	OtherIdentification *GenericIdentification55 `xml:"OthrId"`
}

func (m *ModificationScope23) SetModificationScopeIndication(value string) {
	m.ModificationScopeIndication = (*DataModification2Code)(&value)
}

func (m *ModificationScope23) AddOtherIdentification() *GenericIdentification55 {
	m.OtherIdentification = new(GenericIdentification55)
	return m.OtherIdentification
}
