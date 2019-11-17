package model

// Scope of the modification to be applied on an identified set of information.
type ModificationScope38 struct {

	// Type of modification to be applied.
	ModificationScopeIndication *DataModification1Code `xml:"ModScpIndctn"`

	// Intermediary or other party related to the management of the account.
	Intermediary *Intermediary36 `xml:"Intrmy"`
}

func (m *ModificationScope38) SetModificationScopeIndication(value string) {
	m.ModificationScopeIndication = (*DataModification1Code)(&value)
}

func (m *ModificationScope38) AddIntermediary() *Intermediary36 {
	m.Intermediary = new(Intermediary36)
	return m.Intermediary
}
