package model

// Provides the details for the tax calculation method B.
type BillingMethod2 struct {

	// Amount of the original charge expressed in the host currency.
	ServiceChargeHostAmount *AmountAndDirection34 `xml:"SvcChrgHstAmt"`

	// Provides for the regional taxes on the service. Up to three regional taxes may be defined for the same service.
	ServiceTax *BillingServicesAmount1 `xml:"SvcTax"`

	// Provides for the specific tax identification within the same tax region.
	//
	// Usage: This element allows for a maximum of three regional taxes on the same service.
	TaxIdentification []*BillingServicesTax1 `xml:"TaxId"`
}

func (b *BillingMethod2) AddServiceChargeHostAmount() *AmountAndDirection34 {
	b.ServiceChargeHostAmount = new(AmountAndDirection34)
	return b.ServiceChargeHostAmount
}

func (b *BillingMethod2) AddServiceTax() *BillingServicesAmount1 {
	b.ServiceTax = new(BillingServicesAmount1)
	return b.ServiceTax
}

func (b *BillingMethod2) AddTaxIdentification() *BillingServicesTax1 {
	newValue := new(BillingServicesTax1)
	b.TaxIdentification = append(b.TaxIdentification, newValue)
	return newValue
}
