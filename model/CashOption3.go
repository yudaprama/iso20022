package model

// Provides information about the cash option.
type CashOption3 struct {

	// Indicates whether the value is a debit or a credit.
	CreditDebitIndicator *CreditDebitCode `xml:"CdtDbtInd"`

	// Indicates whether the cash payment occurs or will occur in advance of receipt of proceeds from the issuer and based on a contractual agreement established with the account servicer or upon receipt of proceeds from the issuer.
	ContractualPaymentIndicator *Payment1Code `xml:"CtrctlPmtInd,omitempty"`

	// Specifies information regarding outturn resources that cannot be processed by the Central Securities Depository (CSD). Special delivery instruction is required from the account owner for the corporate action outcome to be credited.
	NonEligibleProceedsIndicator *NonEligibleProceedsIndicator1Choice `xml:"NonElgblPrcdsInd,omitempty"`

	// Specifies the type of income.
	// The lists of income type codes to be used, are available on the SMPG website at www.smpg.info.
	IncomeType *GenericIdentification20 `xml:"IncmTp,omitempty"`

	// Identification of the account in which cash is maintained.
	CashAccountIdentification *CashAccountIdentification5Choice `xml:"CshAcctId,omitempty"`

	// Provides information about the amounts related to a cash movement.
	AmountDetails *CorporateActionAmounts3 `xml:"AmtDtls,omitempty"`

	// Provides information about the dates related to a cash movement.
	DateDetails *CorporateActionDate9 `xml:"DtDtls"`

	// Exchange rate between the amount and the resulting amount.
	ForeignExchangeDetails *ForeignExchangeTerms11 `xml:"FXDtls,omitempty"`

	// Generic cash price received per product by the underlying security holder either as a percentage or an amount, for example, redemption price.
	GenericCashPriceReceivedPerProduct *PriceFormat10Choice `xml:"GncCshPricRcvdPerPdct,omitempty"`
}

func (c *CashOption3) SetCreditDebitIndicator(value string) {
	c.CreditDebitIndicator = (*CreditDebitCode)(&value)
}

func (c *CashOption3) SetContractualPaymentIndicator(value string) {
	c.ContractualPaymentIndicator = (*Payment1Code)(&value)
}

func (c *CashOption3) AddNonEligibleProceedsIndicator() *NonEligibleProceedsIndicator1Choice {
	c.NonEligibleProceedsIndicator = new(NonEligibleProceedsIndicator1Choice)
	return c.NonEligibleProceedsIndicator
}

func (c *CashOption3) AddIncomeType() *GenericIdentification20 {
	c.IncomeType = new(GenericIdentification20)
	return c.IncomeType
}

func (c *CashOption3) AddCashAccountIdentification() *CashAccountIdentification5Choice {
	c.CashAccountIdentification = new(CashAccountIdentification5Choice)
	return c.CashAccountIdentification
}

func (c *CashOption3) AddAmountDetails() *CorporateActionAmounts3 {
	c.AmountDetails = new(CorporateActionAmounts3)
	return c.AmountDetails
}

func (c *CashOption3) AddDateDetails() *CorporateActionDate9 {
	c.DateDetails = new(CorporateActionDate9)
	return c.DateDetails
}

func (c *CashOption3) AddForeignExchangeDetails() *ForeignExchangeTerms11 {
	c.ForeignExchangeDetails = new(ForeignExchangeTerms11)
	return c.ForeignExchangeDetails
}

func (c *CashOption3) AddGenericCashPriceReceivedPerProduct() *PriceFormat10Choice {
	c.GenericCashPriceReceivedPerProduct = new(PriceFormat10Choice)
	return c.GenericCashPriceReceivedPerProduct
}
