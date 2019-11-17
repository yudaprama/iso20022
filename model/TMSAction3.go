package model

// Single terminal management action to be performed by the point of interaction.
type TMSAction3 struct {

	// Types of action to be performed by a point of interaction (POI).
	Type *TerminalManagementAction1Code `xml:"Tp"`

	// Communication parameters of the terminal management system to contact.
	Address *NetworkParameters1 `xml:"Adr,omitempty"`

	// Data set on which the action has to be performed.
	DataSetIdentification *DataSetIdentification3 `xml:"DataSetId,omitempty"`

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
}

func (t *TMSAction3) SetType(value string) {
	t.Type = (*TerminalManagementAction1Code)(&value)
}

func (t *TMSAction3) AddAddress() *NetworkParameters1 {
	t.Address = new(NetworkParameters1)
	return t.Address
}

func (t *TMSAction3) AddDataSetIdentification() *DataSetIdentification3 {
	t.DataSetIdentification = new(DataSetIdentification3)
	return t.DataSetIdentification
}

func (t *TMSAction3) SetTrigger(value string) {
	t.Trigger = (*TerminalManagementActionTrigger1Code)(&value)
}

func (t *TMSAction3) AddAdditionalProcess(value string) {
	t.AdditionalProcess = append(t.AdditionalProcess, (*TerminalManagementAdditionalProcess1Code)(&value))
}

func (t *TMSAction3) AddReTry() *ProcessRetry2 {
	t.ReTry = new(ProcessRetry2)
	return t.ReTry
}

func (t *TMSAction3) AddTimeCondition() *ProcessTiming3 {
	t.TimeCondition = new(ProcessTiming3)
	return t.TimeCondition
}

func (t *TMSAction3) SetTMChallenge(value string) {
	t.TMChallenge = (*Max140Binary)(&value)
}

func (t *TMSAction3) AddKeyEnciphermentCertificate(value string) {
	t.KeyEnciphermentCertificate = append(t.KeyEnciphermentCertificate, (*Max10KBinary)(&value))
}

func (t *TMSAction3) AddErrorAction() *ErrorAction2 {
	newValue := new(ErrorAction2)
	t.ErrorAction = append(t.ErrorAction, newValue)
	return newValue
}
