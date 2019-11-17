package model

// Decision of the voting party for one resolution. Several types of decisions can be indicated to allow for split vote specification.
type Vote4 struct {

	// Numbering of the resolution as specified by the issuer or  its agent.
	IssuerLabel *Max35Text `xml:"IssrLabl"`

	// Number of votes in favour of one resolution
	For *Number `xml:"For,omitempty"`

	// Number of votes against one resolution
	Against *Number `xml:"Agnst,omitempty"`

	// Number of votes expressed as abstain for one resolution.
	Abstain *Number `xml:"Abstn,omitempty"`

	// Number of votes withheld for one resolution
	Withhold *Number `xml:"Wthhld,omitempty"`

	// Number of votes in line with the votes of the management.
	WithManagement *Number `xml:"WthMgmt,omitempty"`

	// Number of votes against the voting recommendation of the management.
	AgainstManagement *Number `xml:"AgnstMgmt,omitempty"`

	// Number of votes for which decision is left to the party that will exercise the voting right.
	Discretionary *Number `xml:"Dscrtnry,omitempty"`

	// Number of votes for which no action has been taken.
	NoAction *Number `xml:"NoActn,omitempty"`
}

func (v *Vote4) SetIssuerLabel(value string) {
	v.IssuerLabel = (*Max35Text)(&value)
}

func (v *Vote4) SetFor(value string) {
	v.For = (*Number)(&value)
}

func (v *Vote4) SetAgainst(value string) {
	v.Against = (*Number)(&value)
}

func (v *Vote4) SetAbstain(value string) {
	v.Abstain = (*Number)(&value)
}

func (v *Vote4) SetWithhold(value string) {
	v.Withhold = (*Number)(&value)
}

func (v *Vote4) SetWithManagement(value string) {
	v.WithManagement = (*Number)(&value)
}

func (v *Vote4) SetAgainstManagement(value string) {
	v.AgainstManagement = (*Number)(&value)
}

func (v *Vote4) SetDiscretionary(value string) {
	v.Discretionary = (*Number)(&value)
}

func (v *Vote4) SetNoAction(value string) {
	v.NoAction = (*Number)(&value)
}
