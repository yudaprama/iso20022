package model

// Scope of the modification to be applied on an identified set of information.
type ModificationScope31 struct {

	// Specifies the type of modification to be applied.
	ModificationScopeIndication *DataModification1Code `xml:"ModScpIndctn"`

	// Identification of information which is part of a service level agreement.
	ServiceLevelAgreement *DocumentToSend3 `xml:"SvcLvlAgrmt"`
}

func (m *ModificationScope31) SetModificationScopeIndication(value string) {
	m.ModificationScopeIndication = (*DataModification1Code)(&value)
}

func (m *ModificationScope31) AddServiceLevelAgreement() *DocumentToSend3 {
	m.ServiceLevelAgreement = new(DocumentToSend3)
	return m.ServiceLevelAgreement
}
