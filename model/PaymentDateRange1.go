package model

// Specifies expected and due payment date.
type PaymentDateRange1 struct {

	// Unique and unambiguous identification of the payment schedule.
	PaymentScheduleIdentification *Max35Text `xml:"PmtSchdlId,omitempty"`

	// Expected payment date.
	ExpectedDate *ISODate `xml:"XpctdDt,omitempty"`

	// Latest date whereby the amount must be paid.
	DueDate *ISODate `xml:"DueDt,omitempty"`
}

func (p *PaymentDateRange1) SetPaymentScheduleIdentification(value string) {
	p.PaymentScheduleIdentification = (*Max35Text)(&value)
}

func (p *PaymentDateRange1) SetExpectedDate(value string) {
	p.ExpectedDate = (*ISODate)(&value)
}

func (p *PaymentDateRange1) SetDueDate(value string) {
	p.DueDate = (*ISODate)(&value)
}
