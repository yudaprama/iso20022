package model

// Financial loan (instalment) or a recurring transaction.
type RecurringTransaction1 struct {

	// Indicates the recurring/instalment occurrence of the transaction (1 = 1st instalment, 2 = 2nd instalment, etc.).
	SequenceNumber *Max2NumericText `xml:"SeqNb"`

	// Period unit between consecutive payments (for example day, month, year).
	PeriodUnit *Frequency4Code `xml:"PrdUnit"`

	// Number of period units between consecutive payments.
	InstalmentPeriod *Number `xml:"InstlmtPrd"`

	// Total number of instalment payments.
	TotalNumberOfPayments *Number `xml:"TtlNbOfPmts"`

	// Interest charged in percentage for the total amount of payments.
	InterestCharges *ImpliedCurrencyAndAmount `xml:"IntrstChrgs,omitempty"`
}

func (r *RecurringTransaction1) SetSequenceNumber(value string) {
	r.SequenceNumber = (*Max2NumericText)(&value)
}

func (r *RecurringTransaction1) SetPeriodUnit(value string) {
	r.PeriodUnit = (*Frequency4Code)(&value)
}

func (r *RecurringTransaction1) SetInstalmentPeriod(value string) {
	r.InstalmentPeriod = (*Number)(&value)
}

func (r *RecurringTransaction1) SetTotalNumberOfPayments(value string) {
	r.TotalNumberOfPayments = (*Number)(&value)
}

func (r *RecurringTransaction1) SetInterestCharges(value, currency string) {
	r.InterestCharges = NewImpliedCurrencyAndAmount(value, currency)
}
