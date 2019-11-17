package model

// Provides the reason for requesting a debit authorisation as well as the amount of the requested debit.
type DebitAuthorisation2 struct {

	// Specifies the reason for the cancellation request.
	CancellationReason *CancellationReason33Choice `xml:"CxlRsn"`

	// Amount of money requested for debit authorisation.
	AmountToDebit *ActiveOrHistoricCurrencyAndAmount `xml:"AmtToDbt,omitempty"`

	// Value date for debiting the amount.
	ValueDateToDebit *ISODate `xml:"ValDtToDbt,omitempty"`

	// Further details on the cancellation request reason.
	AdditionalCancellationReasonInformation []*Max105Text `xml:"AddtlCxlRsnInf,omitempty"`
}

func (d *DebitAuthorisation2) AddCancellationReason() *CancellationReason33Choice {
	d.CancellationReason = new(CancellationReason33Choice)
	return d.CancellationReason
}

func (d *DebitAuthorisation2) SetAmountToDebit(value, currency string) {
	d.AmountToDebit = NewActiveOrHistoricCurrencyAndAmount(value, currency)
}

func (d *DebitAuthorisation2) SetValueDateToDebit(value string) {
	d.ValueDateToDebit = (*ISODate)(&value)
}

func (d *DebitAuthorisation2) AddAdditionalCancellationReasonInformation(value string) {
	d.AdditionalCancellationReasonInformation = append(d.AdditionalCancellationReasonInformation, (*Max105Text)(&value))
}
