package model

// Choice of format for the payment direction.
type PaymentDirection2Choice struct {

	// Indicates the direction of payment for asset or mortgage backed securities, ie, whether the repaid capital is distributed (payment direction is down) or capitalized (payment direction is up).
	Indicator *PaymentDirectionIndicator `xml:"Ind"`

	// Payment direction expressed as a proprietary code.
	Proprietary *GenericIdentification20 `xml:"Prtry"`
}

func (p *PaymentDirection2Choice) SetIndicator(value string) {
	p.Indicator = (*PaymentDirectionIndicator)(&value)
}

func (p *PaymentDirection2Choice) AddProprietary() *GenericIdentification20 {
	p.Proprietary = new(GenericIdentification20)
	return p.Proprietary
}
