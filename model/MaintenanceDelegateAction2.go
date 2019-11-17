package model

// Information for the MTM to build or include delegated actions in the management plan of the POI.
type MaintenanceDelegateAction2 struct {

	// Flag to indicate that the delegated actions have to be included in a periodic sequence of actions.
	PeriodicAction *TrueFalseIndicator `xml:"PrdcActn,omitempty"`

	// Network address and parameters of the terminal manager host which will performs the delegated actions.
	TMRemoteAccess *NetworkParameters5 `xml:"TMRmotAccs,omitempty"`

	// TMS protocol to use to perform the maintenance action.
	TMSProtocol *Max35Text `xml:"TMSPrtcol,omitempty"`

	// Version of the TMS protocol to use to perform the maintenance action.
	TMSProtocolVersion *Max35Text `xml:"TMSPrtcolVrsn,omitempty"`

	// Data set  on which the delegated action has to be performed.
	DataSetIdentification *DataSetIdentification6 `xml:"DataSetId,omitempty"`

	// Definition of retry process when activation of the action fails.
	ReTry *ProcessRetry2 `xml:"ReTry,omitempty"`

	// Additional information to include in the maintenance action.
	AdditionalInformation []*Max3000Binary `xml:"AddtlInf,omitempty"`

	// Sequence of action to include in the next MTM management plan.
	Action []*TMSAction5 `xml:"Actn,omitempty"`
}

func (m *MaintenanceDelegateAction2) SetPeriodicAction(value string) {
	m.PeriodicAction = (*TrueFalseIndicator)(&value)
}

func (m *MaintenanceDelegateAction2) AddTMRemoteAccess() *NetworkParameters5 {
	m.TMRemoteAccess = new(NetworkParameters5)
	return m.TMRemoteAccess
}

func (m *MaintenanceDelegateAction2) SetTMSProtocol(value string) {
	m.TMSProtocol = (*Max35Text)(&value)
}

func (m *MaintenanceDelegateAction2) SetTMSProtocolVersion(value string) {
	m.TMSProtocolVersion = (*Max35Text)(&value)
}

func (m *MaintenanceDelegateAction2) AddDataSetIdentification() *DataSetIdentification6 {
	m.DataSetIdentification = new(DataSetIdentification6)
	return m.DataSetIdentification
}

func (m *MaintenanceDelegateAction2) AddReTry() *ProcessRetry2 {
	m.ReTry = new(ProcessRetry2)
	return m.ReTry
}

func (m *MaintenanceDelegateAction2) AddAdditionalInformation(value string) {
	m.AdditionalInformation = append(m.AdditionalInformation, (*Max3000Binary)(&value))
}

func (m *MaintenanceDelegateAction2) AddAction() *TMSAction5 {
	newValue := new(TMSAction5)
	m.Action = append(m.Action, newValue)
	return newValue
}
