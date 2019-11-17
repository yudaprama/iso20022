package model

// Information on the delegation of a maintenance action or maintenance function.
type MaintenanceDelegation1 struct {

	// Maintenance service to be delegated.
	MaintenanceService []*DataSetCategory6Code `xml:"MntncSvc"`

	// Flag to indicate that the delegated maintenance must be performed on a subset of the terminal estate.
	PartialDelegation *TrueFalseIndicator `xml:"PrtlDlgtn,omitempty"`

	// Subset of the terminal estate for the delegated actions, for instance for pilot or key deactivation). The subset may be expressed as a list of POI or terminal estate subset identifier.
	POISubset []*Max35Text `xml:"POISubset,omitempty"`

	// Information for the MTM to build or include delegated actions in the management plan of the POI.
	DelegatedAction *MaintenanceDelegateAction1 `xml:"DlgtdActn,omitempty"`

	// Identification of the parameters subset assigned by the MTM.
	ParametersSubsetIdentification *Max35Text `xml:"ParamsSubsetId,omitempty"`

	// Definition of the subset of application parameters, for instance the range of application profiles, the RID (registered application provider identification).
	ParametersSubsetDefinition *Max3000Binary `xml:"ParamsSubsetDef,omitempty"`

	// Certificate path of the terminal manager.
	Certificate []*Max5000Binary `xml:"Cert,omitempty"`

	// Association of the TM identifier and the MTM identifier of a POI.
	POIIdentificationAssociation []*MaintenanceIdentificationAssociation1 `xml:"POIIdAssoctn,omitempty"`

	// Identification of the key to manage or to download.
	SymmetricKey []*KEKIdentifier2 `xml:"SmmtrcKey,omitempty"`

	// Configuration parameters of the terminal manager to be sent by the MTM.
	ParameterDataSet *TerminalManagementDataSet14 `xml:"ParamDataSet,omitempty"`
}

func (m *MaintenanceDelegation1) AddMaintenanceService(value string) {
	m.MaintenanceService = append(m.MaintenanceService, (*DataSetCategory6Code)(&value))
}

func (m *MaintenanceDelegation1) SetPartialDelegation(value string) {
	m.PartialDelegation = (*TrueFalseIndicator)(&value)
}

func (m *MaintenanceDelegation1) AddPOISubset(value string) {
	m.POISubset = append(m.POISubset, (*Max35Text)(&value))
}

func (m *MaintenanceDelegation1) AddDelegatedAction() *MaintenanceDelegateAction1 {
	m.DelegatedAction = new(MaintenanceDelegateAction1)
	return m.DelegatedAction
}

func (m *MaintenanceDelegation1) SetParametersSubsetIdentification(value string) {
	m.ParametersSubsetIdentification = (*Max35Text)(&value)
}

func (m *MaintenanceDelegation1) SetParametersSubsetDefinition(value string) {
	m.ParametersSubsetDefinition = (*Max3000Binary)(&value)
}

func (m *MaintenanceDelegation1) AddCertificate(value string) {
	m.Certificate = append(m.Certificate, (*Max5000Binary)(&value))
}

func (m *MaintenanceDelegation1) AddPOIIdentificationAssociation() *MaintenanceIdentificationAssociation1 {
	newValue := new(MaintenanceIdentificationAssociation1)
	m.POIIdentificationAssociation = append(m.POIIdentificationAssociation, newValue)
	return newValue
}

func (m *MaintenanceDelegation1) AddSymmetricKey() *KEKIdentifier2 {
	newValue := new(KEKIdentifier2)
	m.SymmetricKey = append(m.SymmetricKey, newValue)
	return newValue
}

func (m *MaintenanceDelegation1) AddParameterDataSet() *TerminalManagementDataSet14 {
	m.ParameterDataSet = new(TerminalManagementDataSet14)
	return m.ParameterDataSet
}
