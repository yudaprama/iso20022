package model

// Provides additional information such as the information conditions.
type CorporateActionNarrative10 struct {

	// Provides additional information or specifies in more detail the content of a message. This field may only be used when the information to be transmitted, cannot be coded.
	AdditionalText []*Max350Text `xml:"AddtlTxt,omitempty"`

	// Provides additional information regarding the party, for example, the contact unit or person responsible for the transaction identified in the message.
	PartyContactNarrative []*Max350Text `xml:"PtyCtctNrrtv,omitempty"`
}

func (c *CorporateActionNarrative10) AddAdditionalText(value string) {
	c.AdditionalText = append(c.AdditionalText, (*Max350Text)(&value))
}

func (c *CorporateActionNarrative10) AddPartyContactNarrative(value string) {
	c.PartyContactNarrative = append(c.PartyContactNarrative, (*Max350Text)(&value))
}
