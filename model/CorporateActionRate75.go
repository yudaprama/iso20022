package model

// Specifies rate details.
type CorporateActionRate75 struct {

	// Quantity of additional intermediate securities/new equities awarded for a given quantity of securities derived from subscription.
	AdditionalQuantityForSubscribedResultantSecurities *RatioFormat21Choice `xml:"AddtlQtyForSbcbdRsltntScties,omitempty"`

	// Quantity of additional securities for a given quantity of underlying securities where underlying securities are not exchanged or debited, for example, 1 for 1: 1 new equity credited for every 1 underlying equity = 2 resulting equities.
	AdditionalQuantityForExistingSecurities *RatioFormat21Choice `xml:"AddtlQtyForExstgScties,omitempty"`

	// Quantity of new securities for a given quantity of underlying securities, where the underlying securities will be exchanged or debited, for example, 2 for 1: 2 new equities credited for every 1 underlying equity debited = 2 resulting equities.
	NewToOld *RatioFormat22Choice `xml:"NewToOd,omitempty"`

	// Rate used to calculate the amount of the charges/fees that cannot be categorised.
	ChargesFees *RateAndAmountFormat43Choice `xml:"ChrgsFees,omitempty"`

	// Percentage of fiscal tax to apply.
	FiscalStamp *PercentageRate `xml:"FsclStmp,omitempty"`

	// Rate applicable to the event announced, for example, redemption rate for a redemption event.
	ApplicableRate *PercentageRate `xml:"AplblRate,omitempty"`

	// Amount of money per equity allocated as the result of a tax credit.
	TaxCreditRate []*TaxCreditRateFormat9Choice `xml:"TaxCdtRate,omitempty"`

	// Rate of financial transaction tax.
	FinancialTransactionTaxRate *PercentageRate `xml:"FinTxTaxRate,omitempty"`
}

func (c *CorporateActionRate75) AddAdditionalQuantityForSubscribedResultantSecurities() *RatioFormat21Choice {
	c.AdditionalQuantityForSubscribedResultantSecurities = new(RatioFormat21Choice)
	return c.AdditionalQuantityForSubscribedResultantSecurities
}

func (c *CorporateActionRate75) AddAdditionalQuantityForExistingSecurities() *RatioFormat21Choice {
	c.AdditionalQuantityForExistingSecurities = new(RatioFormat21Choice)
	return c.AdditionalQuantityForExistingSecurities
}

func (c *CorporateActionRate75) AddNewToOld() *RatioFormat22Choice {
	c.NewToOld = new(RatioFormat22Choice)
	return c.NewToOld
}

func (c *CorporateActionRate75) AddChargesFees() *RateAndAmountFormat43Choice {
	c.ChargesFees = new(RateAndAmountFormat43Choice)
	return c.ChargesFees
}

func (c *CorporateActionRate75) SetFiscalStamp(value string) {
	c.FiscalStamp = (*PercentageRate)(&value)
}

func (c *CorporateActionRate75) SetApplicableRate(value string) {
	c.ApplicableRate = (*PercentageRate)(&value)
}

func (c *CorporateActionRate75) AddTaxCreditRate() *TaxCreditRateFormat9Choice {
	newValue := new(TaxCreditRateFormat9Choice)
	c.TaxCreditRate = append(c.TaxCreditRate, newValue)
	return newValue
}

func (c *CorporateActionRate75) SetFinancialTransactionTaxRate(value string) {
	c.FinancialTransactionTaxRate = (*PercentageRate)(&value)
}
