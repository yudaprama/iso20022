package model

// Scope of the modification to be applied on an identified set of information.
type ModificationScope10 struct {

	// Specifies the type of modification to be applied on a set of information.
	ModificationScopeIndication *DataModification1Code `xml:"ModScpIndctn"`

	// Identification of information which is part of a service level agreement.
	ServiceLevelAgreement *DocumentToSend1 `xml:"SvcLvlAgrmt"`
}

func (m *ModificationScope10) SetModificationScopeIndication(value string) {
	m.ModificationScopeIndication = (*DataModification1Code)(&value)
}

func (m *ModificationScope10) AddServiceLevelAgreement() *DocumentToSend1 {
	m.ServiceLevelAgreement = new(DocumentToSend1)
	return m.ServiceLevelAgreement
}
