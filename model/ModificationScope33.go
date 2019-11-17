package model

// Scope of the modification to be applied on an identified set of information.
type ModificationScope33 struct {

	// Specifies the type of modification to be applied.
	ModificationScopeIndication *DataModification1Code `xml:"ModScpIndctn"`

	// Referral information.
	Placement *ReferredAgent2 `xml:"Plcmnt"`
}

func (m *ModificationScope33) SetModificationScopeIndication(value string) {
	m.ModificationScopeIndication = (*DataModification1Code)(&value)
}

func (m *ModificationScope33) AddPlacement() *ReferredAgent2 {
	m.Placement = new(ReferredAgent2)
	return m.Placement
}
