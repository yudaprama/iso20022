package model

// Scope of the modification to be applied on an identified set of information.
type ModificationScope1 struct {

	// Specifies the type of modification to be applied on a set of information.
	ModificationScopeIndication *DataModification1Code `xml:"ModScpIndctn"`

	// Information that locates and identifies a specific address, as defined by postal services.
	PostalAddress *PostalAddress3 `xml:"PstlAdr"`
}

func (m *ModificationScope1) SetModificationScopeIndication(value string) {
	m.ModificationScopeIndication = (*DataModification1Code)(&value)
}

func (m *ModificationScope1) AddPostalAddress() *PostalAddress3 {
	m.PostalAddress = new(PostalAddress3)
	return m.PostalAddress
}
