package model

// Provides information about the cash movement.
type CashMovement3 struct {

	// Date and time of the posting.
	PostingDateTime *DateAndDateTimeChoice `xml:"PstngDtTm,omitempty"`

	// Value date.
	ValueDate *ISODate `xml:"ValDt"`

	// Cash amount that is posted.
	PostingAmount *ActiveCurrencyAndAmount `xml:"PstngAmt"`

	// Provides information about the account which is debited/credited.
	AccountDetails []*CashAccount18 `xml:"AcctDtls"`
}

func (c *CashMovement3) AddPostingDateTime() *DateAndDateTimeChoice {
	c.PostingDateTime = new(DateAndDateTimeChoice)
	return c.PostingDateTime
}

func (c *CashMovement3) SetValueDate(value string) {
	c.ValueDate = (*ISODate)(&value)
}

func (c *CashMovement3) SetPostingAmount(value, currency string) {
	c.PostingAmount = NewActiveCurrencyAndAmount(value, currency)
}

func (c *CashMovement3) AddAccountDetails() *CashAccount18 {
	newValue := new(CashAccount18)
	c.AccountDetails = append(c.AccountDetails, newValue)
	return newValue
}
