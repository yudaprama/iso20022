package model

// Scope of the modification to be applied on an identified set of information.
type ModificationScope3 struct {

	// Specifies the type of modification to be applied on a set of information.
	ModificationScopeIndication *DataModification2Code `xml:"ModScpIndctn"`

	// Information about the nationality and the legal status (minor or major) of a person.
	Citizenship *CitizenshipInformation `xml:"Ctznsh"`
}

func (m *ModificationScope3) SetModificationScopeIndication(value string) {
	m.ModificationScopeIndication = (*DataModification2Code)(&value)
}

func (m *ModificationScope3) AddCitizenship() *CitizenshipInformation {
	m.Citizenship = new(CitizenshipInformation)
	return m.Citizenship
}
