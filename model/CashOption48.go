package model

// Provides information about the cash option.
type CashOption48 struct {

	// Indicates whether the value is a debit or credit.
	CreditDebitIndicator *CreditDebitCode `xml:"CdtDbtInd"`

	// Choice between a cash account, a charges account or a tax account.
	Account *Account9Choice `xml:"Acct,omitempty"`

	// Date/Time of the posting (credit or debit) to the account.
	PostingDate *DateAndDateTimeChoice `xml:"PstngDt"`

	// Date/Time of the posting (credit or debit) to the account that was initially communicated in the confirmation.
	OriginalPostingDate *DateAndDateTimeChoice `xml:"OrgnlPstngDt,omitempty"`

	// Date/time when calculating economic benefit for a cash amount.
	ValueDate *DateAndDateTimeChoice `xml:"ValDt,omitempty"`

	// Amount of money that is to be/was posted to the account.
	PostingAmount *RestrictedFINActiveCurrencyAndAmount `xml:"PstngAmt"`
}

func (c *CashOption48) SetCreditDebitIndicator(value string) {
	c.CreditDebitIndicator = (*CreditDebitCode)(&value)
}

func (c *CashOption48) AddAccount() *Account9Choice {
	c.Account = new(Account9Choice)
	return c.Account
}

func (c *CashOption48) AddPostingDate() *DateAndDateTimeChoice {
	c.PostingDate = new(DateAndDateTimeChoice)
	return c.PostingDate
}

func (c *CashOption48) AddOriginalPostingDate() *DateAndDateTimeChoice {
	c.OriginalPostingDate = new(DateAndDateTimeChoice)
	return c.OriginalPostingDate
}

func (c *CashOption48) AddValueDate() *DateAndDateTimeChoice {
	c.ValueDate = new(DateAndDateTimeChoice)
	return c.ValueDate
}

func (c *CashOption48) SetPostingAmount(value, currency string) {
	c.PostingAmount = NewRestrictedFINActiveCurrencyAndAmount(value, currency)
}
