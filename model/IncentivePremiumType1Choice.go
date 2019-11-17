package model

// Choice of cash premium paid to the security holder when voting.
type IncentivePremiumType1Choice struct {

	// Number of securities giving right to a premium.
	PerSecurity *Number `xml:"PerScty"`

	// Number of votes giving right to a premium.
	PerVote *Number `xml:"PerVote"`

	// Indicates that the premium is given per attendee.
	PerAttendee *YesNoIndicator `xml:"PerAttndee"`
}

func (i *IncentivePremiumType1Choice) SetPerSecurity(value string) {
	i.PerSecurity = (*Number)(&value)
}

func (i *IncentivePremiumType1Choice) SetPerVote(value string) {
	i.PerVote = (*Number)(&value)
}

func (i *IncentivePremiumType1Choice) SetPerAttendee(value string) {
	i.PerAttendee = (*YesNoIndicator)(&value)
}
