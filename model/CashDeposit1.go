package model

// Deposit of an amount of money defined in cash notes and/or coins.
type CashDeposit1 struct {

	// Specifies the note or coin denomination, including the currency, such as a 50 euro note.
	NoteDenomination *ActiveCurrencyAndAmount `xml:"NoteDnmtn"`

	// Specifies the number of notes of the same denomination in the deposit.
	NumberOfNotes *Max15NumericText `xml:"NbOfNotes"`

	// Specifies the total amount of money in the cash deposit, that is the note denomination times the number of notes.
	Amount *ActiveCurrencyAndAmount `xml:"Amt"`
}

func (c *CashDeposit1) SetNoteDenomination(value, currency string) {
	c.NoteDenomination = NewActiveCurrencyAndAmount(value, currency)
}

func (c *CashDeposit1) SetNumberOfNotes(value string) {
	c.NumberOfNotes = (*Max15NumericText)(&value)
}

func (c *CashDeposit1) SetAmount(value, currency string) {
	c.Amount = NewActiveCurrencyAndAmount(value, currency)
}
