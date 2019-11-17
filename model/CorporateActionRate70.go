package model

// Specifies rates.
type CorporateActionRate70 struct {

	// Cash dividend amount per equity before deductions or allowances have been made.
	GrossDividendRate []*GrossDividendRateFormat21Choice `xml:"GrssDvddRate,omitempty"`

	// Cash dividend amount per equity after deductions or allowances have been made.
	NetDividendRate []*NetDividendRateFormat23Choice `xml:"NetDvddRate,omitempty"`

	// Public index rate applied to the amount paid to adjust it to inflation.
	IndexFactor *RateAndAmountFormat39Choice `xml:"IndxFctr,omitempty"`

	// Actual interest rate used for the payment of the interest for the specified interest period.
	InterestRateUsedForPayment []*InterestRateUsedForPaymentFormat7Choice `xml:"IntrstRateUsdForPmt,omitempty"`

	// A maximum percentage of shares available through the over subscription privilege, usually a percentage of the basic subscription shares, for example, an account owner subscribing to 100 shares may over subscribe to a maximum of 50 additional shares when the over subscription maximum is 50 percent.
	MaximumAllowedOversubscriptionRate *PercentageRate `xml:"MaxAllwdOvrsbcptRate,omitempty"`

	// Proportionate allocation used for the offer.
	ProrationRate *PercentageRate `xml:"PrratnRate,omitempty"`

	// Percentage of a cash distribution that will be withheld by the tax authorities of the jurisdiction of the issuer, for which a relief at source and/or reclaim may be possible.
	WithholdingTaxRate []*RateAndAmountFormat40Choice `xml:"WhldgTaxRate,omitempty"`

	// Rate at which the income will be withheld by a jurisdiction other than the jurisdiction of the issuer’s country of tax incorporation, for which a relief at source and/or reclaim may be possible. It is levied in complement or offset of the withholding tax rate (TAXR) levied by the jurisdiction of the issuer’s tax domicile.
	SecondLevelTax []*RateAndAmountFormat40Choice `xml:"ScndLvlTax,omitempty"`

	// Rate used for additional tax that cannot be categorised.
	AdditionalTax *RateAndAmountFormat39Choice `xml:"AddtlTax,omitempty"`

	// Amount included in the dividend/NAV that is identified as gains directly or indirectly derived from interest payments, for example, in the context of the EU Savings directive.
	TaxableIncomePerDividendShare []*RateTypeAndAmountAndStatus26 `xml:"TaxblIncmPerDvddShr,omitempty"`
}

func (c *CorporateActionRate70) AddGrossDividendRate() *GrossDividendRateFormat21Choice {
	newValue := new(GrossDividendRateFormat21Choice)
	c.GrossDividendRate = append(c.GrossDividendRate, newValue)
	return newValue
}

func (c *CorporateActionRate70) AddNetDividendRate() *NetDividendRateFormat23Choice {
	newValue := new(NetDividendRateFormat23Choice)
	c.NetDividendRate = append(c.NetDividendRate, newValue)
	return newValue
}

func (c *CorporateActionRate70) AddIndexFactor() *RateAndAmountFormat39Choice {
	c.IndexFactor = new(RateAndAmountFormat39Choice)
	return c.IndexFactor
}

func (c *CorporateActionRate70) AddInterestRateUsedForPayment() *InterestRateUsedForPaymentFormat7Choice {
	newValue := new(InterestRateUsedForPaymentFormat7Choice)
	c.InterestRateUsedForPayment = append(c.InterestRateUsedForPayment, newValue)
	return newValue
}

func (c *CorporateActionRate70) SetMaximumAllowedOversubscriptionRate(value string) {
	c.MaximumAllowedOversubscriptionRate = (*PercentageRate)(&value)
}

func (c *CorporateActionRate70) SetProrationRate(value string) {
	c.ProrationRate = (*PercentageRate)(&value)
}

func (c *CorporateActionRate70) AddWithholdingTaxRate() *RateAndAmountFormat40Choice {
	newValue := new(RateAndAmountFormat40Choice)
	c.WithholdingTaxRate = append(c.WithholdingTaxRate, newValue)
	return newValue
}

func (c *CorporateActionRate70) AddSecondLevelTax() *RateAndAmountFormat40Choice {
	newValue := new(RateAndAmountFormat40Choice)
	c.SecondLevelTax = append(c.SecondLevelTax, newValue)
	return newValue
}

func (c *CorporateActionRate70) AddAdditionalTax() *RateAndAmountFormat39Choice {
	c.AdditionalTax = new(RateAndAmountFormat39Choice)
	return c.AdditionalTax
}

func (c *CorporateActionRate70) AddTaxableIncomePerDividendShare() *RateTypeAndAmountAndStatus26 {
	newValue := new(RateTypeAndAmountAndStatus26)
	c.TaxableIncomePerDividendShare = append(c.TaxableIncomePerDividendShare, newValue)
	return newValue
}
