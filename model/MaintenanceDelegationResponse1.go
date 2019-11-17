package model

// Information related to the request of maintenance delegations.
type MaintenanceDelegationResponse1 struct {

	// Terminal manager identification.
	TMIdentification *GenericIdentification72 `xml:"TMId"`

	// Master terminal manager identification.
	MasterTMIdentification *GenericIdentification72 `xml:"MstrTMId,omitempty"`

	// Information on the delegation of a maintenance action.
	DelegationResponse []*MaintenanceDelegation2 `xml:"DlgtnRspn"`
}

func (m *MaintenanceDelegationResponse1) AddTMIdentification() *GenericIdentification72 {
	m.TMIdentification = new(GenericIdentification72)
	return m.TMIdentification
}

func (m *MaintenanceDelegationResponse1) AddMasterTMIdentification() *GenericIdentification72 {
	m.MasterTMIdentification = new(GenericIdentification72)
	return m.MasterTMIdentification
}

func (m *MaintenanceDelegationResponse1) AddDelegationResponse() *MaintenanceDelegation2 {
	newValue := new(MaintenanceDelegation2)
	m.DelegationResponse = append(m.DelegationResponse, newValue)
	return newValue
}
