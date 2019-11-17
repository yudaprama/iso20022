package model

// Specifies the underlying reason for the payment transaction.
// Usage: Purpose is used by the end-customers, that is initiating party, (ultimate) debtor, (ultimate) creditor to provide information concerning the nature of the payment. Purpose is a content element, which is not used for processing by any of the agents involved in the payment chain.
type Purpose2Choice struct {

	// Underlying reason for the payment transaction, as published in an external purpose code list.
	Code *ExternalPurpose1Code `xml:"Cd"`

	// Purpose, in a proprietary form.
	Proprietary *Max35Text `xml:"Prtry"`
}

func (p *Purpose2Choice) SetCode(value string) {
	p.Code = (*ExternalPurpose1Code)(&value)
}

func (p *Purpose2Choice) SetProprietary(value string) {
	p.Proprietary = (*Max35Text)(&value)
}
