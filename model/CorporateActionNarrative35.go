package model

// Provides additional information such as the taxation conditions.
type CorporateActionNarrative35 struct {

	// Provides additional information or specifies in more detail the content of a message. This field may only be used when the information to be transmitted, cannot be coded.
	AdditionalText []*RestrictedFINXMax350Text `xml:"AddtlTxt,omitempty"`

	// Provides information that can be ignored for automated processing; - reiteration of information that has been included within structured fields, - or narrative information not needed for automatic processing.
	NarrativeVersion []*RestrictedFINXMax350Text `xml:"NrrtvVrsn,omitempty"`

	// Provides additional information regarding the party, for example, the contact unit or person responsible for the transaction identified in the message.
	PartyContactNarrative []*RestrictedFINXMax350Text `xml:"PtyCtctNrrtv,omitempty"`

	// Provides taxation conditions that cannot be included within the structured fields of this message and has not been mentioned in the Service Level Agreement (SLA)
	TaxationConditions []*RestrictedFINXMax350Text `xml:"TaxtnConds,omitempty"`
}

func (c *CorporateActionNarrative35) AddAdditionalText(value string) {
	c.AdditionalText = append(c.AdditionalText, (*RestrictedFINXMax350Text)(&value))
}

func (c *CorporateActionNarrative35) AddNarrativeVersion(value string) {
	c.NarrativeVersion = append(c.NarrativeVersion, (*RestrictedFINXMax350Text)(&value))
}

func (c *CorporateActionNarrative35) AddPartyContactNarrative(value string) {
	c.PartyContactNarrative = append(c.PartyContactNarrative, (*RestrictedFINXMax350Text)(&value))
}

func (c *CorporateActionNarrative35) AddTaxationConditions(value string) {
	c.TaxationConditions = append(c.TaxationConditions, (*RestrictedFINXMax350Text)(&value))
}
