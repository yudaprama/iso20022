package model

// Taxable service charge amount conversions to host currency.
type BillingServicesAmount2 struct {

	// Sum of the original charge host amount and the service tax host amount values. It represents the total charge for a service (including taxes) expressed in the host currency.
	HostAmount *AmountAndDirection34 `xml:"HstAmt"`

	// Sum of the original charge host amount and the service tax host amount values but expressed in the settlement currency.
	SettlementAmount *AmountAndDirection34 `xml:"SttlmAmt,omitempty"`

	// Sum of the original charge host amount and the service tax host amount values but expressed in the pricing currency.
	PricingAmount *AmountAndDirection34 `xml:"PricgAmt,omitempty"`
}

func (b *BillingServicesAmount2) AddHostAmount() *AmountAndDirection34 {
	b.HostAmount = new(AmountAndDirection34)
	return b.HostAmount
}

func (b *BillingServicesAmount2) AddSettlementAmount() *AmountAndDirection34 {
	b.SettlementAmount = new(AmountAndDirection34)
	return b.SettlementAmount
}

func (b *BillingServicesAmount2) AddPricingAmount() *AmountAndDirection34 {
	b.PricingAmount = new(AmountAndDirection34)
	return b.PricingAmount
}
