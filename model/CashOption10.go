package model

// Provides information about the cash option.
type CashOption10 struct {

	// Indicates whether the value is a debit or a credit.
	CreditDebitIndicator *CreditDebitCode `xml:"CdtDbtInd"`

	// Specifies information regarding outturn resources that cannot be processed by the Central Securities Depository (CSD). Special delivery instruction is required from the account owner for the corporate action outcome to be credited.
	NonEligibleProceedsIndicator *NonEligibleProceedsIndicator1Choice `xml:"NonElgblPrcdsInd,omitempty"`

	// Specifies the type of income.
	// The lists of income type codes to be used, are available on the SMPG website at www.smpg.info.
	IncomeType *GenericIdentification20 `xml:"IncmTp,omitempty"`

	// Identification of the account in which cash is maintained.
	CashAccountIdentification *CashAccountIdentification5Choice `xml:"CshAcctId,omitempty"`

	// Provides information about the amounts related to a cash movement.
	AmountDetails *CorporateActionAmounts10 `xml:"AmtDtls,omitempty"`

	// Provides information about the dates related to a cash movement.
	DateDetails *CorporateActionDate17 `xml:"DtDtls"`

	// Exchange rate between the amount and the resulting amount
	ForeignExchangeDetails *ForeignExchangeTerms13 `xml:"FXDtls,omitempty"`

	// Provides information about the corporate action option.
	RateAndAmountDetails *RateDetails3 `xml:"RateAndAmtDtls,omitempty"`

	// Provides information about the prices related to a corporate action option.
	PriceDetails *PriceDetails3 `xml:"PricDtls,omitempty"`
}

func (c *CashOption10) SetCreditDebitIndicator(value string) {
	c.CreditDebitIndicator = (*CreditDebitCode)(&value)
}

func (c *CashOption10) AddNonEligibleProceedsIndicator() *NonEligibleProceedsIndicator1Choice {
	c.NonEligibleProceedsIndicator = new(NonEligibleProceedsIndicator1Choice)
	return c.NonEligibleProceedsIndicator
}

func (c *CashOption10) AddIncomeType() *GenericIdentification20 {
	c.IncomeType = new(GenericIdentification20)
	return c.IncomeType
}

func (c *CashOption10) AddCashAccountIdentification() *CashAccountIdentification5Choice {
	c.CashAccountIdentification = new(CashAccountIdentification5Choice)
	return c.CashAccountIdentification
}

func (c *CashOption10) AddAmountDetails() *CorporateActionAmounts10 {
	c.AmountDetails = new(CorporateActionAmounts10)
	return c.AmountDetails
}

func (c *CashOption10) AddDateDetails() *CorporateActionDate17 {
	c.DateDetails = new(CorporateActionDate17)
	return c.DateDetails
}

func (c *CashOption10) AddForeignExchangeDetails() *ForeignExchangeTerms13 {
	c.ForeignExchangeDetails = new(ForeignExchangeTerms13)
	return c.ForeignExchangeDetails
}

func (c *CashOption10) AddRateAndAmountDetails() *RateDetails3 {
	c.RateAndAmountDetails = new(RateDetails3)
	return c.RateAndAmountDetails
}

func (c *CashOption10) AddPriceDetails() *PriceDetails3 {
	c.PriceDetails = new(PriceDetails3)
	return c.PriceDetails
}
