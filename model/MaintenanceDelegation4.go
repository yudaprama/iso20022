package model

// Information on the delegation of a maintenance action or maintenance function.
type MaintenanceDelegation4 struct {

	// Maintenance service to be delegated.
	MaintenanceService []*DataSetCategory6Code `xml:"MntncSvc"`

	// Response of the MTM to the delegation of the maintenance service.
	Response *Response2Code `xml:"Rspn"`

	// Reason of the response of the MTM.
	ResponseReason *Max35Text `xml:"RspnRsn,omitempty"`

	// Type of delegation action.
	DelegationType *TerminalManagementAction3Code `xml:"DlgtnTp"`

	// Subset of the terminal estate for the delegated actions, for instance for pilot or key deactivation). The subset may be expressed as a list of POI or terminal estate subset identifier.
	POISubset []*Max35Text `xml:"POISubset,omitempty"`

	// Identification of the parameters subset assigned by the MTM.
	DelegationScopeIdentification *Max35Text `xml:"DlgtnScpId,omitempty"`

	// Definition of the delegation scope, for instance inside the payment application parameters the range of application profiles, the RID (registered application provider identification).
	DelegationScopeDefinition *Max3000Binary `xml:"DlgtnScpDef,omitempty"`

	// Proof of delegation to be verified by the POI, when performing the delegated actions.
	DelegationProof *Max5000Binary `xml:"DlgtnProof,omitempty"`

	// Protected proof of delegation.
	ProtectedDelegationProof *ContentInformationType12 `xml:"PrtctdDlgtnProof,omitempty"`

	// Association of the TM identifier and the MTM identifier of a POI.
	POIIdentificationAssociation []*MaintenanceIdentificationAssociation1 `xml:"POIIdAssoctn,omitempty"`
}

func (m *MaintenanceDelegation4) AddMaintenanceService(value string) {
	m.MaintenanceService = append(m.MaintenanceService, (*DataSetCategory6Code)(&value))
}

func (m *MaintenanceDelegation4) SetResponse(value string) {
	m.Response = (*Response2Code)(&value)
}

func (m *MaintenanceDelegation4) SetResponseReason(value string) {
	m.ResponseReason = (*Max35Text)(&value)
}

func (m *MaintenanceDelegation4) SetDelegationType(value string) {
	m.DelegationType = (*TerminalManagementAction3Code)(&value)
}

func (m *MaintenanceDelegation4) AddPOISubset(value string) {
	m.POISubset = append(m.POISubset, (*Max35Text)(&value))
}

func (m *MaintenanceDelegation4) SetDelegationScopeIdentification(value string) {
	m.DelegationScopeIdentification = (*Max35Text)(&value)
}

func (m *MaintenanceDelegation4) SetDelegationScopeDefinition(value string) {
	m.DelegationScopeDefinition = (*Max3000Binary)(&value)
}

func (m *MaintenanceDelegation4) SetDelegationProof(value string) {
	m.DelegationProof = (*Max5000Binary)(&value)
}

func (m *MaintenanceDelegation4) AddProtectedDelegationProof() *ContentInformationType12 {
	m.ProtectedDelegationProof = new(ContentInformationType12)
	return m.ProtectedDelegationProof
}

func (m *MaintenanceDelegation4) AddPOIIdentificationAssociation() *MaintenanceIdentificationAssociation1 {
	newValue := new(MaintenanceIdentificationAssociation1)
	m.POIIdentificationAssociation = append(m.POIIdentificationAssociation, newValue)
	return newValue
}
