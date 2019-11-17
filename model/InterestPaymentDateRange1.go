package model

// Specifies an expected date and a due date for the interest payment.
type InterestPaymentDateRange1 struct {

	// Unique and unambiguous identification of the interest payment schedule.
	InterestScheduleIdentification *Max35Text `xml:"IntrstSchdlId,omitempty"`

	// Expected interest payment date.
	ExpectedDate *ISODate `xml:"XpctdDt,omitempty"`

	// Latest date whereby the interest must be paid.
	DueDate *ISODate `xml:"DueDt,omitempty"`
}

func (i *InterestPaymentDateRange1) SetInterestScheduleIdentification(value string) {
	i.InterestScheduleIdentification = (*Max35Text)(&value)
}

func (i *InterestPaymentDateRange1) SetExpectedDate(value string) {
	i.ExpectedDate = (*ISODate)(&value)
}

func (i *InterestPaymentDateRange1) SetDueDate(value string) {
	i.DueDate = (*ISODate)(&value)
}
