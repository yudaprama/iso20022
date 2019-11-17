package model

// Parties used for acting parties that apply either to the whole message or to individual sides.
type ConfirmationPartyDetails5 struct {

	// Unique and unambiguous identifier for an organisation that is allocated by an institution, eg, Dun & Bradstreet Identification.
	Identification *PartyIdentification32Choice `xml:"Id"`

	// Alternate identification for a party.
	AlternateIdentification *AlternatePartyIdentification5 `xml:"AltrnId,omitempty"`

	// Unambiguous identification of the transaction for the party identified.
	ProcessingIdentification *Max35Text `xml:"PrcgId,omitempty"`

	// Provides additional information to a party identification.
	AdditionalInformation *PartyTextInformation5 `xml:"AddtlInf,omitempty"`

	// Indicates whether the confirmation party is a member of the investor protection association required, eg, as per regulation.
	InvestorProtectionAssociationMembership *YesNoIndicator `xml:"InvstrPrtcnAssoctnMmbsh,omitempty"`
}

func (c *ConfirmationPartyDetails5) AddIdentification() *PartyIdentification32Choice {
	c.Identification = new(PartyIdentification32Choice)
	return c.Identification
}

func (c *ConfirmationPartyDetails5) AddAlternateIdentification() *AlternatePartyIdentification5 {
	c.AlternateIdentification = new(AlternatePartyIdentification5)
	return c.AlternateIdentification
}

func (c *ConfirmationPartyDetails5) SetProcessingIdentification(value string) {
	c.ProcessingIdentification = (*Max35Text)(&value)
}

func (c *ConfirmationPartyDetails5) AddAdditionalInformation() *PartyTextInformation5 {
	c.AdditionalInformation = new(PartyTextInformation5)
	return c.AdditionalInformation
}

func (c *ConfirmationPartyDetails5) SetInvestorProtectionAssociationMembership(value string) {
	c.InvestorProtectionAssociationMembership = (*YesNoIndicator)(&value)
}
