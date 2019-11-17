package model

// Provides detailed information about an announcement.
type CorporateActionNotificationAdvice1 struct {

	// Provides detailed information about the corporate action event.
	CorporateActionDetails *CorporateAction2 `xml:"CorpActnDtls"`

	// Provides information about an option of a CA event.
	CorporateActionOptionDetails []*CorporateActionOption1 `xml:"CorpActnOptnDtls,omitempty"`
}

func (c *CorporateActionNotificationAdvice1) AddCorporateActionDetails() *CorporateAction2 {
	c.CorporateActionDetails = new(CorporateAction2)
	return c.CorporateActionDetails
}

func (c *CorporateActionNotificationAdvice1) AddCorporateActionOptionDetails() *CorporateActionOption1 {
	newValue := new(CorporateActionOption1)
	c.CorporateActionOptionDetails = append(c.CorporateActionOptionDetails, newValue)
	return newValue
}
