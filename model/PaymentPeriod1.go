package model

// Specifies the payment terms by means of a code and a period.
type PaymentPeriod1 struct {

	// Code for the payment.
	Code *PaymentTime1Code `xml:"Cd"`

	// Number of days after which the payment must be effected.
	NumberOfDays *Number `xml:"NbOfDays,omitempty"`
}

func (p *PaymentPeriod1) SetCode(value string) {
	p.Code = (*PaymentTime1Code)(&value)
}

func (p *PaymentPeriod1) SetNumberOfDays(value string) {
	p.NumberOfDays = (*Number)(&value)
}
