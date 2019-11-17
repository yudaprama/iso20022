package model

// Specifies the payment terms of the underlying transaction.
type PaymentTerms6 struct {

	// Due date specified for the payment terms.
	DueDate *ISODate `xml:"DueDt,omitempty"`

	// Payment period specified for these payment terms.
	PaymentPeriod *PaymentPeriod1 `xml:"PmtPrd,omitempty"`

	// Textual description of these payment terms.
	Description []*Max140Text `xml:"Desc,omitempty"`

	// Partial payment, expressed as a percentage, for the payment terms.
	PartialPaymentPercent *PercentageRate `xml:"PrtlPmtPct,omitempty"`

	// Direct debit mandate identification specified for these payment terms.
	DirectDebitMandateIdentification []*Max35Text `xml:"DrctDbtMndtId,omitempty"`

	// Amount used as a basis to calculate the discount amount for these payment terms.
	BasisAmount *CurrencyAndAmount `xml:"BsisAmt,omitempty"`

	// Amount of money that results from the application of an agreed discount percentage to the basis amount and payable to the creditor.
	DiscountAmount *CurrencyAndAmount `xml:"DscntAmt,omitempty"`

	// Percent rate used to calculate the discount for these payment terms.
	DiscountPercentRate *PercentageRate `xml:"DscntPctRate,omitempty"`

	// Amount of money that results from the application of an agreed penalty percentage to the basis amount and payable by the creditor.
	PenaltyAmount *CurrencyAndAmount `xml:"PnltyAmt,omitempty"`

	// Percent rate used to calculate the penalty for these payment terms.
	PenaltyPercentRate *PercentageRate `xml:"PnltyPctRate,omitempty"`
}

func (p *PaymentTerms6) SetDueDate(value string) {
	p.DueDate = (*ISODate)(&value)
}

func (p *PaymentTerms6) AddPaymentPeriod() *PaymentPeriod1 {
	p.PaymentPeriod = new(PaymentPeriod1)
	return p.PaymentPeriod
}

func (p *PaymentTerms6) AddDescription(value string) {
	p.Description = append(p.Description, (*Max140Text)(&value))
}

func (p *PaymentTerms6) SetPartialPaymentPercent(value string) {
	p.PartialPaymentPercent = (*PercentageRate)(&value)
}

func (p *PaymentTerms6) AddDirectDebitMandateIdentification(value string) {
	p.DirectDebitMandateIdentification = append(p.DirectDebitMandateIdentification, (*Max35Text)(&value))
}

func (p *PaymentTerms6) SetBasisAmount(value, currency string) {
	p.BasisAmount = NewCurrencyAndAmount(value, currency)
}

func (p *PaymentTerms6) SetDiscountAmount(value, currency string) {
	p.DiscountAmount = NewCurrencyAndAmount(value, currency)
}

func (p *PaymentTerms6) SetDiscountPercentRate(value string) {
	p.DiscountPercentRate = (*PercentageRate)(&value)
}

func (p *PaymentTerms6) SetPenaltyAmount(value, currency string) {
	p.PenaltyAmount = NewCurrencyAndAmount(value, currency)
}

func (p *PaymentTerms6) SetPenaltyPercentRate(value string) {
	p.PenaltyPercentRate = (*PercentageRate)(&value)
}
