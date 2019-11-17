package model

// Provides the reason for requesting a debit authorisation as well as the amount of the requested debit.
type DebitAuthorisationDetails struct {

	// Indicates the reason for cancellation.
	CancellationReason *CancellationReason1Code `xml:"CxlRsn"`

	// Amount of money to be moved between the debtor and creditor, before deduction of charges, expressed in the currency as ordered by the initiating party.
	AmountToDebit *CurrencyAndAmount `xml:"AmtToDbt,omitempty"`

	// Value date for debiting the amount.
	ValueDateToDebit *ISODate `xml:"ValDtToDbt,omitempty"`
}

func (d *DebitAuthorisationDetails) SetCancellationReason(value string) {
	d.CancellationReason = (*CancellationReason1Code)(&value)
}

func (d *DebitAuthorisationDetails) SetAmountToDebit(value, currency string) {
	d.AmountToDebit = NewCurrencyAndAmount(value, currency)
}

func (d *DebitAuthorisationDetails) SetValueDateToDebit(value string) {
	d.ValueDateToDebit = (*ISODate)(&value)
}
