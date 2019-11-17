package model

// Payment transaction with an aggregated amount.
type AggregationTransaction2 struct {

	// Date and time of the first payment.
	FirstPaymentDateTime *ISODateTime `xml:"FrstPmtDtTm,omitempty"`

	// Date and time of the last payment.
	LastPaymentDateTime *ISODateTime `xml:"LastPmtDtTm,omitempty"`

	// Total number of payments that has been aggregated.
	NumberOfPayments *Number `xml:"NbOfPmts,omitempty"`

	// Individual payment that has been aggregated.
	IndividualPayment []*DetailedAmount14 `xml:"IndvPmt,omitempty"`
}

func (a *AggregationTransaction2) SetFirstPaymentDateTime(value string) {
	a.FirstPaymentDateTime = (*ISODateTime)(&value)
}

func (a *AggregationTransaction2) SetLastPaymentDateTime(value string) {
	a.LastPaymentDateTime = (*ISODateTime)(&value)
}

func (a *AggregationTransaction2) SetNumberOfPayments(value string) {
	a.NumberOfPayments = (*Number)(&value)
}

func (a *AggregationTransaction2) AddIndividualPayment() *DetailedAmount14 {
	newValue := new(DetailedAmount14)
	a.IndividualPayment = append(a.IndividualPayment, newValue)
	return newValue
}
