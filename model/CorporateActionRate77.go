package model

// Specifies security rate details.
type CorporateActionRate77 struct {

	// Quantity of additional intermediate securities/new equities awarded for a given quantity of securities derived from subscription.
	AdditionalQuantityForSubscribedResultantSecurities *RatioFormat23Choice `xml:"AddtlQtyForSbcbdRsltntScties,omitempty"`

	// Quantity of additional securities for a given quantity of underlying securities where underlying securities are not exchanged or debited, for example, 1 for 1: 1 new equity credited for every 1 underlying equity = 2 resulting equities.
	AdditionalQuantityForExistingSecurities *RatioFormat23Choice `xml:"AddtlQtyForExstgScties,omitempty"`

	// Quantity of new securities for a given quantity of underlying securities, where the underlying securities will be exchanged or debited, for example, 2 for 1: 2 new equities credited for every 1 underlying equity debited = 2 resulting equities.
	NewToOld *RatioFormat24Choice `xml:"NewToOd,omitempty"`

	// Rate used to determine the cash consideration split across outturn settlement transactions that are the result of a transformation of the parent transaction.
	TransformationRate *PercentageRate `xml:"TrfrmatnRate,omitempty"`

	// Rate used to calculate the amount of the charges/fees that cannot be categorised.
	ChargesFees *RateAndAmountFormat46Choice `xml:"ChrgsFees,omitempty"`

	// Percentage of fiscal tax to apply.
	FiscalStamp *RateFormat3Choice `xml:"FsclStmp,omitempty"`

	// Rate applicable to the event announced, for example, redemption rate for a redemption event.
	ApplicableRate *RateFormat3Choice `xml:"AplblRate,omitempty"`

	// Amount of money per equity allocated as the result of a tax credit.
	TaxCreditRate []*TaxCreditRateFormat10Choice `xml:"TaxCdtRate,omitempty"`

	// Rate of financial transaction tax.
	FinancialTransactionTaxRate *RateFormat3Choice `xml:"FinTxTaxRate,omitempty"`
}

func (c *CorporateActionRate77) AddAdditionalQuantityForSubscribedResultantSecurities() *RatioFormat23Choice {
	c.AdditionalQuantityForSubscribedResultantSecurities = new(RatioFormat23Choice)
	return c.AdditionalQuantityForSubscribedResultantSecurities
}

func (c *CorporateActionRate77) AddAdditionalQuantityForExistingSecurities() *RatioFormat23Choice {
	c.AdditionalQuantityForExistingSecurities = new(RatioFormat23Choice)
	return c.AdditionalQuantityForExistingSecurities
}

func (c *CorporateActionRate77) AddNewToOld() *RatioFormat24Choice {
	c.NewToOld = new(RatioFormat24Choice)
	return c.NewToOld
}

func (c *CorporateActionRate77) SetTransformationRate(value string) {
	c.TransformationRate = (*PercentageRate)(&value)
}

func (c *CorporateActionRate77) AddChargesFees() *RateAndAmountFormat46Choice {
	c.ChargesFees = new(RateAndAmountFormat46Choice)
	return c.ChargesFees
}

func (c *CorporateActionRate77) AddFiscalStamp() *RateFormat3Choice {
	c.FiscalStamp = new(RateFormat3Choice)
	return c.FiscalStamp
}

func (c *CorporateActionRate77) AddApplicableRate() *RateFormat3Choice {
	c.ApplicableRate = new(RateFormat3Choice)
	return c.ApplicableRate
}

func (c *CorporateActionRate77) AddTaxCreditRate() *TaxCreditRateFormat10Choice {
	newValue := new(TaxCreditRateFormat10Choice)
	c.TaxCreditRate = append(c.TaxCreditRate, newValue)
	return newValue
}

func (c *CorporateActionRate77) AddFinancialTransactionTaxRate() *RateFormat3Choice {
	c.FinancialTransactionTaxRate = new(RateFormat3Choice)
	return c.FinancialTransactionTaxRate
}
