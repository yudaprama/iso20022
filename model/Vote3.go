package model

// Decision of the voting party for one resolution.
type Vote3 struct {

	// Numbering of the resolution as specified by the issuer or  its agent.
	IssuerLabel *Max35Text `xml:"IssrLabl"`

	// Specifies the different instructions which can be used to vote.
	VoteOption *VoteInstruction2Code `xml:"VoteOptn"`
}

func (v *Vote3) SetIssuerLabel(value string) {
	v.IssuerLabel = (*Max35Text)(&value)
}

func (v *Vote3) SetVoteOption(value string) {
	v.VoteOption = (*VoteInstruction2Code)(&value)
}
