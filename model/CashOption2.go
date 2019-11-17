package model

// Provides information about the cash option.
type CashOption2 struct {

	// Indicates whether the value is a debit or credit.
	CreditDebitIndicator *CreditDebitCode `xml:"CdtDbtInd"`

	// Indicates whether the cash payment occurs or will occur in advance of receipt of proceeds from the issuer and based on a contractual agreement established with the account servicer or upon receipt of proceeds from the issuer.
	ContractualPaymentIndicator *Payment1Code `xml:"CtrctlPmtInd,omitempty"`

	// Specifies the type of income.
	// The lists of income type codes to be used, are available on the SMPG website at www.smpg.info.
	IncomeType *GenericIdentification20 `xml:"IncmTp,omitempty"`

	// Choice between a cash account, a charges account or a tax account.
	Account *Account8Choice `xml:"Acct,omitempty"`

	// Provides information about cash parties.
	CashParties *CashParties2 `xml:"CshPties,omitempty"`

	// Provides information about the amounts related to a cash movement.
	AmountDetails *CorporateActionAmounts2 `xml:"AmtDtls"`

	// Provides information about the dates related to a cash movement.
	DateDetails *CorporateActionDate7 `xml:"DtDtls"`

	// Exchange rate between the amount and the resulting amount.
	ForeignExchangeDetails *ForeignExchangeTerms11 `xml:"FXDtls,omitempty"`

	// Generic cash price received per product by the underlying security holder either as a percentage or an amount, for example, redemption price.
	GenericCashPriceReceivedPerProduct *PriceFormat7Choice `xml:"GncCshPricRcvdPerPdct,omitempty"`

	// Provides information about the tax voucher related to a cash movement.
	TaxVoucherDetails *TaxVoucher2 `xml:"TaxVchrDtls,omitempty"`
}

func (c *CashOption2) SetCreditDebitIndicator(value string) {
	c.CreditDebitIndicator = (*CreditDebitCode)(&value)
}

func (c *CashOption2) SetContractualPaymentIndicator(value string) {
	c.ContractualPaymentIndicator = (*Payment1Code)(&value)
}

func (c *CashOption2) AddIncomeType() *GenericIdentification20 {
	c.IncomeType = new(GenericIdentification20)
	return c.IncomeType
}

func (c *CashOption2) AddAccount() *Account8Choice {
	c.Account = new(Account8Choice)
	return c.Account
}

func (c *CashOption2) AddCashParties() *CashParties2 {
	c.CashParties = new(CashParties2)
	return c.CashParties
}

func (c *CashOption2) AddAmountDetails() *CorporateActionAmounts2 {
	c.AmountDetails = new(CorporateActionAmounts2)
	return c.AmountDetails
}

func (c *CashOption2) AddDateDetails() *CorporateActionDate7 {
	c.DateDetails = new(CorporateActionDate7)
	return c.DateDetails
}

func (c *CashOption2) AddForeignExchangeDetails() *ForeignExchangeTerms11 {
	c.ForeignExchangeDetails = new(ForeignExchangeTerms11)
	return c.ForeignExchangeDetails
}

func (c *CashOption2) AddGenericCashPriceReceivedPerProduct() *PriceFormat7Choice {
	c.GenericCashPriceReceivedPerProduct = new(PriceFormat7Choice)
	return c.GenericCashPriceReceivedPerProduct
}

func (c *CashOption2) AddTaxVoucherDetails() *TaxVoucher2 {
	c.TaxVoucherDetails = new(TaxVoucher2)
	return c.TaxVoucherDetails
}
