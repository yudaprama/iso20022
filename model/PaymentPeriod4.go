package model

// Specifies the payment terms by means of a code and a period.
type PaymentPeriod4 struct {

	// Code for the payment.
	Code *PaymentTime4Code `xml:"Cd"`

	// Number of days after which the payment must be effected.
	NumberOfDays *Number `xml:"NbOfDays,omitempty"`
}

func (p *PaymentPeriod4) SetCode(value string) {
	p.Code = (*PaymentTime4Code)(&value)
}

func (p *PaymentPeriod4) SetNumberOfDays(value string) {
	p.NumberOfDays = (*Number)(&value)
}
