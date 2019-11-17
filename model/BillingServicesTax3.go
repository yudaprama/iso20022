package model

// Provides for regional taxes on the service.
type BillingServicesTax3 struct {

	// Identification number of the specific region tax used to calculate the tax.
	Number *Max35Text `xml:"Nb"`

	// Name used to describe the tax (such as the national value added tax).
	Description *Max40Text `xml:"Desc,omitempty"`

	// Rate used to calculate the tax.
	Rate *DecimalNumber `xml:"Rate"`

	// Specifies the tax obligation for taxable services within a tax region for a specific tax identifier (such as national value added tax equals 34,00), and expressed in the tax region’s host currency.
	TotalTaxAmount *AmountAndDirection34 `xml:"TtlTaxAmt"`
}

func (b *BillingServicesTax3) SetNumber(value string) {
	b.Number = (*Max35Text)(&value)
}

func (b *BillingServicesTax3) SetDescription(value string) {
	b.Description = (*Max40Text)(&value)
}

func (b *BillingServicesTax3) SetRate(value string) {
	b.Rate = (*DecimalNumber)(&value)
}

func (b *BillingServicesTax3) AddTotalTaxAmount() *AmountAndDirection34 {
	b.TotalTaxAmount = new(AmountAndDirection34)
	return b.TotalTaxAmount
}
