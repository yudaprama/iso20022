package model

// Specifies the payment terms of the underlying transaction.
type PaymentTerms2 struct {

	// Specifies payment terms not present in a code list.
	OtherPaymentTerms *Max140Text `xml:"OthrPmtTerms"`

	// Specifies the payment period in coded form and a number of days.
	PaymentCode *PaymentPeriod2 `xml:"PmtCd"`

	// Specifies that the payment conditions apply to a percentage of the amount due.
	Percentage *PercentageRate `xml:"Pctg"`

	// Amount of money to be moved between the debtor and creditor, before deduction of charges, expressed in the currency as ordered by the initiating party.
	Amount *CurrencyAndAmount `xml:"Amt"`
}

func (p *PaymentTerms2) SetOtherPaymentTerms(value string) {
	p.OtherPaymentTerms = (*Max140Text)(&value)
}

func (p *PaymentTerms2) AddPaymentCode() *PaymentPeriod2 {
	p.PaymentCode = new(PaymentPeriod2)
	return p.PaymentCode
}

func (p *PaymentTerms2) SetPercentage(value string) {
	p.Percentage = (*PercentageRate)(&value)
}

func (p *PaymentTerms2) SetAmount(value, currency string) {
	p.Amount = NewCurrencyAndAmount(value, currency)
}
