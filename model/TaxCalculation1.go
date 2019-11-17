package model

// Specifies the details for the tax calculation.
type TaxCalculation1 struct {

	// Currency that all totals for taxable services must be converted to for calculating taxes owed for this tax region.  This also is the currency in which the payment of tax obligations is usually submitted to the taxing authority.
	HostCurrency *ActiveOrHistoricCurrencyCode `xml:"HstCcy"`

	// Taxable service charge amount conversions to host currency.
	//
	// Usage: One occurrence must be present for each different service pricing currency in the statement.
	TaxableServiceChargeConversion []*BillingServicesAmount3 `xml:"TaxblSvcChrgConvs"`

	// Total of all services subject to tax for a specific tax region.
	//
	// Usage:
	// This field will equal the sum of all the separate host tax charge for service equivalent totals for each individual currency.  It is expressed in the tax region’s Host currency. This total is used to determine the tax due by calculating using each tax identifications rate.
	TotalTaxableServiceChargeHostAmount *AmountAndDirection34 `xml:"TtlTaxblSvcChrgHstAmt"`

	// Provides for the specific tax identification within the same tax region.
	//
	// Usage: A maximum of three specific tax identifications may be provided. These elements use the total host currency taxable amount as the basis of the calculation.
	// This element is only valid for method C.
	TaxIdentification []*BillingServicesTax3 `xml:"TaxId"`

	// Total amount of all taxes for a specific customer within the tax region.  This is a sum of all individual total tax amounts for tax identification ’s expressed in the tax region’s host currency.
	TotalTax *AmountAndDirection34 `xml:"TtlTax"`
}

func (t *TaxCalculation1) SetHostCurrency(value string) {
	t.HostCurrency = (*ActiveOrHistoricCurrencyCode)(&value)
}

func (t *TaxCalculation1) AddTaxableServiceChargeConversion() *BillingServicesAmount3 {
	newValue := new(BillingServicesAmount3)
	t.TaxableServiceChargeConversion = append(t.TaxableServiceChargeConversion, newValue)
	return newValue
}

func (t *TaxCalculation1) AddTotalTaxableServiceChargeHostAmount() *AmountAndDirection34 {
	t.TotalTaxableServiceChargeHostAmount = new(AmountAndDirection34)
	return t.TotalTaxableServiceChargeHostAmount
}

func (t *TaxCalculation1) AddTaxIdentification() *BillingServicesTax3 {
	newValue := new(BillingServicesTax3)
	t.TaxIdentification = append(t.TaxIdentification, newValue)
	return newValue
}

func (t *TaxCalculation1) AddTotalTax() *AmountAndDirection34 {
	t.TotalTax = new(AmountAndDirection34)
	return t.TotalTax
}
