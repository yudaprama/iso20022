package model

// Provides the reason for requesting a debit authorisation as well as the amount of the requested debit.
type DebitAuthorisation1 struct {

	// Specifies the reason for the cancellation request.
	CancellationReason *CancellationReason14Choice `xml:"CxlRsn"`

	// Amount of money requested for debit authorisation.
	AmountToDebit *ActiveOrHistoricCurrencyAndAmount `xml:"AmtToDbt,omitempty"`

	// Value date for debiting the amount.
	ValueDateToDebit *ISODate `xml:"ValDtToDbt,omitempty"`

	// Further details on the cancellation request reason.
	AdditionalCancellationReasonInformation []*Max105Text `xml:"AddtlCxlRsnInf,omitempty"`
}

func (d *DebitAuthorisation1) AddCancellationReason() *CancellationReason14Choice {
	d.CancellationReason = new(CancellationReason14Choice)
	return d.CancellationReason
}

func (d *DebitAuthorisation1) SetAmountToDebit(value, currency string) {
	d.AmountToDebit = NewActiveOrHistoricCurrencyAndAmount(value, currency)
}

func (d *DebitAuthorisation1) SetValueDateToDebit(value string) {
	d.ValueDateToDebit = (*ISODate)(&value)
}

func (d *DebitAuthorisation1) AddAdditionalCancellationReasonInformation(value string) {
	d.AdditionalCancellationReasonInformation = append(d.AdditionalCancellationReasonInformation, (*Max105Text)(&value))
}
