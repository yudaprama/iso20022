package model

// Scope of the modification to be applied on an identified set of information.
type ModificationScope39 struct {

	// Specifies the type of modification to be applied on a set of information.
	ModificationScopeIndication *DataModification2Code `xml:"ModScpIndctn"`

	// Nationality and legal status (minor or major).
	Citizenship *CitizenshipInformation2 `xml:"Ctznsh"`
}

func (m *ModificationScope39) SetModificationScopeIndication(value string) {
	m.ModificationScopeIndication = (*DataModification2Code)(&value)
}

func (m *ModificationScope39) AddCitizenship() *CitizenshipInformation2 {
	m.Citizenship = new(CitizenshipInformation2)
	return m.Citizenship
}
