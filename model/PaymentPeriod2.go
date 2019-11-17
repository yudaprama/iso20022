package model

// Specifies the payment terms by means of a code and a period.
type PaymentPeriod2 struct {

	// Code for the payment.
	Code *PaymentTime2Code `xml:"Cd"`

	// Number of days after which the payment must be effected.
	NumberOfDays *Number `xml:"NbOfDays,omitempty"`
}

func (p *PaymentPeriod2) SetCode(value string) {
	p.Code = (*PaymentTime2Code)(&value)
}

func (p *PaymentPeriod2) SetNumberOfDays(value string) {
	p.NumberOfDays = (*Number)(&value)
}
