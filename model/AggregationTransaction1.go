package model

// Payment transaction with an aggregated amount.
type AggregationTransaction1 struct {

	// Date and time of the first payment.
	FirstPaymentDateTime *ISODateTime `xml:"FrstPmtDtTm,omitempty"`

	// Date and time of the last payment.
	LastPaymentDateTime *ISODateTime `xml:"LastPmtDtTm,omitempty"`

	// Total number of payments that has been aggregated.
	NumberOfPayments *Number `xml:"NbOfPmts,omitempty"`

	// Individual payment that has been aggregated.
	IndividualPayment []*DetailedAmount6 `xml:"IndvPmt,omitempty"`
}

func (a *AggregationTransaction1) SetFirstPaymentDateTime(value string) {
	a.FirstPaymentDateTime = (*ISODateTime)(&value)
}

func (a *AggregationTransaction1) SetLastPaymentDateTime(value string) {
	a.LastPaymentDateTime = (*ISODateTime)(&value)
}

func (a *AggregationTransaction1) SetNumberOfPayments(value string) {
	a.NumberOfPayments = (*Number)(&value)
}

func (a *AggregationTransaction1) AddIndividualPayment() *DetailedAmount6 {
	newValue := new(DetailedAmount6)
	a.IndividualPayment = append(a.IndividualPayment, newValue)
	return newValue
}
