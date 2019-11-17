package model

// Taxable service charge amount conversions to host currency.
type BillingServicesAmount1 struct {

	// Sum of all the individual taxes on the service expressed in the host currency.
	HostAmount *AmountAndDirection34 `xml:"HstAmt"`

	// Amount of the tax obligation expressed in the tax region's pricing currency.
	// Usage: This is the same amount as carried in the host amount but allows the sender to optionally express the value in the pricing currency.
	PricingAmount *AmountAndDirection34 `xml:"PricgAmt,omitempty"`
}

func (b *BillingServicesAmount1) AddHostAmount() *AmountAndDirection34 {
	b.HostAmount = new(AmountAndDirection34)
	return b.HostAmount
}

func (b *BillingServicesAmount1) AddPricingAmount() *AmountAndDirection34 {
	b.PricingAmount = new(AmountAndDirection34)
	return b.PricingAmount
}
