package model

// Status applying to individual instructions of a MeetingInstruction.
type DetailedInstructionStatus10 struct {

	// Identifies the detailed instruction with an instruction message.
	InstructionIdentification *Max35Text `xml:"InstrId"`

	// Identifies the safekeeping account.
	AccountIdentification *Max35Text `xml:"AcctId,omitempty"`

	// Party that legally owns the account.
	AccountOwner *PartyIdentification40Choice `xml:"AcctOwnr,omitempty"`

	// Identifies the subaccount of the safekeeping account.
	SubAccountIdentification *Max35Text `xml:"SubAcctId,omitempty"`

	// Owner of the voting rights.
	RightsHolder []*PartyIdentification40Choice `xml:"RghtsHldr,omitempty"`

	// Indicates whether standing instructions have been applied or not.
	StandingInstruction *YesNoIndicator `xml:"StgInstr"`

	// Details of the vote.
	VotePerResolution []*Vote6 `xml:"VotePerRsltn,omitempty"`
}

func (d *DetailedInstructionStatus10) SetInstructionIdentification(value string) {
	d.InstructionIdentification = (*Max35Text)(&value)
}

func (d *DetailedInstructionStatus10) SetAccountIdentification(value string) {
	d.AccountIdentification = (*Max35Text)(&value)
}

func (d *DetailedInstructionStatus10) AddAccountOwner() *PartyIdentification40Choice {
	d.AccountOwner = new(PartyIdentification40Choice)
	return d.AccountOwner
}

func (d *DetailedInstructionStatus10) SetSubAccountIdentification(value string) {
	d.SubAccountIdentification = (*Max35Text)(&value)
}

func (d *DetailedInstructionStatus10) AddRightsHolder() *PartyIdentification40Choice {
	newValue := new(PartyIdentification40Choice)
	d.RightsHolder = append(d.RightsHolder, newValue)
	return newValue
}

func (d *DetailedInstructionStatus10) SetStandingInstruction(value string) {
	d.StandingInstruction = (*YesNoIndicator)(&value)
}

func (d *DetailedInstructionStatus10) AddVotePerResolution() *Vote6 {
	newValue := new(Vote6)
	d.VotePerResolution = append(d.VotePerResolution, newValue)
	return newValue
}
