package model

// Status applying to individual instructions of a MeetingInstruction.
type DetailedInstructionStatus11 struct {

	// Identifies the detailed instruction within an instruction message.
	InstructionIdentification *Max35Text `xml:"InstrId"`

	// Identifies the safekeeping account.
	AccountIdentification *Max35Text `xml:"AcctId,omitempty"`

	// Identifies the subaccount of the safekeeping account.
	SubAccountIdentification *Max35Text `xml:"SubAcctId,omitempty"`

	// Status applying to individual instructions of a MeetingInstruction.
	InstructionStatus *InstructionStatus6Choice `xml:"InstrSts"`
}

func (d *DetailedInstructionStatus11) SetInstructionIdentification(value string) {
	d.InstructionIdentification = (*Max35Text)(&value)
}

func (d *DetailedInstructionStatus11) SetAccountIdentification(value string) {
	d.AccountIdentification = (*Max35Text)(&value)
}

func (d *DetailedInstructionStatus11) SetSubAccountIdentification(value string) {
	d.SubAccountIdentification = (*Max35Text)(&value)
}

func (d *DetailedInstructionStatus11) AddInstructionStatus() *InstructionStatus6Choice {
	d.InstructionStatus = new(InstructionStatus6Choice)
	return d.InstructionStatus
}
