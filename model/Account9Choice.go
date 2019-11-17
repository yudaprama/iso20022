package model

// Choice between a cash account, a charges account or a tax account.
type Account9Choice struct {

	// Account in which cash is maintained.
	CashAccount *CashAccountIdentification6Choice `xml:"CshAcct"`

	// Account to be used for charges if different from the account for payment.
	ChargesAccount *CashAccountIdentification6Choice `xml:"ChrgsAcct"`

	// Account to be used for taxes if different from the account for payment.
	TaxAccount *CashAccountIdentification6Choice `xml:"TaxAcct"`
}

func (a *Account9Choice) AddCashAccount() *CashAccountIdentification6Choice {
	a.CashAccount = new(CashAccountIdentification6Choice)
	return a.CashAccount
}

func (a *Account9Choice) AddChargesAccount() *CashAccountIdentification6Choice {
	a.ChargesAccount = new(CashAccountIdentification6Choice)
	return a.ChargesAccount
}

func (a *Account9Choice) AddTaxAccount() *CashAccountIdentification6Choice {
	a.TaxAccount = new(CashAccountIdentification6Choice)
	return a.TaxAccount
}
