package model

// Scope of the modification to be applied on an identified set of information.
type ModificationScope26 struct {

	// Specifies the type of modification to be applied.
	ModificationScopeIndication *DataModification1Code `xml:"ModScpIndctn"`

	// Intermediary or other party related to the management of the account. In some markets, when this intermediary is a party acting on behalf of the investor for which it has opened an account at, for example, a central securities depository or international central securities depository, this party is known by the investor as the 'account controller'.
	Intermediary *Intermediary36 `xml:"Intrmy"`
}

func (m *ModificationScope26) SetModificationScopeIndication(value string) {
	m.ModificationScopeIndication = (*DataModification1Code)(&value)
}

func (m *ModificationScope26) AddIntermediary() *Intermediary36 {
	m.Intermediary = new(Intermediary36)
	return m.Intermediary
}
