package model

// Specifies an payment schedule, that is an amount that must be paid no sooner than the expected payment date and no later than the due date.
type PaymentDateRange2 struct {

	// Unique and unambiguous identification of the payment schedule.
	PaymentScheduleIdentification *Max35Text `xml:"PmtSchdlId,omitempty"`

	// Amount that must be paid no sooner than the expected payment date and no later than the due date.
	Amount *ActiveCurrencyAndAmount `xml:"Amt"`

	// Expected date whereby the amount must be paid.
	ExpectedDate *ISODate `xml:"XpctdDt,omitempty"`

	// Latest date whereby the amount of money must be paid.
	DueDate *ISODate `xml:"DueDt"`
}

func (p *PaymentDateRange2) SetPaymentScheduleIdentification(value string) {
	p.PaymentScheduleIdentification = (*Max35Text)(&value)
}

func (p *PaymentDateRange2) SetAmount(value, currency string) {
	p.Amount = NewActiveCurrencyAndAmount(value, currency)
}

func (p *PaymentDateRange2) SetExpectedDate(value string) {
	p.ExpectedDate = (*ISODate)(&value)
}

func (p *PaymentDateRange2) SetDueDate(value string) {
	p.DueDate = (*ISODate)(&value)
}
