package model

// Provides information about the cash option.
type CashOption49 struct {

	// Indicates whether the value is a debit or a credit.
	CreditDebitIndicator *CreditDebitCode `xml:"CdtDbtInd"`

	// Specifies information regarding outturn resources that cannot be processed by the Central Securities Depository (CSD). Special delivery instruction is required from the account owner for the corporate action outcome to be credited.
	NonEligibleProceedsIndicator *NonEligibleProceedsIndicator4Choice `xml:"NonElgblPrcdsInd,omitempty"`

	// Proceeds are taxable according to the information provided by the issuer / offeror.
	IssuerOfferorTaxabilityIndicator *GenericIdentification47 `xml:"IssrOfferrTaxbltyInd,omitempty"`

	// Specifies the type of income.
	// The lists of income type codes to be used, are available on the SMPG website at www.smpg.info.
	IncomeType *GenericIdentification47 `xml:"IncmTp,omitempty"`

	// Specifies the basis for the reduced rate of withholding.
	ExemptionType []*GenericIdentification47 `xml:"XmptnTp,omitempty"`

	// Indicates the country from which the income originates.
	CountryOfIncomeSource *CountryCode `xml:"CtryOfIncmSrc,omitempty"`

	// Identification of the account in which cash is maintained.
	CashAccountIdentification *CashAccountIdentification6Choice `xml:"CshAcctId,omitempty"`

	// Provides information about the amounts related to a cash movement.
	AmountDetails *CorporateActionAmounts41 `xml:"AmtDtls,omitempty"`

	// Provides information about the dates related to a cash movement.
	DateDetails *CorporateActionDate56 `xml:"DtDtls"`

	// Exchange rate between the amount and the resulting amount
	ForeignExchangeDetails *ForeignExchangeTerms28 `xml:"FXDtls,omitempty"`

	// Provides information about the corporate action option.
	RateAndAmountDetails *RateDetails25 `xml:"RateAndAmtDtls,omitempty"`

	// Provides information about the prices related to a corporate action option.
	PriceDetails *PriceDetails25 `xml:"PricDtls,omitempty"`
}

func (c *CashOption49) SetCreditDebitIndicator(value string) {
	c.CreditDebitIndicator = (*CreditDebitCode)(&value)
}

func (c *CashOption49) AddNonEligibleProceedsIndicator() *NonEligibleProceedsIndicator4Choice {
	c.NonEligibleProceedsIndicator = new(NonEligibleProceedsIndicator4Choice)
	return c.NonEligibleProceedsIndicator
}

func (c *CashOption49) AddIssuerOfferorTaxabilityIndicator() *GenericIdentification47 {
	c.IssuerOfferorTaxabilityIndicator = new(GenericIdentification47)
	return c.IssuerOfferorTaxabilityIndicator
}

func (c *CashOption49) AddIncomeType() *GenericIdentification47 {
	c.IncomeType = new(GenericIdentification47)
	return c.IncomeType
}

func (c *CashOption49) AddExemptionType() *GenericIdentification47 {
	newValue := new(GenericIdentification47)
	c.ExemptionType = append(c.ExemptionType, newValue)
	return newValue
}

func (c *CashOption49) SetCountryOfIncomeSource(value string) {
	c.CountryOfIncomeSource = (*CountryCode)(&value)
}

func (c *CashOption49) AddCashAccountIdentification() *CashAccountIdentification6Choice {
	c.CashAccountIdentification = new(CashAccountIdentification6Choice)
	return c.CashAccountIdentification
}

func (c *CashOption49) AddAmountDetails() *CorporateActionAmounts41 {
	c.AmountDetails = new(CorporateActionAmounts41)
	return c.AmountDetails
}

func (c *CashOption49) AddDateDetails() *CorporateActionDate56 {
	c.DateDetails = new(CorporateActionDate56)
	return c.DateDetails
}

func (c *CashOption49) AddForeignExchangeDetails() *ForeignExchangeTerms28 {
	c.ForeignExchangeDetails = new(ForeignExchangeTerms28)
	return c.ForeignExchangeDetails
}

func (c *CashOption49) AddRateAndAmountDetails() *RateDetails25 {
	c.RateAndAmountDetails = new(RateDetails25)
	return c.RateAndAmountDetails
}

func (c *CashOption49) AddPriceDetails() *PriceDetails25 {
	c.PriceDetails = new(PriceDetails25)
	return c.PriceDetails
}
