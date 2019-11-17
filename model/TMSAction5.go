package model

// Single terminal management action to be performed by the point of interaction.
type TMSAction5 struct {

	// Types of action to be performed by a point of interaction (POI).
	Type *TerminalManagementAction2Code `xml:"Tp"`

	// Host access information.
	RemoteAccess *NetworkParameters5 `xml:"RmotAccs,omitempty"`

	// Identification of the master terminal manager or the terminal manager with which the POI has to perform the action.
	TerminalManagerIdentification *GenericIdentification71 `xml:"TermnlMgrId,omitempty"`

	// TMS protocol to use for performing the maintenance action.
	TMSProtocol *Max35Text `xml:"TMSPrtcol,omitempty"`

	// Version of the TMS protocol to use to perform the maintenance action.
	TMSProtocolVersion *Max35Text `xml:"TMSPrtcolVrsn,omitempty"`

	// Data set on which the action has to be performed.
	DataSetIdentification *DataSetIdentification6 `xml:"DataSetId,omitempty"`

	// Type of POI components to send in a status report.
	ComponentType []*DataSetCategory9Code `xml:"CmpntTp,omitempty"`

	// Identification of the delegation scope assigned by the MTM.
	DelegationScopeIdentification *Max35Text `xml:"DlgtnScpId,omitempty"`

	// Definition of the delegation scope, for instance inside the payment application parameters the range of application profiles, the RID (Registered application provider Identification).
	DelegationScopeDefinition *Max3000Binary `xml:"DlgtnScpDef,omitempty"`

	// Proof of delegation to be verified by the POI, when performing the delegated actions.
	DelegationProof *Max5000Binary `xml:"DlgtnProof,omitempty"`

	// Protected proof of delegation.
	ProtectedDelegationProof *ContentInformationType12 `xml:"PrtctdDlgtnProof,omitempty"`

	// Event on which the action has to be activated by the point of interaction (POI).
	Trigger *TerminalManagementActionTrigger1Code `xml:"Trggr"`

	// Additional process to perform before starting or after completing the action by the point of interaction (POI).
	AdditionalProcess []*TerminalManagementAdditionalProcess1Code `xml:"AddtlPrc,omitempty"`

	// Definition of retry process if activation of the action fails.
	ReTry *ProcessRetry2 `xml:"ReTry,omitempty"`

	// Date and time the action has to be performed.
	TimeCondition *ProcessTiming3 `xml:"TmCond,omitempty"`

	// Terminal manager challenge for cryptographic key injection.
	TMChallenge *Max140Binary `xml:"TMChllng,omitempty"`

	// Certificate chain for the encryption of temporary transport key of the key to inject.
	KeyEnciphermentCertificate []*Max10KBinary `xml:"KeyNcphrmntCert,omitempty"`

	// Action to perform in case of error on the related action in progress.
	ErrorAction []*ErrorAction2 `xml:"ErrActn,omitempty"`

	// Additional information about the maintenance action.
	AdditionalInformation []*Max3000Binary `xml:"AddtlInf,omitempty"`
}

func (t *TMSAction5) SetType(value string) {
	t.Type = (*TerminalManagementAction2Code)(&value)
}

func (t *TMSAction5) AddRemoteAccess() *NetworkParameters5 {
	t.RemoteAccess = new(NetworkParameters5)
	return t.RemoteAccess
}

func (t *TMSAction5) AddTerminalManagerIdentification() *GenericIdentification71 {
	t.TerminalManagerIdentification = new(GenericIdentification71)
	return t.TerminalManagerIdentification
}

func (t *TMSAction5) SetTMSProtocol(value string) {
	t.TMSProtocol = (*Max35Text)(&value)
}

func (t *TMSAction5) SetTMSProtocolVersion(value string) {
	t.TMSProtocolVersion = (*Max35Text)(&value)
}

func (t *TMSAction5) AddDataSetIdentification() *DataSetIdentification6 {
	t.DataSetIdentification = new(DataSetIdentification6)
	return t.DataSetIdentification
}

func (t *TMSAction5) AddComponentType(value string) {
	t.ComponentType = append(t.ComponentType, (*DataSetCategory9Code)(&value))
}

func (t *TMSAction5) SetDelegationScopeIdentification(value string) {
	t.DelegationScopeIdentification = (*Max35Text)(&value)
}

func (t *TMSAction5) SetDelegationScopeDefinition(value string) {
	t.DelegationScopeDefinition = (*Max3000Binary)(&value)
}

func (t *TMSAction5) SetDelegationProof(value string) {
	t.DelegationProof = (*Max5000Binary)(&value)
}

func (t *TMSAction5) AddProtectedDelegationProof() *ContentInformationType12 {
	t.ProtectedDelegationProof = new(ContentInformationType12)
	return t.ProtectedDelegationProof
}

func (t *TMSAction5) SetTrigger(value string) {
	t.Trigger = (*TerminalManagementActionTrigger1Code)(&value)
}

func (t *TMSAction5) AddAdditionalProcess(value string) {
	t.AdditionalProcess = append(t.AdditionalProcess, (*TerminalManagementAdditionalProcess1Code)(&value))
}

func (t *TMSAction5) AddReTry() *ProcessRetry2 {
	t.ReTry = new(ProcessRetry2)
	return t.ReTry
}

func (t *TMSAction5) AddTimeCondition() *ProcessTiming3 {
	t.TimeCondition = new(ProcessTiming3)
	return t.TimeCondition
}

func (t *TMSAction5) SetTMChallenge(value string) {
	t.TMChallenge = (*Max140Binary)(&value)
}

func (t *TMSAction5) AddKeyEnciphermentCertificate(value string) {
	t.KeyEnciphermentCertificate = append(t.KeyEnciphermentCertificate, (*Max10KBinary)(&value))
}

func (t *TMSAction5) AddErrorAction() *ErrorAction2 {
	newValue := new(ErrorAction2)
	t.ErrorAction = append(t.ErrorAction, newValue)
	return newValue
}

func (t *TMSAction5) AddAdditionalInformation(value string) {
	t.AdditionalInformation = append(t.AdditionalInformation, (*Max3000Binary)(&value))
}
