package model

// Provides information on the instruction.
type Instruction2 struct {

	// Identifies the detailed instruction.
	InstructionIdentification *Max35Text `xml:"InstrId"`

	// Date at which the instruction must be executed.
	RequestedExecutionDate *ISODateTime `xml:"ReqdExctnDt,omitempty"`

	// Indicates that a Vote execution confirmation is requested.
	VoteExecutionConfirmation *YesNoIndicator `xml:"VoteExctnConf"`

	// Identification of the securities account.
	AccountDetails *SafekeepingAccount4 `xml:"AcctDtls"`

	// Identification of the person appointed by the security holder as proxy.
	Proxy *Proxy4 `xml:"Prxy,omitempty"`

	// Specifies detailed voting instructions.
	VoteDetails *VoteDetails2 `xml:"VoteDtls,omitempty"`

	// Identification of the security holder who will attend and vote at the meeting in person and/or a person assigned by the security holder to attend the meeting without having any voting rights or taking any action.
	MeetingAttendee []*IndividualPerson17 `xml:"MtgAttndee,omitempty"`

	// Request to execute specific instructions, such as participation registration, securities registration or blocking of securities.
	SpecificInstructionRequest *SpecificInstructionRequest1 `xml:"SpcfcInstrReq,omitempty"`
}

func (i *Instruction2) SetInstructionIdentification(value string) {
	i.InstructionIdentification = (*Max35Text)(&value)
}

func (i *Instruction2) SetRequestedExecutionDate(value string) {
	i.RequestedExecutionDate = (*ISODateTime)(&value)
}

func (i *Instruction2) SetVoteExecutionConfirmation(value string) {
	i.VoteExecutionConfirmation = (*YesNoIndicator)(&value)
}

func (i *Instruction2) AddAccountDetails() *SafekeepingAccount4 {
	i.AccountDetails = new(SafekeepingAccount4)
	return i.AccountDetails
}

func (i *Instruction2) AddProxy() *Proxy4 {
	i.Proxy = new(Proxy4)
	return i.Proxy
}

func (i *Instruction2) AddVoteDetails() *VoteDetails2 {
	i.VoteDetails = new(VoteDetails2)
	return i.VoteDetails
}

func (i *Instruction2) AddMeetingAttendee() *IndividualPerson17 {
	newValue := new(IndividualPerson17)
	i.MeetingAttendee = append(i.MeetingAttendee, newValue)
	return newValue
}

func (i *Instruction2) AddSpecificInstructionRequest() *SpecificInstructionRequest1 {
	i.SpecificInstructionRequest = new(SpecificInstructionRequest1)
	return i.SpecificInstructionRequest
}
