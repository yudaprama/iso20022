package model

// Describes account taxing parameters.
type AccountTax1 struct {

	// Defines the calculation method on how the taxes are applied on the account.
	CalculationMethod *BillingTaxCalculationMethod1Code `xml:"ClctnMtd"`

	// Identifies the tax region in which the account resides.
	Region *Max40Text `xml:"Rgn,omitempty"`

	// Specifies the country of residence, when the account owner does not reside in the account's tax region.
	//
	// Usage: If present, the account owner does not reside in the account's tax region.
	NonResidenceCountry *ResidenceLocation1Choice `xml:"NonResCtry,omitempty"`
}

func (a *AccountTax1) SetCalculationMethod(value string) {
	a.CalculationMethod = (*BillingTaxCalculationMethod1Code)(&value)
}

func (a *AccountTax1) SetRegion(value string) {
	a.Region = (*Max40Text)(&value)
}

func (a *AccountTax1) AddNonResidenceCountry() *ResidenceLocation1Choice {
	a.NonResidenceCountry = new(ResidenceLocation1Choice)
	return a.NonResidenceCountry
}
