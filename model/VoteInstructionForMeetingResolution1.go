package model

// Vote instruction applying to resolutions added during the meeting.
type VoteInstructionForMeetingResolution1 struct {

	// Specifies the vote recommendation for resolutions added during the meeting.
	VoteIndication *VoteInstructionAtMeeting1Code `xml:"VoteIndctn"`

	// Specifies the name and address of the shareholder to whom the rights to vote are transferred for resolutions added during the meeting.
	Shareholder *NameAndAddress9 `xml:"Shrhldr"`
}

func (v *VoteInstructionForMeetingResolution1) SetVoteIndication(value string) {
	v.VoteIndication = (*VoteInstructionAtMeeting1Code)(&value)
}

func (v *VoteInstructionForMeetingResolution1) AddShareholder() *NameAndAddress9 {
	v.Shareholder = new(NameAndAddress9)
	return v.Shareholder
}
