package model

// Provides information on the instruction.
type Instruction3 struct {

	// Identifies the detailed instruction.
	InstructionIdentification *Max35Text `xml:"InstrId"`

	// Date at which the instruction must be executed.
	RequestedExecutionDate *ISODateTime `xml:"ReqdExctnDt,omitempty"`

	// Indicates that a vote execution confirmation is requested.
	VoteExecutionConfirmation *YesNoIndicator `xml:"VoteExctnConf"`

	// Identification of the securities account.
	AccountDetails *SafekeepingAccount6 `xml:"AcctDtls"`

	// Identification of the person appointed by the security holder as proxy.
	Proxy *Proxy6 `xml:"Prxy,omitempty"`

	// Specifies detailed voting instructions.
	VoteDetails *VoteDetails3 `xml:"VoteDtls,omitempty"`

	// Identification of the security holder who will attend and vote at the meeting in person and/or the person assigned by the security holder to attend the meeting without having any voting rights or taking any action.
	MeetingAttendee []*IndividualPerson26 `xml:"MtgAttndee,omitempty"`

	// Request to execute specific instructions, such as participation registration, securities registration or blocking of securities.
	SpecificInstructionRequest *SpecificInstructionRequest1 `xml:"SpcfcInstrReq,omitempty"`
}

func (i *Instruction3) SetInstructionIdentification(value string) {
	i.InstructionIdentification = (*Max35Text)(&value)
}

func (i *Instruction3) SetRequestedExecutionDate(value string) {
	i.RequestedExecutionDate = (*ISODateTime)(&value)
}

func (i *Instruction3) SetVoteExecutionConfirmation(value string) {
	i.VoteExecutionConfirmation = (*YesNoIndicator)(&value)
}

func (i *Instruction3) AddAccountDetails() *SafekeepingAccount6 {
	i.AccountDetails = new(SafekeepingAccount6)
	return i.AccountDetails
}

func (i *Instruction3) AddProxy() *Proxy6 {
	i.Proxy = new(Proxy6)
	return i.Proxy
}

func (i *Instruction3) AddVoteDetails() *VoteDetails3 {
	i.VoteDetails = new(VoteDetails3)
	return i.VoteDetails
}

func (i *Instruction3) AddMeetingAttendee() *IndividualPerson26 {
	newValue := new(IndividualPerson26)
	i.MeetingAttendee = append(i.MeetingAttendee, newValue)
	return newValue
}

func (i *Instruction3) AddSpecificInstructionRequest() *SpecificInstructionRequest1 {
	i.SpecificInstructionRequest = new(SpecificInstructionRequest1)
	return i.SpecificInstructionRequest
}
