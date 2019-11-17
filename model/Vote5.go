package model

// Indicates the number of voting rights cast to a resolution.
type Vote5 struct {

	// Numbering of the resolution as specified by the issuer or its agent.
	IssuerLabel *Max35Text `xml:"IssrLabl"`

	// Specifies whether a resolution is accepted or not.
	Accepted *YesNoIndicator `xml:"Accptd"`

	// Number of votes in favour of one resolution
	For *Number `xml:"For,omitempty"`

	// Number of votes against one resolution
	Against *Number `xml:"Agnst,omitempty"`

	// Number of votes expressed as abstain.
	Abstain *Number `xml:"Abstn,omitempty"`

	// Total votes withheld, eg in the case where a shareholder wishes not to endorse the election of a board member.
	Withhold *Number `xml:"Wthhld,omitempty"`

	// Number of votes for which no action has been taken.
	NoAction *Number `xml:"NoActn,omitempty"`
}

func (v *Vote5) SetIssuerLabel(value string) {
	v.IssuerLabel = (*Max35Text)(&value)
}

func (v *Vote5) SetAccepted(value string) {
	v.Accepted = (*YesNoIndicator)(&value)
}

func (v *Vote5) SetFor(value string) {
	v.For = (*Number)(&value)
}

func (v *Vote5) SetAgainst(value string) {
	v.Against = (*Number)(&value)
}

func (v *Vote5) SetAbstain(value string) {
	v.Abstain = (*Number)(&value)
}

func (v *Vote5) SetWithhold(value string) {
	v.Withhold = (*Number)(&value)
}

func (v *Vote5) SetNoAction(value string) {
	v.NoAction = (*Number)(&value)
}
