package model

// Specifies the payment terms of the underlying transaction.
type PaymentTerms3 struct {

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

	// Monetary value used as a basis to calculate the discount in these payment terms.
	DiscountAmount *CurrencyAndAmount `xml:"DscntAmt,omitempty"`

	// Percent rate used to calculate the discount for these payment terms.
	DiscountPercentRate *PercentageRate `xml:"DscntPctRate,omitempty"`

	// Monetary value used as a basis to calculate the discount in these payment terms.
	DiscountBasisAmount *CurrencyAndAmount `xml:"DscntBsisAmt,omitempty"`

	// Monetary value used as a basis to calculate the penalty in the payment terms.
	PenaltyAmount *CurrencyAndAmount `xml:"PnltyAmt,omitempty"`

	// Percent rate used to calculate the penalty for these payment terms.
	PenaltyPercentRate *PercentageRate `xml:"PnltyPctRate,omitempty"`

	// Amount used as a basis to calculate the penalty amount.
	PenaltyBasisAmount *CurrencyAndAmount `xml:"PnltyBsisAmt,omitempty"`
}

func (p *PaymentTerms3) SetDueDate(value string) {
	p.DueDate = (*ISODate)(&value)
}

func (p *PaymentTerms3) AddPaymentPeriod() *PaymentPeriod1 {
	p.PaymentPeriod = new(PaymentPeriod1)
	return p.PaymentPeriod
}

func (p *PaymentTerms3) AddDescription(value string) {
	p.Description = append(p.Description, (*Max140Text)(&value))
}

func (p *PaymentTerms3) SetPartialPaymentPercent(value string) {
	p.PartialPaymentPercent = (*PercentageRate)(&value)
}

func (p *PaymentTerms3) AddDirectDebitMandateIdentification(value string) {
	p.DirectDebitMandateIdentification = append(p.DirectDebitMandateIdentification, (*Max35Text)(&value))
}

func (p *PaymentTerms3) SetDiscountAmount(value, currency string) {
	p.DiscountAmount = NewCurrencyAndAmount(value, currency)
}

func (p *PaymentTerms3) SetDiscountPercentRate(value string) {
	p.DiscountPercentRate = (*PercentageRate)(&value)
}

func (p *PaymentTerms3) SetDiscountBasisAmount(value, currency string) {
	p.DiscountBasisAmount = NewCurrencyAndAmount(value, currency)
}

func (p *PaymentTerms3) SetPenaltyAmount(value, currency string) {
	p.PenaltyAmount = NewCurrencyAndAmount(value, currency)
}

func (p *PaymentTerms3) SetPenaltyPercentRate(value string) {
	p.PenaltyPercentRate = (*PercentageRate)(&value)
}

func (p *PaymentTerms3) SetPenaltyBasisAmount(value, currency string) {
	p.PenaltyBasisAmount = NewCurrencyAndAmount(value, currency)
}
