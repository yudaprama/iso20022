package model

// Result of an individual terminal management action performed by the point of interaction.
type TMSEvent3 struct {

	// Date time of the terminal management action performed by the point of interaction.
	TimeStamp *ISODateTime `xml:"TmStmp"`

	// Final result of the processed terminal management action.
	Result *TerminalManagementActionResult1Code `xml:"Rslt"`

	// Identification of the terminal management action performed by the point of interaction.
	ActionIdentification *TMSActionIdentification3 `xml:"ActnId"`

	// Additional information related to a failure.
	AdditionalErrorInformation *Max70Text `xml:"AddtlErrInf,omitempty"`
}

func (t *TMSEvent3) SetTimeStamp(value string) {
	t.TimeStamp = (*ISODateTime)(&value)
}

func (t *TMSEvent3) SetResult(value string) {
	t.Result = (*TerminalManagementActionResult1Code)(&value)
}

func (t *TMSEvent3) AddActionIdentification() *TMSActionIdentification3 {
	t.ActionIdentification = new(TMSActionIdentification3)
	return t.ActionIdentification
}

func (t *TMSEvent3) SetAdditionalErrorInformation(value string) {
	t.AdditionalErrorInformation = (*Max70Text)(&value)
}
