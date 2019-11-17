package model

// Information related to the request of maintenance delegations.
type MaintenanceDelegationRequest2 struct {

	// Terminal manager identification.
	TMIdentification *GenericIdentification72 `xml:"TMId"`

	// Master terminal manager identification.
	MasterTMIdentification *GenericIdentification72 `xml:"MstrTMId,omitempty"`

	// Information on the delegation of a maintenance action.
	RequestedDelegation []*MaintenanceDelegation3 `xml:"ReqdDlgtn"`
}

func (m *MaintenanceDelegationRequest2) AddTMIdentification() *GenericIdentification72 {
	m.TMIdentification = new(GenericIdentification72)
	return m.TMIdentification
}

func (m *MaintenanceDelegationRequest2) AddMasterTMIdentification() *GenericIdentification72 {
	m.MasterTMIdentification = new(GenericIdentification72)
	return m.MasterTMIdentification
}

func (m *MaintenanceDelegationRequest2) AddRequestedDelegation() *MaintenanceDelegation3 {
	newValue := new(MaintenanceDelegation3)
	m.RequestedDelegation = append(m.RequestedDelegation, newValue)
	return newValue
}
