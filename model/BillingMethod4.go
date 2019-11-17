package model

// Provides the details for the tax calculation method C.
type BillingMethod4 struct {

	// Specifies the details of the taxable services using tax calculation method C.
	ServiceDetail []*BillingServiceParameters2 `xml:"SvcDtl"`

	// Total amount of service charge to be taxed in the tax region’s host currency along with the supporting tax calculations.
	//
	// Usage: Used for tax calculation method C only, and only one per tax region may be specified.
	TaxCalculation *TaxCalculation1 `xml:"TaxClctn"`
}

func (b *BillingMethod4) AddServiceDetail() *BillingServiceParameters2 {
	newValue := new(BillingServiceParameters2)
	b.ServiceDetail = append(b.ServiceDetail, newValue)
	return newValue
}

func (b *BillingMethod4) AddTaxCalculation() *TaxCalculation1 {
	b.TaxCalculation = new(TaxCalculation1)
	return b.TaxCalculation
}
