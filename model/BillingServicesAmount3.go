package model

// Taxable service charge amount conversions to host currency.
type BillingServicesAmount3 struct {

	// Represents the total of all taxable services in a specific tax region for a specific currency.  For example, all taxable services for a tax region in Euro would be totalled here in the Euro currency.
	SourceAmount *AmountAndDirection34 `xml:"SrcAmt"`

	// Represents the total of all taxable services in a specific tax region for a specific currency and is a one-to-one relationship with total taxable charge of services, but represented in the host currency after conversion.
	HostAmount *AmountAndDirection34 `xml:"HstAmt"`
}

func (b *BillingServicesAmount3) AddSourceAmount() *AmountAndDirection34 {
	b.SourceAmount = new(AmountAndDirection34)
	return b.SourceAmount
}

func (b *BillingServicesAmount3) AddHostAmount() *AmountAndDirection34 {
	b.HostAmount = new(AmountAndDirection34)
	return b.HostAmount
}
