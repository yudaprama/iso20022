package model

// Provides information about the cash movement resulting from the election instruction.
type CorporateActionCashMovements2 struct {

	// Posting identification of the cash movement.
	PostingIdentification *Max35Text `xml:"PstngId,omitempty"`

	// Posting date of the cash movement.
	PostingDateTime *DateAndDateTimeChoice `xml:"PstngDtTm,omitempty"`

	// Amount posted as a result of the cash movement.
	PostingAmount *ActiveCurrencyAndAmount `xml:"PstngAmt"`

	// Provides information about the account which is debited/credited as a result of the movement.
	AccountDetails []*CashAccount19 `xml:"AcctDtls"`
}

func (c *CorporateActionCashMovements2) SetPostingIdentification(value string) {
	c.PostingIdentification = (*Max35Text)(&value)
}

func (c *CorporateActionCashMovements2) AddPostingDateTime() *DateAndDateTimeChoice {
	c.PostingDateTime = new(DateAndDateTimeChoice)
	return c.PostingDateTime
}

func (c *CorporateActionCashMovements2) SetPostingAmount(value, currency string) {
	c.PostingAmount = NewActiveCurrencyAndAmount(value, currency)
}

func (c *CorporateActionCashMovements2) AddAccountDetails() *CashAccount19 {
	newValue := new(CashAccount19)
	c.AccountDetails = append(c.AccountDetails, newValue)
	return newValue
}
