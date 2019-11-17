package model

// Specifies an interest payment schedule, that is an amount that must be paid no sooner than the expected payment date and no later than the due date.
type InterestPaymentDateRange2 struct {

	// Unique and unambiguous identification of the interest payment schedule.
	InterestScheduleIdentification *Max35Text `xml:"IntrstSchdlId,omitempty"`

	// Interest amount that must be paid at due date.
	Amount *ActiveCurrencyAndAmount `xml:"Amt"`

	// Latest date whereby the interest must be paid.
	DueDate *ISODate `xml:"DueDt"`

	// Further details on the interest payments.
	AdditionalInformation *Max1025Text `xml:"AddtlInf,omitempty"`
}

func (i *InterestPaymentDateRange2) SetInterestScheduleIdentification(value string) {
	i.InterestScheduleIdentification = (*Max35Text)(&value)
}

func (i *InterestPaymentDateRange2) SetAmount(value, currency string) {
	i.Amount = NewActiveCurrencyAndAmount(value, currency)
}

func (i *InterestPaymentDateRange2) SetDueDate(value string) {
	i.DueDate = (*ISODate)(&value)
}

func (i *InterestPaymentDateRange2) SetAdditionalInformation(value string) {
	i.AdditionalInformation = (*Max1025Text)(&value)
}
