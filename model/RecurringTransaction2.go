package model

// Financial loan (instalment) or a recurring transaction.
type RecurringTransaction2 struct {

	// Type of instalment plan.
	InstalmentPlan []*InstalmentPlan1Code `xml:"InstlmtPlan,omitempty"`

	// Identification of the instalment plan.
	PlanIdentification *Max35Text `xml:"PlanId,omitempty"`

	// Indicates the recurring/instalment occurrence of the transaction (1 = 1st instalment, 2 = 2nd instalment, etc.).
	SequenceNumber *Number `xml:"SeqNb,omitempty"`

	// Period unit between consecutive payments (for example day, month, year).
	PeriodUnit *Frequency3Code `xml:"PrdUnit,omitempty"`

	// Number of period units between consecutive payments.
	InstalmentPeriod *Number `xml:"InstlmtPrd,omitempty"`

	// Total number of instalment payments.
	TotalNumberOfPayments *Number `xml:"TtlNbOfPmts,omitempty"`

	// Date of the first payment.
	FirstPaymentDate *ISODate `xml:"FrstPmtDt,omitempty"`

	// Cumulative amount of all the instalments.
	TotalAmount *CurrencyAndAmount `xml:"TtlAmt,omitempty"`

	// Amount of the first payment.
	FirstAmount *ImpliedCurrencyAndAmount `xml:"FrstAmt,omitempty"`

	// Charges related to the transaction.
	Charges *ImpliedCurrencyAndAmount `xml:"Chrgs,omitempty"`
}

func (r *RecurringTransaction2) AddInstalmentPlan(value string) {
	r.InstalmentPlan = append(r.InstalmentPlan, (*InstalmentPlan1Code)(&value))
}

func (r *RecurringTransaction2) SetPlanIdentification(value string) {
	r.PlanIdentification = (*Max35Text)(&value)
}

func (r *RecurringTransaction2) SetSequenceNumber(value string) {
	r.SequenceNumber = (*Number)(&value)
}

func (r *RecurringTransaction2) SetPeriodUnit(value string) {
	r.PeriodUnit = (*Frequency3Code)(&value)
}

func (r *RecurringTransaction2) SetInstalmentPeriod(value string) {
	r.InstalmentPeriod = (*Number)(&value)
}

func (r *RecurringTransaction2) SetTotalNumberOfPayments(value string) {
	r.TotalNumberOfPayments = (*Number)(&value)
}

func (r *RecurringTransaction2) SetFirstPaymentDate(value string) {
	r.FirstPaymentDate = (*ISODate)(&value)
}

func (r *RecurringTransaction2) SetTotalAmount(value, currency string) {
	r.TotalAmount = NewCurrencyAndAmount(value, currency)
}

func (r *RecurringTransaction2) SetFirstAmount(value, currency string) {
	r.FirstAmount = NewImpliedCurrencyAndAmount(value, currency)
}

func (r *RecurringTransaction2) SetCharges(value, currency string) {
	r.Charges = NewImpliedCurrencyAndAmount(value, currency)
}
