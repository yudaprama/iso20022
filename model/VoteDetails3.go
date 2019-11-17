package model

// Specifies detailed voting instructions.
type VoteDetails3 struct {

	// Indicates the vote instruction for the resolutions that are announced via the meeting agenda in advance of the meeting.
	VoteInstructionForAgendaResolution *Vote3Choice `xml:"VoteInstrForAgndRsltn"`

	// Indicates the vote instruction for the resolutions that may arise at the meeting but were not previously provided in the agenda.
	VoteInstructionForMeetingResolution *VoteInstructionForMeetingResolution2Choice `xml:"VoteInstrForMtgRsltn,omitempty"`
}

func (v *VoteDetails3) AddVoteInstructionForAgendaResolution() *Vote3Choice {
	v.VoteInstructionForAgendaResolution = new(Vote3Choice)
	return v.VoteInstructionForAgendaResolution
}

func (v *VoteDetails3) AddVoteInstructionForMeetingResolution() *VoteInstructionForMeetingResolution2Choice {
	v.VoteInstructionForMeetingResolution = new(VoteInstructionForMeetingResolution2Choice)
	return v.VoteInstructionForMeetingResolution
}
