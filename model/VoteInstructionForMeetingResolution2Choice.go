package model

// Vote instruction applying to resolutions added during the meeting.
type VoteInstructionForMeetingResolution2Choice struct {

	// Specifies the vote recommendation for resolutions added during the meeting.
	VoteIndication *VoteInstruction4Code `xml:"VoteIndctn"`

	// Specifies the name and address of the shareholder to whom the rights to vote are transferred for resolutions added during the meeting.
	Shareholder *NameAndAddress9 `xml:"Shrhldr"`
}

func (v *VoteInstructionForMeetingResolution2Choice) SetVoteIndication(value string) {
	v.VoteIndication = (*VoteInstruction4Code)(&value)
}

func (v *VoteInstructionForMeetingResolution2Choice) AddShareholder() *NameAndAddress9 {
	v.Shareholder = new(NameAndAddress9)
	return v.Shareholder
}
