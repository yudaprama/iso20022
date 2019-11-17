package model

// Provides information about the cash movement.
type CashMovement1 struct {

	// Identification of the movement.
	MovementIdentification *Max35Text `xml:"MvmntId,omitempty"`

	// Cash amount.
	Amount *ActiveCurrencyAndAmount `xml:"Amt"`

	// Amount of taxes.
	TaxAmount *ActiveCurrencyAndAmount `xml:"TaxAmt,omitempty"`

	// Specifies the charges.
	Charges []*Charges1 `xml:"Chrgs,omitempty"`

	// Provides information about the account which is debited/credited.
	AccountDetails []*CashAccount18 `xml:"AcctDtls"`
}

func (c *CashMovement1) SetMovementIdentification(value string) {
	c.MovementIdentification = (*Max35Text)(&value)
}

func (c *CashMovement1) SetAmount(value, currency string) {
	c.Amount = NewActiveCurrencyAndAmount(value, currency)
}

func (c *CashMovement1) SetTaxAmount(value, currency string) {
	c.TaxAmount = NewActiveCurrencyAndAmount(value, currency)
}

func (c *CashMovement1) AddCharges() *Charges1 {
	newValue := new(Charges1)
	c.Charges = append(c.Charges, newValue)
	return newValue
}

func (c *CashMovement1) AddAccountDetails() *CashAccount18 {
	newValue := new(CashAccount18)
	c.AccountDetails = append(c.AccountDetails, newValue)
	return newValue
}
