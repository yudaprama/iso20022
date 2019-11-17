package model

// Information on the delegation of a maintenance action or maintenance function.
type MaintenanceDelegation3 struct {

	// Type of delegation action.
	DelegationType *TerminalManagementAction3Code `xml:"DlgtnTp"`

	// Maintenance service to be delegated.
	MaintenanceService []*DataSetCategory11Code `xml:"MntncSvc"`

	// Flag to indicate that the delegated maintenance must be performed on a subset of the terminal estate.
	PartialDelegation *TrueFalseIndicator `xml:"PrtlDlgtn,omitempty"`

	// Subset of the terminal estate for the delegated actions, for instance for pilot or key deactivation). The subset may be expressed as a list of POI or terminal estate subset identifier.
	POISubset []*Max35Text `xml:"POISubset,omitempty"`

	// Information for the MTM to build or include delegated actions in the management plan of the POI.
	DelegatedAction *MaintenanceDelegateAction2 `xml:"DlgtdActn,omitempty"`

	// Identification of the delegation scope assigned by the MTM.
	DelegationScopeIdentification *Max35Text `xml:"DlgtnScpId,omitempty"`

	// Definition of the delegation scope, for instance inside the payment application parameters the range of application profiles, the RID (registered application provider identification).
	DelegationScopeDefinition *Max3000Binary `xml:"DlgtnScpDef,omitempty"`

	// Certificate path of the terminal manager.
	Certificate []*Max10KBinary `xml:"Cert,omitempty"`

	// Association of the TM identifier and the MTM identifier of a POI.
	POIIdentificationAssociation []*MaintenanceIdentificationAssociation1 `xml:"POIIdAssoctn,omitempty"`

	// Identification of the key to manage or to download.
	SymmetricKey []*KEKIdentifier5 `xml:"SmmtrcKey,omitempty"`

	// Configuration parameters of the terminal manager to be sent by the MTM.
	ParameterDataSet *TerminalManagementDataSet19 `xml:"ParamDataSet,omitempty"`
}

func (m *MaintenanceDelegation3) SetDelegationType(value string) {
	m.DelegationType = (*TerminalManagementAction3Code)(&value)
}

func (m *MaintenanceDelegation3) AddMaintenanceService(value string) {
	m.MaintenanceService = append(m.MaintenanceService, (*DataSetCategory11Code)(&value))
}

func (m *MaintenanceDelegation3) SetPartialDelegation(value string) {
	m.PartialDelegation = (*TrueFalseIndicator)(&value)
}

func (m *MaintenanceDelegation3) AddPOISubset(value string) {
	m.POISubset = append(m.POISubset, (*Max35Text)(&value))
}

func (m *MaintenanceDelegation3) AddDelegatedAction() *MaintenanceDelegateAction2 {
	m.DelegatedAction = new(MaintenanceDelegateAction2)
	return m.DelegatedAction
}

func (m *MaintenanceDelegation3) SetDelegationScopeIdentification(value string) {
	m.DelegationScopeIdentification = (*Max35Text)(&value)
}

func (m *MaintenanceDelegation3) SetDelegationScopeDefinition(value string) {
	m.DelegationScopeDefinition = (*Max3000Binary)(&value)
}

func (m *MaintenanceDelegation3) AddCertificate(value string) {
	m.Certificate = append(m.Certificate, (*Max10KBinary)(&value))
}

func (m *MaintenanceDelegation3) AddPOIIdentificationAssociation() *MaintenanceIdentificationAssociation1 {
	newValue := new(MaintenanceIdentificationAssociation1)
	m.POIIdentificationAssociation = append(m.POIIdentificationAssociation, newValue)
	return newValue
}

func (m *MaintenanceDelegation3) AddSymmetricKey() *KEKIdentifier5 {
	newValue := new(KEKIdentifier5)
	m.SymmetricKey = append(m.SymmetricKey, newValue)
	return newValue
}

func (m *MaintenanceDelegation3) AddParameterDataSet() *TerminalManagementDataSet19 {
	m.ParameterDataSet = new(TerminalManagementDataSet19)
	return m.ParameterDataSet
}
