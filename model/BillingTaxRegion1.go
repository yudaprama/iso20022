package model

// Tax region that levies a tax on the services in a statement.
type BillingTaxRegion1 struct {

	// Specifies a particular unique zone of taxes assigned by taxing authorities.  A tax region number is unique.  Every account is considered to reside within a tax region, although some tax regions may not charge taxes on services.
	RegionNumber *Max40Text `xml:"RgnNb"`

	// Name associated with a specific tax region number.
	RegionName *Max40Text `xml:"RgnNm"`

	// Specifies the financial institution’s customer tax identification number.
	//
	// Usage:
	// This is the number passed from the financial institution to the taxing authority to indicate the specific customer tied to the taxes calculated for this tax region and group of delivered services.  It is typically the tax identification tied to a customer’s account.
	CustomerTaxIdentification *Max40Text `xml:"CstmrTaxId"`

	// Date on which the tax calculation was performed.
	//
	// Usage:
	// This date can be used to verify the tax rate value on the calculation date.
	PointDate *ISODate `xml:"PtDt,omitempty"`

	// Tax information that relates to the sending financial institution.
	SendingFinancialInstitution *BillingTaxIdentification1 `xml:"SndgFI,omitempty"`

	// Unique number to be used by the customer to cross-reference between the tax region information and a tax invoice or notice.
	InvoiceNumber *Max40Text `xml:"InvcNb,omitempty"`

	// Tax values are based on tax calculation method C.
	MethodC *BillingMethod4 `xml:"MtdC,omitempty"`

	// Total tax amount expressed in the account’s settlement (or charging) currency.
	//
	// Usage: This total sums the individual service level taxes as calculated for each service by methods A, B and D. The sum of these amounts across all tax regions for the statement is displayed as the tax total sum in the compensation section.
	SettlementAmount *AmountAndDirection34 `xml:"SttlmAmt"`

	// Total amount of all taxes for a specific customer within the tax region expressed in the tax region’s host currency.
	//
	// Usage: It is the same value as total tax amount and is included for the specific use of tax calculation methods A , B and D.
	TaxDueToRegion *AmountAndDirection34 `xml:"TaxDueToRgn"`
}

func (b *BillingTaxRegion1) SetRegionNumber(value string) {
	b.RegionNumber = (*Max40Text)(&value)
}

func (b *BillingTaxRegion1) SetRegionName(value string) {
	b.RegionName = (*Max40Text)(&value)
}

func (b *BillingTaxRegion1) SetCustomerTaxIdentification(value string) {
	b.CustomerTaxIdentification = (*Max40Text)(&value)
}

func (b *BillingTaxRegion1) SetPointDate(value string) {
	b.PointDate = (*ISODate)(&value)
}

func (b *BillingTaxRegion1) AddSendingFinancialInstitution() *BillingTaxIdentification1 {
	b.SendingFinancialInstitution = new(BillingTaxIdentification1)
	return b.SendingFinancialInstitution
}

func (b *BillingTaxRegion1) SetInvoiceNumber(value string) {
	b.InvoiceNumber = (*Max40Text)(&value)
}

func (b *BillingTaxRegion1) AddMethodC() *BillingMethod4 {
	b.MethodC = new(BillingMethod4)
	return b.MethodC
}

func (b *BillingTaxRegion1) AddSettlementAmount() *AmountAndDirection34 {
	b.SettlementAmount = new(AmountAndDirection34)
	return b.SettlementAmount
}

func (b *BillingTaxRegion1) AddTaxDueToRegion() *AmountAndDirection34 {
	b.TaxDueToRegion = new(AmountAndDirection34)
	return b.TaxDueToRegion
}
