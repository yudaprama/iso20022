package model

// Provides information about the cash option.
type CashOption46 struct {

	// Indicates whether the value is a debit or credit.
	CreditDebitIndicator *CreditDebitCode `xml:"CdtDbtInd"`

	// Indicates whether the cash payment occurs or will occur in advance of receipt of proceeds from the issuer and based on a contractual agreement established with the account servicer or upon receipt of proceeds from the issuer.
	ContractualPaymentIndicator *Payment1Code `xml:"CtrctlPmtInd,omitempty"`

	// Proceeds are taxable according to the information provided by the issuer / offeror.
	IssuerOfferorTaxabilityIndicator *GenericIdentification47 `xml:"IssrOfferrTaxbltyInd,omitempty"`

	// Specifies the type of income.
	// The lists of income type codes to be used, are available on the SMPG website at www.smpg.info.
	IncomeType *GenericIdentification47 `xml:"IncmTp,omitempty"`

	// Specifies the basis for the reduced rate of withholding.
	ExemptionType []*GenericIdentification47 `xml:"XmptnTp,omitempty"`

	// Indicates the country from which the income originates.
	CountryOfIncomeSource *CountryCode `xml:"CtryOfIncmSrc,omitempty"`

	// Choice between a cash account, a charges account or a tax account.
	Account *Account9Choice `xml:"Acct,omitempty"`

	// Provides information about cash parties.
	CashParties *CashParties29 `xml:"CshPties,omitempty"`

	// Provides information about the amounts related to a cash movement.
	AmountDetails *CorporateActionAmounts39 `xml:"AmtDtls"`

	// Provides information about the dates related to a cash movement.
	DateDetails *CorporateActionDate24 `xml:"DtDtls"`

	// Exchange rate between the amount and the resulting amount.
	ForeignExchangeDetails *ForeignExchangeTerms27 `xml:"FXDtls,omitempty"`

	// Provides information about the tax voucher related to a cash movement.
	TaxVoucherDetails *TaxVoucher3 `xml:"TaxVchrDtls,omitempty"`

	// Provides information about the corporate action option.
	RateAndAmountDetails *RateDetails24 `xml:"RateAndAmtDtls,omitempty"`

	// Provides information about the prices related to a corporate action option.
	PriceDetails *PriceDetails24 `xml:"PricDtls,omitempty"`
}

func (c *CashOption46) SetCreditDebitIndicator(value string) {
	c.CreditDebitIndicator = (*CreditDebitCode)(&value)
}

func (c *CashOption46) SetContractualPaymentIndicator(value string) {
	c.ContractualPaymentIndicator = (*Payment1Code)(&value)
}

func (c *CashOption46) AddIssuerOfferorTaxabilityIndicator() *GenericIdentification47 {
	c.IssuerOfferorTaxabilityIndicator = new(GenericIdentification47)
	return c.IssuerOfferorTaxabilityIndicator
}

func (c *CashOption46) AddIncomeType() *GenericIdentification47 {
	c.IncomeType = new(GenericIdentification47)
	return c.IncomeType
}

func (c *CashOption46) AddExemptionType() *GenericIdentification47 {
	newValue := new(GenericIdentification47)
	c.ExemptionType = append(c.ExemptionType, newValue)
	return newValue
}

func (c *CashOption46) SetCountryOfIncomeSource(value string) {
	c.CountryOfIncomeSource = (*CountryCode)(&value)
}

func (c *CashOption46) AddAccount() *Account9Choice {
	c.Account = new(Account9Choice)
	return c.Account
}

func (c *CashOption46) AddCashParties() *CashParties29 {
	c.CashParties = new(CashParties29)
	return c.CashParties
}

func (c *CashOption46) AddAmountDetails() *CorporateActionAmounts39 {
	c.AmountDetails = new(CorporateActionAmounts39)
	return c.AmountDetails
}

func (c *CashOption46) AddDateDetails() *CorporateActionDate24 {
	c.DateDetails = new(CorporateActionDate24)
	return c.DateDetails
}

func (c *CashOption46) AddForeignExchangeDetails() *ForeignExchangeTerms27 {
	c.ForeignExchangeDetails = new(ForeignExchangeTerms27)
	return c.ForeignExchangeDetails
}

func (c *CashOption46) AddTaxVoucherDetails() *TaxVoucher3 {
	c.TaxVoucherDetails = new(TaxVoucher3)
	return c.TaxVoucherDetails
}

func (c *CashOption46) AddRateAndAmountDetails() *RateDetails24 {
	c.RateAndAmountDetails = new(RateDetails24)
	return c.RateAndAmountDetails
}

func (c *CashOption46) AddPriceDetails() *PriceDetails24 {
	c.PriceDetails = new(PriceDetails24)
	return c.PriceDetails
}
