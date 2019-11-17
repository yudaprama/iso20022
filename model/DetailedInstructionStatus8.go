package model

// Status applying to individual instructions of a MeetingInstruction.
type DetailedInstructionStatus8 struct {

	// Identifies the detailed instruction within an instruction message.
	InstructionIdentification *Max35Text `xml:"InstrId"`

	// Identifies the safekeeping account.
	AccountIdentification *Max35Text `xml:"AcctId,omitempty"`

	// Identifies the subaccount of the safekeeping account.
	SubAccountIdentification *Max35Text `xml:"SubAcctId,omitempty"`

	// Status applying to individual instructions of a MeetingInstruction.
	InstructionStatus *InstructionStatus4Choice `xml:"InstrSts"`
}

func (d *DetailedInstructionStatus8) SetInstructionIdentification(value string) {
	d.InstructionIdentification = (*Max35Text)(&value)
}

func (d *DetailedInstructionStatus8) SetAccountIdentification(value string) {
	d.AccountIdentification = (*Max35Text)(&value)
}

func (d *DetailedInstructionStatus8) SetSubAccountIdentification(value string) {
	d.SubAccountIdentification = (*Max35Text)(&value)
}

func (d *DetailedInstructionStatus8) AddInstructionStatus() *InstructionStatus4Choice {
	d.InstructionStatus = new(InstructionStatus4Choice)
	return d.InstructionStatus
}
