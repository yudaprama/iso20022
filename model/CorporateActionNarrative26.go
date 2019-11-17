package model

// Provides additional information such as the taxation conditions.
type CorporateActionNarrative26 struct {

	// Provides the entity making the offer and is different from the issuing company.
	Offeror []*UpdatedAdditionalInformation3 `xml:"Offerr,omitempty"`

	// Provides the new name of a company following a name change.
	NewCompanyName *UpdatedAdditionalInformation3 `xml:"NewCpnyNm,omitempty"`

	// Provides the web address published for the event, that is, the address for the Universal Resource Locator (URL), for example, used over the www (HTTP) service.
	URLAddress *UpdatedURLlnformation2 `xml:"URLAdr,omitempty"`
}

func (c *CorporateActionNarrative26) AddOfferor() *UpdatedAdditionalInformation3 {
	newValue := new(UpdatedAdditionalInformation3)
	c.Offeror = append(c.Offeror, newValue)
	return newValue
}

func (c *CorporateActionNarrative26) AddNewCompanyName() *UpdatedAdditionalInformation3 {
	c.NewCompanyName = new(UpdatedAdditionalInformation3)
	return c.NewCompanyName
}

func (c *CorporateActionNarrative26) AddURLAddress() *UpdatedURLlnformation2 {
	c.URLAddress = new(UpdatedURLlnformation2)
	return c.URLAddress
}
