package model

// Provides for regional taxes on the service.
type BillingServicesTax1 struct {

	// Identification number of the specific region tax used to calculate the tax.
	Number *Max35Text `xml:"Nb"`

	// Name used to describe the tax (such as the national value added tax).
	Description *Max40Text `xml:"Desc,omitempty"`

	// Rate used to calculate the tax.
	Rate *DecimalNumber `xml:"Rate"`

	// Amount of the tax obligation expressed in the tax region's host currency.
	HostAmount *AmountAndDirection34 `xml:"HstAmt"`

	// Amount of the tax obligation expressed in the tax region's pricing currency.
	PricingAmount *AmountAndDirection34 `xml:"PricgAmt,omitempty"`
}

func (b *BillingServicesTax1) SetNumber(value string) {
	b.Number = (*Max35Text)(&value)
}

func (b *BillingServicesTax1) SetDescription(value string) {
	b.Description = (*Max40Text)(&value)
}

func (b *BillingServicesTax1) SetRate(value string) {
	b.Rate = (*DecimalNumber)(&value)
}

func (b *BillingServicesTax1) AddHostAmount() *AmountAndDirection34 {
	b.HostAmount = new(AmountAndDirection34)
	return b.HostAmount
}

func (b *BillingServicesTax1) AddPricingAmount() *AmountAndDirection34 {
	b.PricingAmount = new(AmountAndDirection34)
	return b.PricingAmount
}
