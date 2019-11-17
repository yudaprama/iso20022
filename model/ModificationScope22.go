package model

// Scope of the modification to be applied on an identified set of information.
type ModificationScope22 struct {

	// Specifies the type of modification to be applied on a set of information.
	ModificationScopeIndication *DataModification1Code `xml:"ModScpIndctn"`

	// Identification of information which is part of a service level agreement.
	ServiceLevelAgreement *DocumentToSend2 `xml:"SvcLvlAgrmt"`
}

func (m *ModificationScope22) SetModificationScopeIndication(value string) {
	m.ModificationScopeIndication = (*DataModification1Code)(&value)
}

func (m *ModificationScope22) AddServiceLevelAgreement() *DocumentToSend2 {
	m.ServiceLevelAgreement = new(DocumentToSend2)
	return m.ServiceLevelAgreement
}
