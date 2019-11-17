package model

// Specifies detailed voting instructions.
type VoteDetails2 struct {

	// Indicates the vote instruction for the resolutions which are announced via the meeting agenda in advance of the meeting.
	VoteInstructionForAgendaResolution *Vote2Choice `xml:"VoteInstrForAgndRsltn"`

	// Indicates the vote instruction for the resolutions that may arise at the meeting but were not previously provided in the agenda.
	VoteInstructionForMeetingResolution *VoteInstructionForMeetingResolution1Choice `xml:"VoteInstrForMtgRsltn,omitempty"`
}

func (v *VoteDetails2) AddVoteInstructionForAgendaResolution() *Vote2Choice {
	v.VoteInstructionForAgendaResolution = new(Vote2Choice)
	return v.VoteInstructionForAgendaResolution
}

func (v *VoteDetails2) AddVoteInstructionForMeetingResolution() *VoteInstructionForMeetingResolution1Choice {
	v.VoteInstructionForMeetingResolution = new(VoteInstructionForMeetingResolution1Choice)
	return v.VoteInstructionForMeetingResolution
}
