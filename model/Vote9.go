package model

// Decision of the voting party for one resolution.
type Vote9 struct {

	// Numbering of the resolution as specified by the issuer or  its agent.
	IssuerLabel *Max35Text `xml:"IssrLabl"`

	// Specifies the different instructions that can be used to vote.
	VoteOption *VoteInstruction3Code `xml:"VoteOptn"`
}

func (v *Vote9) SetIssuerLabel(value string) {
	v.IssuerLabel = (*Max35Text)(&value)
}

func (v *Vote9) SetVoteOption(value string) {
	v.VoteOption = (*VoteInstruction3Code)(&value)
}
