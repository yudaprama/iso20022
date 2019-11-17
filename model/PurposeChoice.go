package model

// Underlying reason for the payment transaction, eg, a charity payment, or a commercial agreement between the creditor and the debtor.
type PurposeChoice struct {

	// Payment transaction purpose that is specific to a user community and is required for use within that user community.
	Proprietary *Max35Text `xml:"Prtry"`

	// Specifies the type of transaction that resulted in the payment initiation in coded form.
	Code *PaymentPurpose1Code `xml:"Cd"`
}

func (p *PurposeChoice) SetProprietary(value string) {
	p.Proprietary = (*Max35Text)(&value)
}

func (p *PurposeChoice) SetCode(value string) {
	p.Code = (*PaymentPurpose1Code)(&value)
}
