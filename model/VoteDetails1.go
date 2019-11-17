package model

// Specifies detailed voting instructions.
type VoteDetails1 struct {

	// Indicates the vote instruction for the resolutions which are announced via the meeting agenda in advance of the meeting.
	VoteInstructionForAgendaResolution *Vote1Choice `xml:"VoteInstrForAgndRsltn"`

	// Indicates the vote instruction for the resolutions that will be added during the meeting.
	VoteInstructionForMeetingResolution *VoteInstructionForMeetingResolution1 `xml:"VoteInstrForMtgRsltn,omitempty"`
}

func (v *VoteDetails1) AddVoteInstructionForAgendaResolution() *Vote1Choice {
	v.VoteInstructionForAgendaResolution = new(Vote1Choice)
	return v.VoteInstructionForAgendaResolution
}

func (v *VoteDetails1) AddVoteInstructionForMeetingResolution() *VoteInstructionForMeetingResolution1 {
	v.VoteInstructionForMeetingResolution = new(VoteInstructionForMeetingResolution1)
	return v.VoteInstructionForMeetingResolution
}
