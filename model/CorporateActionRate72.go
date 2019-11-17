package model

// Specifies rate details.
type CorporateActionRate72 struct {

	// Quantity of additional intermediate securities/new equities awarded for a given quantity of securities derived from subscription.
	AdditionalQuantityForSubscribedResultantSecurities *RatioFormat20Choice `xml:"AddtlQtyForSbcbdRsltntScties,omitempty"`

	// Quantity of additional securities for a given quantity of underlying securities where underlying securities are not exchanged or debited, for example, 1 for 1: 1 new equity credited for every 1 underlying equity = 2 resulting equities.
	AdditionalQuantityForExistingSecurities *RatioFormat20Choice `xml:"AddtlQtyForExstgScties,omitempty"`

	// Quantity of new securities for a given quantity of underlying securities, where the underlying securities will be exchanged or debited, for example, 2 for 1: 2 new equities credited for every 1 underlying equity debited = 2 resulting equities.
	NewToOld *RatioFormat19Choice `xml:"NewToOd,omitempty"`

	// Rate used to calculate the amount of the charges/fees that cannot be categorised.
	ChargesFees *RateAndAmountFormat39Choice `xml:"ChrgsFees,omitempty"`

	// Percentage of fiscal tax to apply.
	FiscalStamp *PercentageRate `xml:"FsclStmp,omitempty"`

	// Rate applicable to the event announced, for example, redemption rate for a redemption event.
	ApplicableRate *PercentageRate `xml:"AplblRate,omitempty"`

	// Amount of money per equity allocated as the result of a tax credit.
	TaxCreditRate []*TaxCreditRateFormat7Choice `xml:"TaxCdtRate,omitempty"`

	// Rate of financial transaction tax.
	FinancialTransactionTaxRate *PercentageRate `xml:"FinTxTaxRate,omitempty"`
}

func (c *CorporateActionRate72) AddAdditionalQuantityForSubscribedResultantSecurities() *RatioFormat20Choice {
	c.AdditionalQuantityForSubscribedResultantSecurities = new(RatioFormat20Choice)
	return c.AdditionalQuantityForSubscribedResultantSecurities
}

func (c *CorporateActionRate72) AddAdditionalQuantityForExistingSecurities() *RatioFormat20Choice {
	c.AdditionalQuantityForExistingSecurities = new(RatioFormat20Choice)
	return c.AdditionalQuantityForExistingSecurities
}

func (c *CorporateActionRate72) AddNewToOld() *RatioFormat19Choice {
	c.NewToOld = new(RatioFormat19Choice)
	return c.NewToOld
}

func (c *CorporateActionRate72) AddChargesFees() *RateAndAmountFormat39Choice {
	c.ChargesFees = new(RateAndAmountFormat39Choice)
	return c.ChargesFees
}

func (c *CorporateActionRate72) SetFiscalStamp(value string) {
	c.FiscalStamp = (*PercentageRate)(&value)
}

func (c *CorporateActionRate72) SetApplicableRate(value string) {
	c.ApplicableRate = (*PercentageRate)(&value)
}

func (c *CorporateActionRate72) AddTaxCreditRate() *TaxCreditRateFormat7Choice {
	newValue := new(TaxCreditRateFormat7Choice)
	c.TaxCreditRate = append(c.TaxCreditRate, newValue)
	return newValue
}

func (c *CorporateActionRate72) SetFinancialTransactionTaxRate(value string) {
	c.FinancialTransactionTaxRate = (*PercentageRate)(&value)
}
