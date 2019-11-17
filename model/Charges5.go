package model

// Charges related to a payment obligation contracted between two financial institutions related to the financing of a commercial transaction.
type Charges5 struct {

	// Bank which will pay the charges.
	ChargesPayer *BankRole1Code `xml:"ChrgsPyer"`

	// Bank which will receive the charges.
	ChargesPayee *BankRole1Code `xml:"ChrgsPyee"`

	// Amount of the charges taken by the payer.
	Amount *CurrencyAndAmount `xml:"Amt,omitempty"`

	// Amount of the charges expressed as a percentage of the amount paid by the obligor bank.
	Percentage *PercentageRate `xml:"Pctg,omitempty"`

	// Type of charges. For example: transaction charges, financing charges, deferred payment, interests.
	Type *Max35Text `xml:"Tp,omitempty"`
}

func (c *Charges5) SetChargesPayer(value string) {
	c.ChargesPayer = (*BankRole1Code)(&value)
}

func (c *Charges5) SetChargesPayee(value string) {
	c.ChargesPayee = (*BankRole1Code)(&value)
}

func (c *Charges5) SetAmount(value, currency string) {
	c.Amount = NewCurrencyAndAmount(value, currency)
}

func (c *Charges5) SetPercentage(value string) {
	c.Percentage = (*PercentageRate)(&value)
}

func (c *Charges5) SetType(value string) {
	c.Type = (*Max35Text)(&value)
}
