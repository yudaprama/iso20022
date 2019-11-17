package model

// Scope of the modification to be applied on an identified set of information.
type ModificationScope34 struct {

	// Specifies the type of modification to be applied on a set of information.
	ModificationScopeIndication *DataModification1Code `xml:"ModScpIndctn"`

	// Address of the organisation.
	PostalAddress *PostalAddress21 `xml:"PstlAdr"`
}

func (m *ModificationScope34) SetModificationScopeIndication(value string) {
	m.ModificationScopeIndication = (*DataModification1Code)(&value)
}

func (m *ModificationScope34) AddPostalAddress() *PostalAddress21 {
	m.PostalAddress = new(PostalAddress21)
	return m.PostalAddress
}
