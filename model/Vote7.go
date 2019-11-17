package model

// Indicates the number of voting rights cast to a resolution.
type Vote7 struct {

	// Numbering of the resolution as specified by the issuer or its agent.
	IssuerLabel *Max35Text `xml:"IssrLabl"`

	// Specifies the acceptance status of a resolution.
	ResolutionStatus *ResolutionStatus2Code `xml:"RsltnSts"`

	// Number of votes in favour of one resolution.
	For *Number `xml:"For,omitempty"`

	// Number of votes against one resolution.
	Against *Number `xml:"Agnst,omitempty"`

	// Number of votes expressed as abstain.
	Abstain *Number `xml:"Abstn,omitempty"`

	// Total votes withheld, for example, in the case where a shareholder wishes not to endorse the election of a board member.
	Withhold *Number `xml:"Wthhld,omitempty"`

	// Number of votes in favour for one year for "say on pay" type of resolution.
	OneYear *Number `xml:"OneYr,omitempty"`

	// Number of votes in favour of two years for "say on pay" type of resolution.
	TwoYears *Number `xml:"TwoYrs,omitempty"`

	// Number of votes in favour of three years for "say on pay" type of resolution.
	ThreeYears *Number `xml:"ThreeYrs,omitempty"`

	// Number of votes for which no action has been taken.
	NoAction *Number `xml:"NoActn,omitempty"`
}

func (v *Vote7) SetIssuerLabel(value string) {
	v.IssuerLabel = (*Max35Text)(&value)
}

func (v *Vote7) SetResolutionStatus(value string) {
	v.ResolutionStatus = (*ResolutionStatus2Code)(&value)
}

func (v *Vote7) SetFor(value string) {
	v.For = (*Number)(&value)
}

func (v *Vote7) SetAgainst(value string) {
	v.Against = (*Number)(&value)
}

func (v *Vote7) SetAbstain(value string) {
	v.Abstain = (*Number)(&value)
}

func (v *Vote7) SetWithhold(value string) {
	v.Withhold = (*Number)(&value)
}

func (v *Vote7) SetOneYear(value string) {
	v.OneYear = (*Number)(&value)
}

func (v *Vote7) SetTwoYears(value string) {
	v.TwoYears = (*Number)(&value)
}

func (v *Vote7) SetThreeYears(value string) {
	v.ThreeYears = (*Number)(&value)
}

func (v *Vote7) SetNoAction(value string) {
	v.NoAction = (*Number)(&value)
}
