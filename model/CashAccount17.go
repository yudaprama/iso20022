package model

// Provides information about the cash account.
type CashAccount17 struct {

	// Identification of the cash account.
	AccountIdentification *CashAccountIdentification1Choice `xml:"AcctId"`

	// Currency of the payment.
	PaymentCurrency *ActiveCurrencyCode `xml:"PmtCcy"`

	// Identification of the party that owns the account.
	AccountOwnerIdentification *PartyIdentification2Choice `xml:"AcctOwnrId,omitempty"`

	// Identification of the cash correspondent back.
	CorrespondentBankIdentification *BICIdentifier `xml:"CrspdtBkId"`
}

func (c *CashAccount17) AddAccountIdentification() *CashAccountIdentification1Choice {
	c.AccountIdentification = new(CashAccountIdentification1Choice)
	return c.AccountIdentification
}

func (c *CashAccount17) SetPaymentCurrency(value string) {
	c.PaymentCurrency = (*ActiveCurrencyCode)(&value)
}

func (c *CashAccount17) AddAccountOwnerIdentification() *PartyIdentification2Choice {
	c.AccountOwnerIdentification = new(PartyIdentification2Choice)
	return c.AccountOwnerIdentification
}

func (c *CashAccount17) SetCorrespondentBankIdentification(value string) {
	c.CorrespondentBankIdentification = (*BICIdentifier)(&value)
}
