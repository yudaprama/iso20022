package model

// Single terminal management action to be performed by the point of interaction.
type TMSAction1 struct {

	// Types of action to be performed by a point of interaction (POI).
	Type *TerminalManagementAction1Code `xml:"Tp"`

	// Communication parameters of the terminal management system to contact.
	Address *NetworkParameters1 `xml:"Adr,omitempty"`

	// Data set on which the action has to be performed.
	DataSetIdentification *DataSetIdentification2 `xml:"DataSetId,omitempty"`

	// Event on which the action has to be activated by the point of interaction (POI).
	Trigger *TerminalManagementActionTrigger1Code `xml:"Trggr"`

	// Additional process to perform before starting or after completing the action by the point of interaction (POI).
	AdditionalProcess *TerminalManagementAdditionalProcess1Code `xml:"AddtlPrc,omitempty"`

	// Date and time the action has to be performed.
	TimeCondition *ProcessTiming1 `xml:"TmCond,omitempty"`

	// Action to perform in case of error on the related action in progress.
	ErrorAction []*ErrorAction1 `xml:"ErrActn,omitempty"`
}

func (t *TMSAction1) SetType(value string) {
	t.Type = (*TerminalManagementAction1Code)(&value)
}

func (t *TMSAction1) AddAddress() *NetworkParameters1 {
	t.Address = new(NetworkParameters1)
	return t.Address
}

func (t *TMSAction1) AddDataSetIdentification() *DataSetIdentification2 {
	t.DataSetIdentification = new(DataSetIdentification2)
	return t.DataSetIdentification
}

func (t *TMSAction1) SetTrigger(value string) {
	t.Trigger = (*TerminalManagementActionTrigger1Code)(&value)
}

func (t *TMSAction1) SetAdditionalProcess(value string) {
	t.AdditionalProcess = (*TerminalManagementAdditionalProcess1Code)(&value)
}

func (t *TMSAction1) AddTimeCondition() *ProcessTiming1 {
	t.TimeCondition = new(ProcessTiming1)
	return t.TimeCondition
}

func (t *TMSAction1) AddErrorAction() *ErrorAction1 {
	newValue := new(ErrorAction1)
	t.ErrorAction = append(t.ErrorAction, newValue)
	return newValue
}
