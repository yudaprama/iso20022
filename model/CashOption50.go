package model

// Provides information about the cash option.
type CashOption50 struct {

	// Indicates whether the value is a debit or a credit.
	CreditDebitIndicator *CreditDebitCode `xml:"CdtDbtInd"`

	// Specifies information regarding outturn resources that cannot be processed by the Central Securities Depository (CSD). Special delivery instruction is required from the account owner for the corporate action outcome to be credited.
	NonEligibleProceedsIndicator *NonEligibleProceedsIndicator3Choice `xml:"NonElgblPrcdsInd,omitempty"`

	// Proceeds are taxable according to the information provided by the issuer / offeror.
	IssuerOfferorTaxabilityIndicator *IssuerOfferorTaxabilityIndicator1Choice `xml:"IssrOfferrTaxbltyInd,omitempty"`

	// Specifies the type of income.
	// The lists of income type codes to be used, are available on the SMPG website at www.smpg.info.
	IncomeType *GenericIdentification30 `xml:"IncmTp,omitempty"`

	// Specifies the basis for the reduced rate of withholding.
	ExemptionType []*GenericIdentification30 `xml:"XmptnTp,omitempty"`

	// Indicates the country from which the income originates.
	CountryOfIncomeSource *CountryCode `xml:"CtryOfIncmSrc,omitempty"`

	// Identification of the account in which cash is maintained.
	CashAccountIdentification *CashAccountIdentification5Choice `xml:"CshAcctId,omitempty"`

	// Provides information about the amounts related to a cash movement.
	AmountDetails *CorporateActionAmounts36 `xml:"AmtDtls,omitempty"`

	// Provides information about the dates related to a cash movement.
	DateDetails *CorporateActionDate47 `xml:"DtDtls"`

	// Exchange rate between the amount and the resulting amount
	ForeignExchangeDetails *ForeignExchangeTerms24 `xml:"FXDtls,omitempty"`

	// Provides information about the corporate action option.
	RateAndAmountDetails *RateDetails26 `xml:"RateAndAmtDtls,omitempty"`

	// Provides information about the prices related to a corporate action option.
	PriceDetails *PriceDetails22 `xml:"PricDtls,omitempty"`
}

func (c *CashOption50) SetCreditDebitIndicator(value string) {
	c.CreditDebitIndicator = (*CreditDebitCode)(&value)
}

func (c *CashOption50) AddNonEligibleProceedsIndicator() *NonEligibleProceedsIndicator3Choice {
	c.NonEligibleProceedsIndicator = new(NonEligibleProceedsIndicator3Choice)
	return c.NonEligibleProceedsIndicator
}

func (c *CashOption50) AddIssuerOfferorTaxabilityIndicator() *IssuerOfferorTaxabilityIndicator1Choice {
	c.IssuerOfferorTaxabilityIndicator = new(IssuerOfferorTaxabilityIndicator1Choice)
	return c.IssuerOfferorTaxabilityIndicator
}

func (c *CashOption50) AddIncomeType() *GenericIdentification30 {
	c.IncomeType = new(GenericIdentification30)
	return c.IncomeType
}

func (c *CashOption50) AddExemptionType() *GenericIdentification30 {
	newValue := new(GenericIdentification30)
	c.ExemptionType = append(c.ExemptionType, newValue)
	return newValue
}

func (c *CashOption50) SetCountryOfIncomeSource(value string) {
	c.CountryOfIncomeSource = (*CountryCode)(&value)
}

func (c *CashOption50) AddCashAccountIdentification() *CashAccountIdentification5Choice {
	c.CashAccountIdentification = new(CashAccountIdentification5Choice)
	return c.CashAccountIdentification
}

func (c *CashOption50) AddAmountDetails() *CorporateActionAmounts36 {
	c.AmountDetails = new(CorporateActionAmounts36)
	return c.AmountDetails
}

func (c *CashOption50) AddDateDetails() *CorporateActionDate47 {
	c.DateDetails = new(CorporateActionDate47)
	return c.DateDetails
}

func (c *CashOption50) AddForeignExchangeDetails() *ForeignExchangeTerms24 {
	c.ForeignExchangeDetails = new(ForeignExchangeTerms24)
	return c.ForeignExchangeDetails
}

func (c *CashOption50) AddRateAndAmountDetails() *RateDetails26 {
	c.RateAndAmountDetails = new(RateDetails26)
	return c.RateAndAmountDetails
}

func (c *CashOption50) AddPriceDetails() *PriceDetails22 {
	c.PriceDetails = new(PriceDetails22)
	return c.PriceDetails
}
