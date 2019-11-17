package model

// Provides additional narrative information about the corporate event.
type CorporateEventNarrative2 struct {

	// Issuer’s disclaimer notice relative to the meeting announcement information provided. It may be ignored for automated processing.
	Disclaimer []*Max350Text `xml:"Dsclmr,omitempty"`
}

func (c *CorporateEventNarrative2) AddDisclaimer(value string) {
	c.Disclaimer = append(c.Disclaimer, (*Max350Text)(&value))
}
