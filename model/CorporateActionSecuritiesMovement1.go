package model

// Information about the securities movements.
type CorporateActionSecuritiesMovement1 struct {

	// Date and time of the posting.
	PostingDateTime *DateAndDateTimeChoice `xml:"PstngDtTm,omitempty"`

	// Identification of the financial instrument.
	SecurityIdentification *SecurityIdentification7 `xml:"SctyId"`

	// Quantity of securities that is posted.
	PostingQuantity *UnitOrFaceAmount1Choice `xml:"PstngQty"`

	// Provides information about the account which is debited/credited.
	AccountDetails []*SecuritiesAccount8 `xml:"AcctDtls"`
}

func (c *CorporateActionSecuritiesMovement1) AddPostingDateTime() *DateAndDateTimeChoice {
	c.PostingDateTime = new(DateAndDateTimeChoice)
	return c.PostingDateTime
}

func (c *CorporateActionSecuritiesMovement1) AddSecurityIdentification() *SecurityIdentification7 {
	c.SecurityIdentification = new(SecurityIdentification7)
	return c.SecurityIdentification
}

func (c *CorporateActionSecuritiesMovement1) AddPostingQuantity() *UnitOrFaceAmount1Choice {
	c.PostingQuantity = new(UnitOrFaceAmount1Choice)
	return c.PostingQuantity
}

func (c *CorporateActionSecuritiesMovement1) AddAccountDetails() *SecuritiesAccount8 {
	newValue := new(SecuritiesAccount8)
	c.AccountDetails = append(c.AccountDetails, newValue)
	return newValue
}
