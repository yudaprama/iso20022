package model

// Provides information about the cash option.
type CashOption1 struct {

	// Indicates wether it is a debit or a credit.
	CreditDebitIndicator *CreditDebitCode `xml:"CdtDbtInd"`

	// Currency of the option.
	Currency *ActiveCurrencyCode `xml:"Ccy,omitempty"`

	// Provides information about the dates related to a cash movement.
	DateDetails *CorporateActionDate5 `xml:"DtDtls,omitempty"`

	// Provides information about the amounts related to a cash movement.
	AmountDetails *CorporateActionAmounts1 `xml:"AmtDtls,omitempty"`

	// Provides information about a foreign exchange.
	ExchangeRate *ForeignExchangeTerms8 `xml:"XchgRate,omitempty"`
}

func (c *CashOption1) SetCreditDebitIndicator(value string) {
	c.CreditDebitIndicator = (*CreditDebitCode)(&value)
}

func (c *CashOption1) SetCurrency(value string) {
	c.Currency = (*ActiveCurrencyCode)(&value)
}

func (c *CashOption1) AddDateDetails() *CorporateActionDate5 {
	c.DateDetails = new(CorporateActionDate5)
	return c.DateDetails
}

func (c *CashOption1) AddAmountDetails() *CorporateActionAmounts1 {
	c.AmountDetails = new(CorporateActionAmounts1)
	return c.AmountDetails
}

func (c *CashOption1) AddExchangeRate() *ForeignExchangeTerms8 {
	c.ExchangeRate = new(ForeignExchangeTerms8)
	return c.ExchangeRate
}
