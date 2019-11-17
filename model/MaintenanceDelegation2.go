package model

// Information on the delegation of a maintenance action or maintenance function.
type MaintenanceDelegation2 struct {

	// Maintenance service to be delegated.
	MaintenanceService []*DataSetCategory6Code `xml:"MntncSvc"`

	// Response of the MTM to the delegation of the maintenance service.
	Response *Response2Code `xml:"Rspn"`

	// Reason of the response of the MTM.
	ResponseReason *Max35Text `xml:"RspnRsn,omitempty"`

	// Subset of the terminal estate for the delegated actions, for instance for pilot or key deactivation). The subset may be expressed as a list of POI or terminal estate subset identifier.
	POISubset []*Max35Text `xml:"POISubset,omitempty"`

	// Identification of the parameters subset assigned by the MTM.
	ParametersSubsetIdentification *Max35Text `xml:"ParamsSubsetId,omitempty"`

	// Definition of the subset of application parameters, for instance the range of application profiles, the RID (registered application provider identification).
	ParametersSubsetDefinition *Max3000Binary `xml:"ParamsSubsetDef,omitempty"`

	// Proof of delegation to be verified by the POI, when performing the delegated actions.
	DelegationProof *Max5000Binary `xml:"DlgtnProof,omitempty"`

	// Protected proof of delegation.
	ProtectedDelegationProof *ContentInformationType12 `xml:"PrtctdDlgtnProof,omitempty"`

	// Association of the TM identifier and the MTM identifier of a POI.
	POIIdentificationAssociation []*MaintenanceIdentificationAssociation1 `xml:"POIIdAssoctn,omitempty"`
}

func (m *MaintenanceDelegation2) AddMaintenanceService(value string) {
	m.MaintenanceService = append(m.MaintenanceService, (*DataSetCategory6Code)(&value))
}

func (m *MaintenanceDelegation2) SetResponse(value string) {
	m.Response = (*Response2Code)(&value)
}

func (m *MaintenanceDelegation2) SetResponseReason(value string) {
	m.ResponseReason = (*Max35Text)(&value)
}

func (m *MaintenanceDelegation2) AddPOISubset(value string) {
	m.POISubset = append(m.POISubset, (*Max35Text)(&value))
}

func (m *MaintenanceDelegation2) SetParametersSubsetIdentification(value string) {
	m.ParametersSubsetIdentification = (*Max35Text)(&value)
}

func (m *MaintenanceDelegation2) SetParametersSubsetDefinition(value string) {
	m.ParametersSubsetDefinition = (*Max3000Binary)(&value)
}

func (m *MaintenanceDelegation2) SetDelegationProof(value string) {
	m.DelegationProof = (*Max5000Binary)(&value)
}

func (m *MaintenanceDelegation2) AddProtectedDelegationProof() *ContentInformationType12 {
	m.ProtectedDelegationProof = new(ContentInformationType12)
	return m.ProtectedDelegationProof
}

func (m *MaintenanceDelegation2) AddPOIIdentificationAssociation() *MaintenanceIdentificationAssociation1 {
	newValue := new(MaintenanceIdentificationAssociation1)
	m.POIIdentificationAssociation = append(m.POIIdentificationAssociation, newValue)
	return newValue
}
