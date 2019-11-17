package model

// Specifies the payment terms of the underlying transaction.
type PaymentTerms5 struct {

	// Specifies the payment terms using a code or other means.
	PaymentTerms *PaymentCodeOrOther2Choice `xml:"PmtTerms"`

	// Specifies if it is a fixed amount or a percentage.
	AmountOrPercentage *AmountOrPercentage2Choice `xml:"AmtOrPctg"`
}

func (p *PaymentTerms5) AddPaymentTerms() *PaymentCodeOrOther2Choice {
	p.PaymentTerms = new(PaymentCodeOrOther2Choice)
	return p.PaymentTerms
}

func (p *PaymentTerms5) AddAmountOrPercentage() *AmountOrPercentage2Choice {
	p.AmountOrPercentage = new(AmountOrPercentage2Choice)
	return p.AmountOrPercentage
}
