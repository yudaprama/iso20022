package model

// Choice between expected/due interest payment date and a interest payment schedule per amount and due date.
type InterestPaymentSchedule1Choice struct {

	// Specifies an expected date and a due date for the interest payment.
	DateRange *InterestPaymentDateRange1 `xml:"DtRg"`

	// Specifies an interest payment schedule, that is an interest amount that must be paid no sooner than the expected payment date and no later than the due date.
	SubSchedule []*InterestPaymentDateRange2 `xml:"SubSchdl"`
}

func (i *InterestPaymentSchedule1Choice) AddDateRange() *InterestPaymentDateRange1 {
	i.DateRange = new(InterestPaymentDateRange1)
	return i.DateRange
}

func (i *InterestPaymentSchedule1Choice) AddSubSchedule() *InterestPaymentDateRange2 {
	newValue := new(InterestPaymentDateRange2)
	i.SubSchedule = append(i.SubSchedule, newValue)
	return newValue
}
