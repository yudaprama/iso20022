package model

// Choice between a cash account, a charges account or a tax account.
type Account8Choice struct {

	// Account in which cash is maintained.
	CashAccount *CashAccountIdentification5Choice `xml:"CshAcct"`

	// Account to be used for charges if different from the account for payment.
	ChargesAccount *CashAccountIdentification5Choice `xml:"ChrgsAcct"`

	// Account to be used for taxes if different from the account for payment.
	TaxAccount *CashAccountIdentification5Choice `xml:"TaxAcct"`
}

func (a *Account8Choice) AddCashAccount() *CashAccountIdentification5Choice {
	a.CashAccount = new(CashAccountIdentification5Choice)
	return a.CashAccount
}

func (a *Account8Choice) AddChargesAccount() *CashAccountIdentification5Choice {
	a.ChargesAccount = new(CashAccountIdentification5Choice)
	return a.ChargesAccount
}

func (a *Account8Choice) AddTaxAccount() *CashAccountIdentification5Choice {
	a.TaxAccount = new(CashAccountIdentification5Choice)
	return a.TaxAccount
}
