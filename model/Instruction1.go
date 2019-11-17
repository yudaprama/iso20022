package model

// Provides information on the instruction.
type Instruction1 struct {

	// Identifies the detailed instruction.
	InstructionIdentification *Max35Text `xml:"InstrId"`

	// Date at which the instruction must be executed.
	RequestedExecutionDate *ISODateTime `xml:"ReqdExctnDt,omitempty"`

	// Indicates that a Vote execution confirmation is requested.
	VoteExecutionConfirmation *YesNoIndicator `xml:"VoteExctnConf"`

	// Identification of the securities account.
	AccountDetails *SafekeepingAccount3 `xml:"AcctDtls"`

	// Identification of the person appointed by the security holder as proxy.
	Proxy *Proxy2 `xml:"Prxy,omitempty"`

	// Specifies detailed voting instructions.
	VoteDetails *VoteDetails1 `xml:"VoteDtls,omitempty"`

	// Identification of the security holder who will attend and vote at the meeting in person and/or a person assigned by the security holder to attend the meeting without having any voting rights or taking any action.
	MeetingAttendee []*IndividualPerson13 `xml:"MtgAttndee,omitempty"`

	// Request to execute specific instructions, such as participation registration, securities registration or blocking of securities.
	SpecificInstructionRequest *SpecificInstructionRequest1 `xml:"SpcfcInstrReq,omitempty"`
}

func (i *Instruction1) SetInstructionIdentification(value string) {
	i.InstructionIdentification = (*Max35Text)(&value)
}

func (i *Instruction1) SetRequestedExecutionDate(value string) {
	i.RequestedExecutionDate = (*ISODateTime)(&value)
}

func (i *Instruction1) SetVoteExecutionConfirmation(value string) {
	i.VoteExecutionConfirmation = (*YesNoIndicator)(&value)
}

func (i *Instruction1) AddAccountDetails() *SafekeepingAccount3 {
	i.AccountDetails = new(SafekeepingAccount3)
	return i.AccountDetails
}

func (i *Instruction1) AddProxy() *Proxy2 {
	i.Proxy = new(Proxy2)
	return i.Proxy
}

func (i *Instruction1) AddVoteDetails() *VoteDetails1 {
	i.VoteDetails = new(VoteDetails1)
	return i.VoteDetails
}

func (i *Instruction1) AddMeetingAttendee() *IndividualPerson13 {
	newValue := new(IndividualPerson13)
	i.MeetingAttendee = append(i.MeetingAttendee, newValue)
	return newValue
}

func (i *Instruction1) AddSpecificInstructionRequest() *SpecificInstructionRequest1 {
	i.SpecificInstructionRequest = new(SpecificInstructionRequest1)
	return i.SpecificInstructionRequest
}
