package model

// Specifies security rate details.
type CorporateActionRate28 struct {

	// Quantity of additional intermediate securities/new equities awarded for a given quantity of securities derived from subscription.
	AdditionalQuantityForSubscribedResultantSecurities *RatioFormat11Choice `xml:"AddtlQtyForSbcbdRsltntScties,omitempty"`

	// Quantity of additional securities for a given quantity of underlying securities where underlying securities are not exchanged or debited, for example, 1 for 1: 1 new equity credited for every 1 underlying equity = 2 resulting equities.
	AdditionalQuantityForExistingSecurities *RatioFormat11Choice `xml:"AddtlQtyForExstgScties,omitempty"`

	// Quantity of new securities for a given quantity of underlying securities, where the underlying securities will be exchanged or debited, for example, 2 for 1: 2 new equities credited for every 1 underlying equity debited = 2 resulting equities.
	NewToOld *RatioFormat12Choice `xml:"NewToOd,omitempty"`

	// Rate used to determine the cash consideration split across outturn settlement transactions that are the result of a transformation of the parent transaction.
	TransformationRate *PercentageRate `xml:"TrfrmatnRate,omitempty"`

	// Rate used to calculate the amount of the charges/fees that cannot be categorised.
	ChargesFees *RateAndAmountFormat14Choice `xml:"ChrgsFees,omitempty"`

	// Percentage of fiscal tax to apply.
	FiscalStamp *RateFormat6Choice `xml:"FsclStmp,omitempty"`

	// Rate applicable to the event announced, for example, redemption rate for a redemption event.
	ApplicableRate *RateFormat6Choice `xml:"AplblRate,omitempty"`

	// Amount of money per equity allocated as the result of a tax credit.
	TaxCreditRate []*TaxCreditRateFormat5Choice `xml:"TaxCdtRate,omitempty"`
}

func (c *CorporateActionRate28) AddAdditionalQuantityForSubscribedResultantSecurities() *RatioFormat11Choice {
	c.AdditionalQuantityForSubscribedResultantSecurities = new(RatioFormat11Choice)
	return c.AdditionalQuantityForSubscribedResultantSecurities
}

func (c *CorporateActionRate28) AddAdditionalQuantityForExistingSecurities() *RatioFormat11Choice {
	c.AdditionalQuantityForExistingSecurities = new(RatioFormat11Choice)
	return c.AdditionalQuantityForExistingSecurities
}

func (c *CorporateActionRate28) AddNewToOld() *RatioFormat12Choice {
	c.NewToOld = new(RatioFormat12Choice)
	return c.NewToOld
}

func (c *CorporateActionRate28) SetTransformationRate(value string) {
	c.TransformationRate = (*PercentageRate)(&value)
}

func (c *CorporateActionRate28) AddChargesFees() *RateAndAmountFormat14Choice {
	c.ChargesFees = new(RateAndAmountFormat14Choice)
	return c.ChargesFees
}

func (c *CorporateActionRate28) AddFiscalStamp() *RateFormat6Choice {
	c.FiscalStamp = new(RateFormat6Choice)
	return c.FiscalStamp
}

func (c *CorporateActionRate28) AddApplicableRate() *RateFormat6Choice {
	c.ApplicableRate = new(RateFormat6Choice)
	return c.ApplicableRate
}

func (c *CorporateActionRate28) AddTaxCreditRate() *TaxCreditRateFormat5Choice {
	newValue := new(TaxCreditRateFormat5Choice)
	c.TaxCreditRate = append(c.TaxCreditRate, newValue)
	return newValue
}
