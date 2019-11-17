package model

// Provides information about the cash option.
type CashOption32 struct {

	// Indicates whether the value is a debit or a credit.
	CreditDebitIndicator *CreditDebitCode `xml:"CdtDbtInd"`

	// Specifies information regarding outturn resources that cannot be processed by the Central Securities Depository (CSD). Special delivery instruction is required from the account owner for the corporate action outcome to be credited.
	NonEligibleProceedsIndicator *NonEligibleProceedsIndicator1Choice `xml:"NonElgblPrcdsInd,omitempty"`

	// Proceeds are taxable according to the information provided by the issuer / offeror.
	IssuerOfferorTaxabilityIndicator *IssuerTaxability1Code `xml:"IssrOfferrTaxbltyInd,omitempty"`

	// Specifies the type of income.
	// The lists of income type codes to be used, are available on the SMPG website at www.smpg.info.
	IncomeType *GenericIdentification20 `xml:"IncmTp,omitempty"`

	// Specifies the basis for the reduced rate of withholding.
	ExemptionType []*GenericIdentification20 `xml:"XmptnTp,omitempty"`

	// Identification of the account in which cash is maintained.
	CashAccountIdentification *CashAccountIdentification5Choice `xml:"CshAcctId,omitempty"`

	// Provides information about the amounts related to a cash movement.
	AmountDetails *CorporateActionAmounts27 `xml:"AmtDtls,omitempty"`

	// Provides information about the dates related to a cash movement.
	DateDetails *CorporateActionDate23 `xml:"DtDtls"`

	// Exchange rate between the amount and the resulting amount
	ForeignExchangeDetails *ForeignExchangeTerms13 `xml:"FXDtls,omitempty"`

	// Provides information about the corporate action option.
	RateAndAmountDetails *RateDetails14 `xml:"RateAndAmtDtls,omitempty"`

	// Provides information about the prices related to a corporate action option.
	PriceDetails *PriceDetails15 `xml:"PricDtls,omitempty"`
}

func (c *CashOption32) SetCreditDebitIndicator(value string) {
	c.CreditDebitIndicator = (*CreditDebitCode)(&value)
}

func (c *CashOption32) AddNonEligibleProceedsIndicator() *NonEligibleProceedsIndicator1Choice {
	c.NonEligibleProceedsIndicator = new(NonEligibleProceedsIndicator1Choice)
	return c.NonEligibleProceedsIndicator
}

func (c *CashOption32) SetIssuerOfferorTaxabilityIndicator(value string) {
	c.IssuerOfferorTaxabilityIndicator = (*IssuerTaxability1Code)(&value)
}

func (c *CashOption32) AddIncomeType() *GenericIdentification20 {
	c.IncomeType = new(GenericIdentification20)
	return c.IncomeType
}

func (c *CashOption32) AddExemptionType() *GenericIdentification20 {
	newValue := new(GenericIdentification20)
	c.ExemptionType = append(c.ExemptionType, newValue)
	return newValue
}

func (c *CashOption32) AddCashAccountIdentification() *CashAccountIdentification5Choice {
	c.CashAccountIdentification = new(CashAccountIdentification5Choice)
	return c.CashAccountIdentification
}

func (c *CashOption32) AddAmountDetails() *CorporateActionAmounts27 {
	c.AmountDetails = new(CorporateActionAmounts27)
	return c.AmountDetails
}

func (c *CashOption32) AddDateDetails() *CorporateActionDate23 {
	c.DateDetails = new(CorporateActionDate23)
	return c.DateDetails
}

func (c *CashOption32) AddForeignExchangeDetails() *ForeignExchangeTerms13 {
	c.ForeignExchangeDetails = new(ForeignExchangeTerms13)
	return c.ForeignExchangeDetails
}

func (c *CashOption32) AddRateAndAmountDetails() *RateDetails14 {
	c.RateAndAmountDetails = new(RateDetails14)
	return c.RateAndAmountDetails
}

func (c *CashOption32) AddPriceDetails() *PriceDetails15 {
	c.PriceDetails = new(PriceDetails15)
	return c.PriceDetails
}
