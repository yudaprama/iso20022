package model

// Provides information about the cash option.
type CashOption52 struct {

	// Indicates whether the value is a debit or credit.
	CreditDebitIndicator *CreditDebitCode `xml:"CdtDbtInd"`

	// Indicates whether the cash payment occurs or will occur in advance of receipt of proceeds from the issuer and based on a contractual agreement established with the account servicer or upon receipt of proceeds from the issuer.
	ContractualPaymentIndicator *Payment1Code `xml:"CtrctlPmtInd,omitempty"`

	// Proceeds are taxable according to the information provided by the issuer / offeror.
	IssuerOfferorTaxabilityIndicator *IssuerOfferorTaxabilityIndicator1Choice `xml:"IssrOfferrTaxbltyInd,omitempty"`

	// Specifies the type of income.
	// The lists of income type codes to be used, are available on the SMPG website at www.smpg.info.
	IncomeType *GenericIdentification30 `xml:"IncmTp,omitempty"`

	// Specifies the basis for the reduced rate of withholding.
	ExemptionType []*GenericIdentification30 `xml:"XmptnTp,omitempty"`

	// Indicates the country from which the income originates.
	CountryOfIncomeSource *CountryCode `xml:"CtryOfIncmSrc,omitempty"`

	// Choice between a cash account, a charges account or a tax account.
	Account *Account8Choice `xml:"Acct,omitempty"`

	// Provides information about cash parties.
	CashParties *CashParties28 `xml:"CshPties,omitempty"`

	// Provides information about the amounts related to a cash movement.
	AmountDetails *CorporateActionAmounts37 `xml:"AmtDtls"`

	// Provides information about the dates related to a cash movement.
	DateDetails *CorporateActionDate24 `xml:"DtDtls"`

	// Exchange rate between the amount and the resulting amount.
	ForeignExchangeDetails *ForeignExchangeTerms23 `xml:"FXDtls,omitempty"`

	// Provides information about the tax voucher related to a cash movement.
	TaxVoucherDetails *TaxVoucher2 `xml:"TaxVchrDtls,omitempty"`

	// Provides information about the corporate action option.
	RateAndAmountDetails *RateDetails27 `xml:"RateAndAmtDtls,omitempty"`

	// Provides information about the prices related to a corporate action option.
	PriceDetails *PriceDetails23 `xml:"PricDtls,omitempty"`
}

func (c *CashOption52) SetCreditDebitIndicator(value string) {
	c.CreditDebitIndicator = (*CreditDebitCode)(&value)
}

func (c *CashOption52) SetContractualPaymentIndicator(value string) {
	c.ContractualPaymentIndicator = (*Payment1Code)(&value)
}

func (c *CashOption52) AddIssuerOfferorTaxabilityIndicator() *IssuerOfferorTaxabilityIndicator1Choice {
	c.IssuerOfferorTaxabilityIndicator = new(IssuerOfferorTaxabilityIndicator1Choice)
	return c.IssuerOfferorTaxabilityIndicator
}

func (c *CashOption52) AddIncomeType() *GenericIdentification30 {
	c.IncomeType = new(GenericIdentification30)
	return c.IncomeType
}

func (c *CashOption52) AddExemptionType() *GenericIdentification30 {
	newValue := new(GenericIdentification30)
	c.ExemptionType = append(c.ExemptionType, newValue)
	return newValue
}

func (c *CashOption52) SetCountryOfIncomeSource(value string) {
	c.CountryOfIncomeSource = (*CountryCode)(&value)
}

func (c *CashOption52) AddAccount() *Account8Choice {
	c.Account = new(Account8Choice)
	return c.Account
}

func (c *CashOption52) AddCashParties() *CashParties28 {
	c.CashParties = new(CashParties28)
	return c.CashParties
}

func (c *CashOption52) AddAmountDetails() *CorporateActionAmounts37 {
	c.AmountDetails = new(CorporateActionAmounts37)
	return c.AmountDetails
}

func (c *CashOption52) AddDateDetails() *CorporateActionDate24 {
	c.DateDetails = new(CorporateActionDate24)
	return c.DateDetails
}

func (c *CashOption52) AddForeignExchangeDetails() *ForeignExchangeTerms23 {
	c.ForeignExchangeDetails = new(ForeignExchangeTerms23)
	return c.ForeignExchangeDetails
}

func (c *CashOption52) AddTaxVoucherDetails() *TaxVoucher2 {
	c.TaxVoucherDetails = new(TaxVoucher2)
	return c.TaxVoucherDetails
}

func (c *CashOption52) AddRateAndAmountDetails() *RateDetails27 {
	c.RateAndAmountDetails = new(RateDetails27)
	return c.RateAndAmountDetails
}

func (c *CashOption52) AddPriceDetails() *PriceDetails23 {
	c.PriceDetails = new(PriceDetails23)
	return c.PriceDetails
}
