package model

// Association of the TM identifier and the MTM identifier of an entity.
type MaintenanceIdentificationAssociation1 struct {

	// Identifier for the master terminal manager.
	MasterTMIdentification *Max35Text `xml:"MstrTMId"`

	// Identifier for the terminal manager requesting the delegation.
	TMIdentification *Max35Text `xml:"TMId"`
}

func (m *MaintenanceIdentificationAssociation1) SetMasterTMIdentification(value string) {
	m.MasterTMIdentification = (*Max35Text)(&value)
}

func (m *MaintenanceIdentificationAssociation1) SetTMIdentification(value string) {
	m.TMIdentification = (*Max35Text)(&value)
}
