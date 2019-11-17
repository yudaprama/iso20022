package model

// Provides the details for the tax calculation method D.
type BillingMethod3 struct {

	// Equivalent amount to the service tax host amount but allows the sender to optionally express the value in the pricing currency.
	ServiceTaxPriceAmount *AmountAndDirection34 `xml:"SvcTaxPricAmt"`

	// Provides for the specific tax identification within the same tax region.
	//
	// Usage: This element allows for a maximum of three regional taxes on the same service.
	TaxIdentification []*BillingServicesTax2 `xml:"TaxId"`
}

func (b *BillingMethod3) AddServiceTaxPriceAmount() *AmountAndDirection34 {
	b.ServiceTaxPriceAmount = new(AmountAndDirection34)
	return b.ServiceTaxPriceAmount
}

func (b *BillingMethod3) AddTaxIdentification() *BillingServicesTax2 {
	newValue := new(BillingServicesTax2)
	b.TaxIdentification = append(b.TaxIdentification, newValue)
	return newValue
}
