package model

// Specifies the payment terms by means of a code and a period.
type PaymentPeriod3 struct {

	// Code for the payment.
	Code *PaymentTime3Code `xml:"Cd"`

	// Number of days after which the payment must be effected.
	NumberOfDays *Number `xml:"NbOfDays,omitempty"`
}

func (p *PaymentPeriod3) SetCode(value string) {
	p.Code = (*PaymentTime3Code)(&value)
}

func (p *PaymentPeriod3) SetNumberOfDays(value string) {
	p.NumberOfDays = (*Number)(&value)
}
