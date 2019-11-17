package model

// Provides additional information such as the information conditions.
type CorporateActionNarrative19 struct {

	// Provides additional information or specifies in more detail the content of a message. This field may only be used when the information to be transmitted, cannot be coded.
	AdditionalText []*RestrictedFINXMax350Text `xml:"AddtlTxt,omitempty"`

	// Provides additional information regarding the party, for example, the contact unit or person responsible for the transaction identified in the message.
	PartyContactNarrative []*RestrictedFINXMax350Text `xml:"PtyCtctNrrtv,omitempty"`
}

func (c *CorporateActionNarrative19) AddAdditionalText(value string) {
	c.AdditionalText = append(c.AdditionalText, (*RestrictedFINXMax350Text)(&value))
}

func (c *CorporateActionNarrative19) AddPartyContactNarrative(value string) {
	c.PartyContactNarrative = append(c.PartyContactNarrative, (*RestrictedFINXMax350Text)(&value))
}
