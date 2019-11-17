package model

// Provides additional information such as the taxation conditions.
type CorporateActionNarrative24 struct {

	// Provides the entity making the offer and is different from the issuing company.
	Offeror []*UpdatedAdditionalInformation3 `xml:"Offerr,omitempty"`

	// Provides the new name of a company following a name change.
	NewCompanyName *UpdatedAdditionalInformation3 `xml:"NewCpnyNm,omitempty"`

	// Provides the web address published for the event, that is, the address for the Universal Resource Locator (URL), for example, used over the www (HTTP) service.
	URLAddress *UpdatedURLlnformation `xml:"URLAdr,omitempty"`
}

func (c *CorporateActionNarrative24) AddOfferor() *UpdatedAdditionalInformation3 {
	newValue := new(UpdatedAdditionalInformation3)
	c.Offeror = append(c.Offeror, newValue)
	return newValue
}

func (c *CorporateActionNarrative24) AddNewCompanyName() *UpdatedAdditionalInformation3 {
	c.NewCompanyName = new(UpdatedAdditionalInformation3)
	return c.NewCompanyName
}

func (c *CorporateActionNarrative24) AddURLAddress() *UpdatedURLlnformation {
	c.URLAddress = new(UpdatedURLlnformation)
	return c.URLAddress
}
