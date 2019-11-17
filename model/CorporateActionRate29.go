package model

// Specifies rate details.
type CorporateActionRate29 struct {

	// Quantity of additional intermediate securities/new equities awarded for a given quantity of securities derived from subscription.
	AdditionalQuantityForSubscribedResultantSecurities *RatioFormat3Choice `xml:"AddtlQtyForSbcbdRsltntScties,omitempty"`

	// Quantity of additional securities for a given quantity of underlying securities where underlying securities are not exchanged or debited, for example, 1 for 1: 1 new equity credited for every 1 underlying equity = 2 resulting equities.
	AdditionalQuantityForExistingSecurities *RatioFormat3Choice `xml:"AddtlQtyForExstgScties,omitempty"`

	// Quantity of new securities for a given quantity of underlying securities, where the underlying securities will be exchanged or debited, for example, 2 for 1: 2 new equities credited for every 1 underlying equity debited = 2 resulting equities.
	NewToOld *RatioFormat15Choice `xml:"NewToOd,omitempty"`

	// Rate used to calculate the amount of the charges/fees that cannot be categorised.
	ChargesFees *RateAndAmountFormat5Choice `xml:"ChrgsFees,omitempty"`

	// Percentage of fiscal tax to apply.
	FiscalStamp *PercentageRate `xml:"FsclStmp,omitempty"`

	// Rate applicable to the event announced, for example, redemption rate for a redemption event.
	ApplicableRate *PercentageRate `xml:"AplblRate,omitempty"`

	// Amount of money per equity allocated as the result of a tax credit.
	TaxCreditRate []*TaxCreditRateFormat2Choice `xml:"TaxCdtRate,omitempty"`
}

func (c *CorporateActionRate29) AddAdditionalQuantityForSubscribedResultantSecurities() *RatioFormat3Choice {
	c.AdditionalQuantityForSubscribedResultantSecurities = new(RatioFormat3Choice)
	return c.AdditionalQuantityForSubscribedResultantSecurities
}

func (c *CorporateActionRate29) AddAdditionalQuantityForExistingSecurities() *RatioFormat3Choice {
	c.AdditionalQuantityForExistingSecurities = new(RatioFormat3Choice)
	return c.AdditionalQuantityForExistingSecurities
}

func (c *CorporateActionRate29) AddNewToOld() *RatioFormat15Choice {
	c.NewToOld = new(RatioFormat15Choice)
	return c.NewToOld
}

func (c *CorporateActionRate29) AddChargesFees() *RateAndAmountFormat5Choice {
	c.ChargesFees = new(RateAndAmountFormat5Choice)
	return c.ChargesFees
}

func (c *CorporateActionRate29) SetFiscalStamp(value string) {
	c.FiscalStamp = (*PercentageRate)(&value)
}

func (c *CorporateActionRate29) SetApplicableRate(value string) {
	c.ApplicableRate = (*PercentageRate)(&value)
}

func (c *CorporateActionRate29) AddTaxCreditRate() *TaxCreditRateFormat2Choice {
	newValue := new(TaxCreditRateFormat2Choice)
	c.TaxCreditRate = append(c.TaxCreditRate, newValue)
	return newValue
}
