package model

// Decision of the voting party for one resolution. Several types of decisions can be indicated to allow for split vote specification.
type Vote6 struct {

	// Numbering of the resolution as specified by the issuer or  its agent.
	IssuerLabel *Max35Text `xml:"IssrLabl"`

	// Number of votes in favour of one resolution.
	For *Number `xml:"For,omitempty"`

	// Number of votes against one resolution.
	Against *Number `xml:"Agnst,omitempty"`

	// Number of votes expressed as abstain for one resolution.
	Abstain *Number `xml:"Abstn,omitempty"`

	// Number of votes withheld for one resolution.
	Withhold *Number `xml:"Wthhld,omitempty"`

	// Number of votes in line with the votes of the management.
	WithManagement *Number `xml:"WthMgmt,omitempty"`

	// Number of votes against the voting recommendation of the management.
	AgainstManagement *Number `xml:"AgnstMgmt,omitempty"`

	// Number of votes for which decision is left to the party that will exercise the voting right.
	Discretionary *Number `xml:"Dscrtnry,omitempty"`

	// Number of votes in favour for one year for "say on pay" type of resolution.
	OneYear *Number `xml:"OneYr,omitempty"`

	// Number of votes in favour of two years for "say on pay" type of resolution.
	TwoYears *Number `xml:"TwoYrs,omitempty"`

	// Number of votes in favour of three years for "say on pay" type of resolution.
	ThreeYears *Number `xml:"ThreeYrs,omitempty"`

	// Number of votes for which no action has been taken.
	NoAction *Number `xml:"NoActn,omitempty"`

	// Resolution withdrawn at the meeting.
	Withdrawn *YesNoIndicator `xml:"Wdrwn,omitempty"`
}

func (v *Vote6) SetIssuerLabel(value string) {
	v.IssuerLabel = (*Max35Text)(&value)
}

func (v *Vote6) SetFor(value string) {
	v.For = (*Number)(&value)
}

func (v *Vote6) SetAgainst(value string) {
	v.Against = (*Number)(&value)
}

func (v *Vote6) SetAbstain(value string) {
	v.Abstain = (*Number)(&value)
}

func (v *Vote6) SetWithhold(value string) {
	v.Withhold = (*Number)(&value)
}

func (v *Vote6) SetWithManagement(value string) {
	v.WithManagement = (*Number)(&value)
}

func (v *Vote6) SetAgainstManagement(value string) {
	v.AgainstManagement = (*Number)(&value)
}

func (v *Vote6) SetDiscretionary(value string) {
	v.Discretionary = (*Number)(&value)
}

func (v *Vote6) SetOneYear(value string) {
	v.OneYear = (*Number)(&value)
}

func (v *Vote6) SetTwoYears(value string) {
	v.TwoYears = (*Number)(&value)
}

func (v *Vote6) SetThreeYears(value string) {
	v.ThreeYears = (*Number)(&value)
}

func (v *Vote6) SetNoAction(value string) {
	v.NoAction = (*Number)(&value)
}

func (v *Vote6) SetWithdrawn(value string) {
	v.Withdrawn = (*YesNoIndicator)(&value)
}
