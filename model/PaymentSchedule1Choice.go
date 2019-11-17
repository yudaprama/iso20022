package model

// Choice between expected/due payment date and a payment schedule per amount and due date.
type PaymentSchedule1Choice struct {

	// Specifies an expected date and a due date for the payment.
	DateRange *PaymentDateRange1 `xml:"DtRg"`

	// Specifies a payment sub-schedule, that is the amount of money that must be paid no sooner than the expected date and no later than the latest shipment date.
	SubSchedule []*PaymentDateRange2 `xml:"SubSchdl"`
}

func (p *PaymentSchedule1Choice) AddDateRange() *PaymentDateRange1 {
	p.DateRange = new(PaymentDateRange1)
	return p.DateRange
}

func (p *PaymentSchedule1Choice) AddSubSchedule() *PaymentDateRange2 {
	newValue := new(PaymentDateRange2)
	p.SubSchedule = append(p.SubSchedule, newValue)
	return newValue
}
