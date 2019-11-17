package model

// Information for the MTM to build or include delegated actions in the management plan of the POI.
type MaintenanceDelegateAction1 struct {

	// Flag to indicate that the delegated actions have to be included in a periodic sequence of actions.
	PeriodicAction *TrueFalseIndicator `xml:"PrdcActn,omitempty"`

	// Network address and parameters of the terminal manager host which will performs the delegated actions.
	TMRemoteAccess *NetworkParameters3 `xml:"TMRmotAccs,omitempty"`

	// TMS protocol to use to perform the maintenance action.
	TMSProtocol *Max35Text `xml:"TMSPrtcol,omitempty"`

	// Version of the TMS protocol to use to perform the maintenance action.
	TMSProtocolVersion *Max35Text `xml:"TMSPrtcolVrsn,omitempty"`

	// Data set  on which the delegated action has to be performed.
	DataSetIdentification *DataSetIdentification4 `xml:"DataSetId,omitempty"`

	// Definition of retry process when activation of the action fails.
	ReTry *ProcessRetry2 `xml:"ReTry,omitempty"`

	// Additional information to include in the maintenance action.
	AdditionalInformation []*Max3000Binary `xml:"AddtlInf,omitempty"`

	// Sequence of action to include in the next MTM management plan.
	Action []*TMSAction4 `xml:"Actn,omitempty"`
}

func (m *MaintenanceDelegateAction1) SetPeriodicAction(value string) {
	m.PeriodicAction = (*TrueFalseIndicator)(&value)
}

func (m *MaintenanceDelegateAction1) AddTMRemoteAccess() *NetworkParameters3 {
	m.TMRemoteAccess = new(NetworkParameters3)
	return m.TMRemoteAccess
}

func (m *MaintenanceDelegateAction1) SetTMSProtocol(value string) {
	m.TMSProtocol = (*Max35Text)(&value)
}

func (m *MaintenanceDelegateAction1) SetTMSProtocolVersion(value string) {
	m.TMSProtocolVersion = (*Max35Text)(&value)
}

func (m *MaintenanceDelegateAction1) AddDataSetIdentification() *DataSetIdentification4 {
	m.DataSetIdentification = new(DataSetIdentification4)
	return m.DataSetIdentification
}

func (m *MaintenanceDelegateAction1) AddReTry() *ProcessRetry2 {
	m.ReTry = new(ProcessRetry2)
	return m.ReTry
}

func (m *MaintenanceDelegateAction1) AddAdditionalInformation(value string) {
	m.AdditionalInformation = append(m.AdditionalInformation, (*Max3000Binary)(&value))
}

func (m *MaintenanceDelegateAction1) AddAction() *TMSAction4 {
	newValue := new(TMSAction4)
	m.Action = append(m.Action, newValue)
	return newValue
}
