package model

// Specifies the payment terms of the underlying transaction.
type PaymentTerms4 struct {

	// Specifies the payment terms using a code or other means.
	PaymentTerms *PaymentCodeOrOther1Choice `xml:"PmtTerms"`

	// Specifies if it is a fixed amount or a percentage.
	AmountOrPercentage *AmountOrPercentage2Choice `xml:"AmtOrPctg"`
}

func (p *PaymentTerms4) AddPaymentTerms() *PaymentCodeOrOther1Choice {
	p.PaymentTerms = new(PaymentCodeOrOther1Choice)
	return p.PaymentTerms
}

func (p *PaymentTerms4) AddAmountOrPercentage() *AmountOrPercentage2Choice {
	p.AmountOrPercentage = new(AmountOrPercentage2Choice)
	return p.AmountOrPercentage
}
