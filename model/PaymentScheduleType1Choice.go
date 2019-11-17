package model

// Choice for payment schedule type.
type PaymentScheduleType1Choice struct {

	// Payment schedule type defined in a coded form.
	Code *PaymentScheduleType1Code `xml:"Cd"`

	// Payment schedule type defined in a proprietary format.
	Proprietary *Max35Text `xml:"Prtry"`
}

func (p *PaymentScheduleType1Choice) SetCode(value string) {
	p.Code = (*PaymentScheduleType1Code)(&value)
}

func (p *PaymentScheduleType1Choice) SetProprietary(value string) {
	p.Proprietary = (*Max35Text)(&value)
}
