package model

// Scope of the modification to be applied on an identified set of information.
type ModificationScope20 struct {

	// Specifies the type of modification to be applied on a set of information.
	ModificationScopeIndication *DataModification1Code `xml:"ModScpIndctn"`

	// Information related to intermediaries.
	Intermediary *Intermediary24 `xml:"Intrmy"`
}

func (m *ModificationScope20) SetModificationScopeIndication(value string) {
	m.ModificationScopeIndication = (*DataModification1Code)(&value)
}

func (m *ModificationScope20) AddIntermediary() *Intermediary24 {
	m.Intermediary = new(Intermediary24)
	return m.Intermediary
}
