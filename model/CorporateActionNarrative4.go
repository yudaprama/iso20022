package model

// Provides additional information such as the taxation conditions.
type CorporateActionNarrative4 struct {

	// Provides declaration details narrative relative to the financial instrument, for example, beneficial ownership.
	DeclarationDetails []*Max350Text `xml:"DclrtnDtls,omitempty"`

	// Provides additional information or specifies in more detail the content of a message. This field may only be used when the information to be transmitted, cannot be coded.
	AdditionalText []*Max350Text `xml:"AddtlTxt,omitempty"`

	// Provides information that can be ignored for automated processing; - reiteration of information that has been included within structured fields of this message, - or narrative information not needed for automatic processing.
	NarrativeVersion []*Max350Text `xml:"NrrtvVrsn,omitempty"`

	// Provides information required for the registration.
	RegistrationDetails []*Max350Text `xml:"RegnDtls,omitempty"`

	// Provides conditional information related to the event, for example, an offer is subject to 50 percent acceptance, the offeror allows the securities holder to set some conditions.
	InformationConditions []*Max350Text `xml:"InfConds,omitempty"`

	// Provides information conditions to the account owner that are to be complied with, for example, not open to US/Canadian residents, Qualified Institutional Buyers (QIB) or SIL (Sophisticated Investor Letter) to be provided.
	InformationToComplyWith []*Max350Text `xml:"InfToCmplyWth,omitempty"`

	// Provides additional information regarding the party, for example, the contact unit or person responsible for the transaction identified in the message.
	PartyContactNarrative []*Max350Text `xml:"PtyCtctNrrtv,omitempty"`

	// Provides taxation conditions that cannot be included within the structured fields of this message and has not been mentioned in the Service Level Agreement (SLA)
	TaxationConditions []*Max350Text `xml:"TaxtnConds,omitempty"`

	// Provides additional information on the basket or index underlying a security, for example a warrant.
	BasketOrIndexInformation []*Max350Text `xml:"BsktOrIndxInf,omitempty"`
}

func (c *CorporateActionNarrative4) AddDeclarationDetails(value string) {
	c.DeclarationDetails = append(c.DeclarationDetails, (*Max350Text)(&value))
}

func (c *CorporateActionNarrative4) AddAdditionalText(value string) {
	c.AdditionalText = append(c.AdditionalText, (*Max350Text)(&value))
}

func (c *CorporateActionNarrative4) AddNarrativeVersion(value string) {
	c.NarrativeVersion = append(c.NarrativeVersion, (*Max350Text)(&value))
}

func (c *CorporateActionNarrative4) AddRegistrationDetails(value string) {
	c.RegistrationDetails = append(c.RegistrationDetails, (*Max350Text)(&value))
}

func (c *CorporateActionNarrative4) AddInformationConditions(value string) {
	c.InformationConditions = append(c.InformationConditions, (*Max350Text)(&value))
}

func (c *CorporateActionNarrative4) AddInformationToComplyWith(value string) {
	c.InformationToComplyWith = append(c.InformationToComplyWith, (*Max350Text)(&value))
}

func (c *CorporateActionNarrative4) AddPartyContactNarrative(value string) {
	c.PartyContactNarrative = append(c.PartyContactNarrative, (*Max350Text)(&value))
}

func (c *CorporateActionNarrative4) AddTaxationConditions(value string) {
	c.TaxationConditions = append(c.TaxationConditions, (*Max350Text)(&value))
}

func (c *CorporateActionNarrative4) AddBasketOrIndexInformation(value string) {
	c.BasketOrIndexInformation = append(c.BasketOrIndexInformation, (*Max350Text)(&value))
}
