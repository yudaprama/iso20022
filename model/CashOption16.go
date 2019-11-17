package model

// Provides information about the cash option.
type CashOption16 struct {

	// Indicates whether the value is a debit or a credit.
	CreditDebitIndicator *CreditDebitCode `xml:"CdtDbtInd"`

	// Indicates whether the cash payment occurs or will occur in advance of receipt of proceeds from the issuer and based on a contractual agreement established with the account servicer or upon receipt of proceeds from the issuer.
	ContractualPaymentIndicator *Payment1Code `xml:"CtrctlPmtInd,omitempty"`

	// Specifies information regarding outturn resources that cannot be processed by the Central Securities Depository (CSD). Special delivery instruction is required from the account owner for the corporate action outcome to be credited.
	NonEligibleProceedsIndicator *NonEligibleProceedsIndicator1Choice `xml:"NonElgblPrcdsInd,omitempty"`

	// Proceeds are taxable according to the information provided by the issuer / offeror.
	IssuerOfferorTaxabilityIndicator *IssuerTaxability1Code `xml:"IssrOfferrTaxbltyInd,omitempty"`

	// Specifies the type of income.
	// The lists of income type codes to be used, are available on the SMPG website at www.smpg.info.
	IncomeType *GenericIdentification20 `xml:"IncmTp,omitempty"`

	// Identification of the account in which cash is maintained.
	CashAccountIdentification *CashAccountIdentification5Choice `xml:"CshAcctId,omitempty"`

	// Provides information about the amounts related to a cash movement.
	AmountDetails *CorporateActionAmounts16 `xml:"AmtDtls,omitempty"`

	// Provides information about the dates related to a cash movement.
	DateDetails *CorporateActionDate23 `xml:"DtDtls"`

	// Exchange rate between the amount and the resulting amount.
	ForeignExchangeDetails *ForeignExchangeTerms13 `xml:"FXDtls,omitempty"`

	// Provides information about the corporate action option.
	RateAndAmountDetails *RateDetails3 `xml:"RateAndAmtDtls,omitempty"`

	// Provides information about the prices related to a corporate action option.
	PriceDetails *PriceDetails6 `xml:"PricDtls,omitempty"`
}

func (c *CashOption16) SetCreditDebitIndicator(value string) {
	c.CreditDebitIndicator = (*CreditDebitCode)(&value)
}

func (c *CashOption16) SetContractualPaymentIndicator(value string) {
	c.ContractualPaymentIndicator = (*Payment1Code)(&value)
}

func (c *CashOption16) AddNonEligibleProceedsIndicator() *NonEligibleProceedsIndicator1Choice {
	c.NonEligibleProceedsIndicator = new(NonEligibleProceedsIndicator1Choice)
	return c.NonEligibleProceedsIndicator
}

func (c *CashOption16) SetIssuerOfferorTaxabilityIndicator(value string) {
	c.IssuerOfferorTaxabilityIndicator = (*IssuerTaxability1Code)(&value)
}

func (c *CashOption16) AddIncomeType() *GenericIdentification20 {
	c.IncomeType = new(GenericIdentification20)
	return c.IncomeType
}

func (c *CashOption16) AddCashAccountIdentification() *CashAccountIdentification5Choice {
	c.CashAccountIdentification = new(CashAccountIdentification5Choice)
	return c.CashAccountIdentification
}

func (c *CashOption16) AddAmountDetails() *CorporateActionAmounts16 {
	c.AmountDetails = new(CorporateActionAmounts16)
	return c.AmountDetails
}

func (c *CashOption16) AddDateDetails() *CorporateActionDate23 {
	c.DateDetails = new(CorporateActionDate23)
	return c.DateDetails
}

func (c *CashOption16) AddForeignExchangeDetails() *ForeignExchangeTerms13 {
	c.ForeignExchangeDetails = new(ForeignExchangeTerms13)
	return c.ForeignExchangeDetails
}

func (c *CashOption16) AddRateAndAmountDetails() *RateDetails3 {
	c.RateAndAmountDetails = new(RateDetails3)
	return c.RateAndAmountDetails
}

func (c *CashOption16) AddPriceDetails() *PriceDetails6 {
	c.PriceDetails = new(PriceDetails6)
	return c.PriceDetails
}
