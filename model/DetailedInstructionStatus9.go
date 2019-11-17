package model

// Status applying to individual instructions of a MeetingInstruction.
type DetailedInstructionStatus9 struct {

	// Identifies the detailed instruction with an instruction message.
	InstructionIdentification *Max35Text `xml:"InstrId"`

	// Identifies the safekeeping account.
	AccountIdentification *Max35Text `xml:"AcctId,omitempty"`

	// Party that legally owns the account.
	AccountOwner *PartyIdentification9Choice `xml:"AcctOwnr,omitempty"`

	// Identifies the subaccount of the safekeeping account.
	SubAccountIdentification *Max35Text `xml:"SubAcctId,omitempty"`

	// Owner of the voting rights.
	RightsHolder []*PartyIdentification9Choice `xml:"RghtsHldr,omitempty"`

	// Indicates whether standing instructions have been applied or not.
	StandingInstruction *YesNoIndicator `xml:"StgInstr"`

	// Details of the vote.
	VotePerResolution []*Vote4 `xml:"VotePerRsltn"`
}

func (d *DetailedInstructionStatus9) SetInstructionIdentification(value string) {
	d.InstructionIdentification = (*Max35Text)(&value)
}

func (d *DetailedInstructionStatus9) SetAccountIdentification(value string) {
	d.AccountIdentification = (*Max35Text)(&value)
}

func (d *DetailedInstructionStatus9) AddAccountOwner() *PartyIdentification9Choice {
	d.AccountOwner = new(PartyIdentification9Choice)
	return d.AccountOwner
}

func (d *DetailedInstructionStatus9) SetSubAccountIdentification(value string) {
	d.SubAccountIdentification = (*Max35Text)(&value)
}

func (d *DetailedInstructionStatus9) AddRightsHolder() *PartyIdentification9Choice {
	newValue := new(PartyIdentification9Choice)
	d.RightsHolder = append(d.RightsHolder, newValue)
	return newValue
}

func (d *DetailedInstructionStatus9) SetStandingInstruction(value string) {
	d.StandingInstruction = (*YesNoIndicator)(&value)
}

func (d *DetailedInstructionStatus9) AddVotePerResolution() *Vote4 {
	newValue := new(Vote4)
	d.VotePerResolution = append(d.VotePerResolution, newValue)
	return newValue
}
