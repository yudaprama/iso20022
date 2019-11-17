package model

// Total of credit or debit transactions
type TransactionTotals5 struct {

	// Cumulative amount of all financial transactions.
	Amount *ImpliedCurrencyAndAmount `xml:"Amt"`

	// Number of all financial transactions.
	Number *Number `xml:"Nb"`

	// Cumulative amount of all chargeback transactions exclusive of any fees.
	ChargeBackAmount *ImpliedCurrencyAndAmount `xml:"ChrgBckAmt"`

	// Total number of chargeback transactions.
	ChargeBackNumber *Number `xml:"ChrgBckNb"`

	// Cumulative amount of all reversal transactions exclusive of any fees.
	ReversalAmount *ImpliedCurrencyAndAmount `xml:"RvslAmt"`

	// Total number of reversal transactions.
	ReversalNumber *Number `xml:"RvslNb"`

	// Sum amount of all fees.
	FeeAmounts *ImpliedCurrencyAndAmount `xml:"FeeAmts,omitempty"`
}

func (t *TransactionTotals5) SetAmount(value, currency string) {
	t.Amount = NewImpliedCurrencyAndAmount(value, currency)
}

func (t *TransactionTotals5) SetNumber(value string) {
	t.Number = (*Number)(&value)
}

func (t *TransactionTotals5) SetChargeBackAmount(value, currency string) {
	t.ChargeBackAmount = NewImpliedCurrencyAndAmount(value, currency)
}

func (t *TransactionTotals5) SetChargeBackNumber(value string) {
	t.ChargeBackNumber = (*Number)(&value)
}

func (t *TransactionTotals5) SetReversalAmount(value, currency string) {
	t.ReversalAmount = NewImpliedCurrencyAndAmount(value, currency)
}

func (t *TransactionTotals5) SetReversalNumber(value string) {
	t.ReversalNumber = (*Number)(&value)
}

func (t *TransactionTotals5) SetFeeAmounts(value, currency string) {
	t.FeeAmounts = NewImpliedCurrencyAndAmount(value, currency)
}
