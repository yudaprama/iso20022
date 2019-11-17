package model

// Scope of the modification to be applied on an identified set of information.
type ModificationScope7 struct {

	// Specifies the type of modification to be applied on a set of information.
	ModificationScopeIndication *DataModification1Code `xml:"ModScpIndctn"`

	// Information related to intermediaries.
	Intermediary *Intermediary13 `xml:"Intrmy"`
}

func (m *ModificationScope7) SetModificationScopeIndication(value string) {
	m.ModificationScopeIndication = (*DataModification1Code)(&value)
}

func (m *ModificationScope7) AddIntermediary() *Intermediary13 {
	m.Intermediary = new(Intermediary13)
	return m.Intermediary
}
