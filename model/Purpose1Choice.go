package model

// Underlying reason for the payment transaction.
//
// Usage: Purpose is used by the end-customers, i.e. initiating party, (ultimate) debtor, (ultimate) creditor to provide information concerning the nature of the payment. Purpose is a content element, which is not used for processing by any of the agents involved in the payment chain.
type Purpose1Choice struct {

	// Specifies the underlying reason for the payment transaction, as published in an external purpose code list.
	Code *ExternalPurposeCode `xml:"Cd"`

	// User community specific purpose.
	//
	// Usage: When available, codes provided by local communities should be used.
	Proprietary *Max35Text `xml:"Prtry"`
}

func (p *Purpose1Choice) SetCode(value string) {
	p.Code = (*ExternalPurposeCode)(&value)
}

func (p *Purpose1Choice) SetProprietary(value string) {
	p.Proprietary = (*Max35Text)(&value)
}
