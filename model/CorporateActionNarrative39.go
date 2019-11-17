package model

// Provides additional information such as the taxation conditions.
type CorporateActionNarrative39 struct {

	// Provides the entity making the offer and is different from the issuing company.
	Offeror []*UpdatedAdditionalInformation6 `xml:"Offerr,omitempty"`

	// Provides the new name of a company following a name change.
	NewCompanyName *UpdatedAdditionalInformation6 `xml:"NewCpnyNm,omitempty"`

	// Provides the web address published for the event, that is, the address for the Universal Resource Locator (URL), for example, used over the www (HTTP) service.
	URLAddress *UpdatedURLlnformation3 `xml:"URLAdr,omitempty"`
}

func (c *CorporateActionNarrative39) AddOfferor() *UpdatedAdditionalInformation6 {
	newValue := new(UpdatedAdditionalInformation6)
	c.Offeror = append(c.Offeror, newValue)
	return newValue
}

func (c *CorporateActionNarrative39) AddNewCompanyName() *UpdatedAdditionalInformation6 {
	c.NewCompanyName = new(UpdatedAdditionalInformation6)
	return c.NewCompanyName
}

func (c *CorporateActionNarrative39) AddURLAddress() *UpdatedURLlnformation3 {
	c.URLAddress = new(UpdatedURLlnformation3)
	return c.URLAddress
}
