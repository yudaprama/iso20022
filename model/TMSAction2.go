package model

// Single terminal management action to be performed by the point of interaction.
type TMSAction2 struct {

	// Types of action to be performed by a point of interaction (POI).
	Type *TerminalManagementAction1Code `xml:"Tp"`

	// Communication parameters of the terminal management system to contact.
	Address *NetworkParameters1 `xml:"Adr,omitempty"`

	// Data set on which the action has to be performed.
	DataSetIdentification *DataSetIdentification3 `xml:"DataSetId,omitempty"`

	// Event on which the action has to be activated by the point of interaction (POI).
	Trigger *TerminalManagementActionTrigger1Code `xml:"Trggr"`

	// Additional process to perform before starting or after completing the action by the point of interaction (POI).
	AdditionalProcess *TerminalManagementAdditionalProcess1Code `xml:"AddtlPrc,omitempty"`

	// Date and time the action has to be performed.
	TimeCondition *ProcessTiming2 `xml:"TmCond,omitempty"`

	// Terminal manager challenge for cryptographic key injection.
	TMChallenge *Max140Binary `xml:"TMChllng,omitempty"`

	// Certificate chain for the encryption of temporary transport key of the key to inject.
	KeyEnciphermentCertificate []*Max5000Binary `xml:"KeyNcphrmntCert,omitempty"`

	// Action to perform in case of error on the related action in progress.
	ErrorAction []*ErrorAction2 `xml:"ErrActn,omitempty"`
}

func (t *TMSAction2) SetType(value string) {
	t.Type = (*TerminalManagementAction1Code)(&value)
}

func (t *TMSAction2) AddAddress() *NetworkParameters1 {
	t.Address = new(NetworkParameters1)
	return t.Address
}

func (t *TMSAction2) AddDataSetIdentification() *DataSetIdentification3 {
	t.DataSetIdentification = new(DataSetIdentification3)
	return t.DataSetIdentification
}

func (t *TMSAction2) SetTrigger(value string) {
	t.Trigger = (*TerminalManagementActionTrigger1Code)(&value)
}

func (t *TMSAction2) SetAdditionalProcess(value string) {
	t.AdditionalProcess = (*TerminalManagementAdditionalProcess1Code)(&value)
}

func (t *TMSAction2) AddTimeCondition() *ProcessTiming2 {
	t.TimeCondition = new(ProcessTiming2)
	return t.TimeCondition
}

func (t *TMSAction2) SetTMChallenge(value string) {
	t.TMChallenge = (*Max140Binary)(&value)
}

func (t *TMSAction2) AddKeyEnciphermentCertificate(value string) {
	t.KeyEnciphermentCertificate = append(t.KeyEnciphermentCertificate, (*Max5000Binary)(&value))
}

func (t *TMSAction2) AddErrorAction() *ErrorAction2 {
	newValue := new(ErrorAction2)
	t.ErrorAction = append(t.ErrorAction, newValue)
	return newValue
}
