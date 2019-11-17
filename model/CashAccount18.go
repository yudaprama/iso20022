package model

// Provides information about the cash account.
type CashAccount18 struct {

	// Specifies whether the value is a debit or credit.
	CreditDebitIndicator *CreditDebitCode `xml:"CdtDbtInd"`

	// Identification of the party that owns the account.
	AccountOwnerIdentification *PartyIdentification2Choice `xml:"AcctOwnrId,omitempty"`

	// Identification of the cash account or the securities account from which the cash account is derived.
	AccountIdentification *AccountIdentification2Choice `xml:"AcctId"`

	// The cash balance type.
	BalanceType *CashBalanceType1FormatType `xml:"BalTp,omitempty"`
}

func (c *CashAccount18) SetCreditDebitIndicator(value string) {
	c.CreditDebitIndicator = (*CreditDebitCode)(&value)
}

func (c *CashAccount18) AddAccountOwnerIdentification() *PartyIdentification2Choice {
	c.AccountOwnerIdentification = new(PartyIdentification2Choice)
	return c.AccountOwnerIdentification
}

func (c *CashAccount18) AddAccountIdentification() *AccountIdentification2Choice {
	c.AccountIdentification = new(AccountIdentification2Choice)
	return c.AccountIdentification
}

func (c *CashAccount18) AddBalanceType() *CashBalanceType1FormatType {
	c.BalanceType = new(CashBalanceType1FormatType)
	return c.BalanceType
}
