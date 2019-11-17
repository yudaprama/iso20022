package model

// Provides addtional information such as the taxation conditions.
type CorporateActionNarrative1 struct {

	// Provides conditional information related to the event, eg, an offer is subject to 50% acceptance, the offeror allows the securities holder to set some conditions.
	InformationConditions *Max350Text `xml:"InfConds,omitempty"`

	// Provides information conditions to the account owner that are to be complied with, eg, not open to US/Canadian residents, QIB or SIL to be provided.
	InformationToComplyWith *Max350Text `xml:"InfToCmplyWth,omitempty"`

	// Provides taxation conditions that cannot be included within the structured fields of this message.
	TaxationConditions *Max350Text `xml:"TaxtnConds,omitempty"`

	// Provide  the new name of a company following a name change
	NewCompanyName *Max350Text `xml:"NewCpnyNm,omitempty"`

	// Provides the entity making the offer and is different from the issuing company.
	Offeror *PartyIdentification2Choice `xml:"Offerr,omitempty"`

	// Provides the web address published for the event, ie the address for the Universal Resource Locator (URL), eg, used over the www (HTTP) service.
	URLAddress *Max256Text `xml:"URLAdr,omitempty"`

	// Provides additional information or specifies in more detail the content of a
	// message.
	AdditionalText *Max350Text `xml:"AddtlTxt,omitempty"`
}

func (c *CorporateActionNarrative1) SetInformationConditions(value string) {
	c.InformationConditions = (*Max350Text)(&value)
}

func (c *CorporateActionNarrative1) SetInformationToComplyWith(value string) {
	c.InformationToComplyWith = (*Max350Text)(&value)
}

func (c *CorporateActionNarrative1) SetTaxationConditions(value string) {
	c.TaxationConditions = (*Max350Text)(&value)
}

func (c *CorporateActionNarrative1) SetNewCompanyName(value string) {
	c.NewCompanyName = (*Max350Text)(&value)
}

func (c *CorporateActionNarrative1) AddOfferor() *PartyIdentification2Choice {
	c.Offeror = new(PartyIdentification2Choice)
	return c.Offeror
}

func (c *CorporateActionNarrative1) SetURLAddress(value string) {
	c.URLAddress = (*Max256Text)(&value)
}

func (c *CorporateActionNarrative1) SetAdditionalText(value string) {
	c.AdditionalText = (*Max350Text)(&value)
}
