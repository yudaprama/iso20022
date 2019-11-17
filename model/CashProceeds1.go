package model

// Provides information about the cash proceeds.
type CashProceeds1 struct {

	// Cash amount which is posted.
	PostingAmount *ActiveCurrencyAndAmount `xml:"PstngAmt"`

	// Reconciliation information.
	ReconciliationDetails *Max350Text `xml:"RcncltnDtls,omitempty"`

	// Provides information about the debited securities account.
	AccountDetails []*CashAccount19 `xml:"AcctDtls"`
}

func (c *CashProceeds1) SetPostingAmount(value, currency string) {
	c.PostingAmount = NewActiveCurrencyAndAmount(value, currency)
}

func (c *CashProceeds1) SetReconciliationDetails(value string) {
	c.ReconciliationDetails = (*Max350Text)(&value)
}

func (c *CashProceeds1) AddAccountDetails() *CashAccount19 {
	newValue := new(CashAccount19)
	c.AccountDetails = append(c.AccountDetails, newValue)
	return newValue
}
