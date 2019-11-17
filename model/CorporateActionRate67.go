package model

// Specifies rates related to a corporate action option.
type CorporateActionRate67 struct {

	// Rate used for additional tax that cannot be categorised.
	AdditionalTax *RateAndAmountFormat37Choice `xml:"AddtlTax,omitempty"`

	// Cash dividend amount per equity before deductions or allowances have been made.
	GrossDividendRate []*GrossDividendRateFormat19Choice `xml:"GrssDvddRate,omitempty"`

	// Cash dividend amount per equity after deductions or allowances have been made.
	NetDividendRate []*NetDividendRateFormat21Choice `xml:"NetDvddRate,omitempty"`

	// Public index rate applied to the amount paid to adjust it to inflation.
	IndexFactor *RateAndAmountFormat37Choice `xml:"IndxFctr,omitempty"`

	// The actual interest rate used for the payment of the interest for the specified interest period.
	// Usage guideline: It is used to provide the applicable rate for the current payment, after all calculations have been performed, that is, application of period and method of interest computation.
	InterestRateUsedForPayment []*InterestRateUsedForPaymentFormat8Choice `xml:"IntrstRateUsdForPmt,omitempty"`

	// Maximum percentage of shares available through the over subscription privilege, usually a percentage of the basic subscription shares, for example, an account owner subscribing to 100 shares may over subscribe to a maximum of 50 additional shares when the over subscription maximum is 50 percent.
	MaximumAllowedOversubscriptionRate *RateFormat3Choice `xml:"MaxAllwdOvrsbcptRate,omitempty"`

	// Proportionate allocation used for the offer.
	ProrationRate *RateFormat3Choice `xml:"PrratnRate,omitempty"`

	// Percentage of a cash distribution that will be withheld by the tax authorities of the jurisdiction of the issuer, for which a relief at source and/or reclaim may be possible.
	WithholdingTaxRate []*RateAndAmountFormat41Choice `xml:"WhldgTaxRate,omitempty"`

	// Rate at which the income will be withheld by a jurisdiction other than the jurisdiction of the issuer’s country of tax incorporation, for which a relief at source and/or reclaim may be possible. It is levied in complement or offset of the withholding tax rate (TAXR) levied by the jurisdiction of the issuer’s tax domicile.
	SecondLevelTax []*RateAndAmountFormat41Choice `xml:"ScndLvlTax,omitempty"`

	// Amount included in the dividend/NAV that is identified as gains directly or indirectly derived from interest payments, for example, in the context of the EU Savings directive.
	TaxableIncomePerDividendShare []*RateTypeAndAmountAndStatus26 `xml:"TaxblIncmPerDvddShr,omitempty"`

	// Overall tax withheld at source by fund managers prior to considering the tax obligation of each unit holder.
	TaxOnIncome *RateAndAmountFormat37Choice `xml:"TaxOnIncm,omitempty"`
}

func (c *CorporateActionRate67) AddAdditionalTax() *RateAndAmountFormat37Choice {
	c.AdditionalTax = new(RateAndAmountFormat37Choice)
	return c.AdditionalTax
}

func (c *CorporateActionRate67) AddGrossDividendRate() *GrossDividendRateFormat19Choice {
	newValue := new(GrossDividendRateFormat19Choice)
	c.GrossDividendRate = append(c.GrossDividendRate, newValue)
	return newValue
}

func (c *CorporateActionRate67) AddNetDividendRate() *NetDividendRateFormat21Choice {
	newValue := new(NetDividendRateFormat21Choice)
	c.NetDividendRate = append(c.NetDividendRate, newValue)
	return newValue
}

func (c *CorporateActionRate67) AddIndexFactor() *RateAndAmountFormat37Choice {
	c.IndexFactor = new(RateAndAmountFormat37Choice)
	return c.IndexFactor
}

func (c *CorporateActionRate67) AddInterestRateUsedForPayment() *InterestRateUsedForPaymentFormat8Choice {
	newValue := new(InterestRateUsedForPaymentFormat8Choice)
	c.InterestRateUsedForPayment = append(c.InterestRateUsedForPayment, newValue)
	return newValue
}

func (c *CorporateActionRate67) AddMaximumAllowedOversubscriptionRate() *RateFormat3Choice {
	c.MaximumAllowedOversubscriptionRate = new(RateFormat3Choice)
	return c.MaximumAllowedOversubscriptionRate
}

func (c *CorporateActionRate67) AddProrationRate() *RateFormat3Choice {
	c.ProrationRate = new(RateFormat3Choice)
	return c.ProrationRate
}

func (c *CorporateActionRate67) AddWithholdingTaxRate() *RateAndAmountFormat41Choice {
	newValue := new(RateAndAmountFormat41Choice)
	c.WithholdingTaxRate = append(c.WithholdingTaxRate, newValue)
	return newValue
}

func (c *CorporateActionRate67) AddSecondLevelTax() *RateAndAmountFormat41Choice {
	newValue := new(RateAndAmountFormat41Choice)
	c.SecondLevelTax = append(c.SecondLevelTax, newValue)
	return newValue
}

func (c *CorporateActionRate67) AddTaxableIncomePerDividendShare() *RateTypeAndAmountAndStatus26 {
	newValue := new(RateTypeAndAmountAndStatus26)
	c.TaxableIncomePerDividendShare = append(c.TaxableIncomePerDividendShare, newValue)
	return newValue
}

func (c *CorporateActionRate67) AddTaxOnIncome() *RateAndAmountFormat37Choice {
	c.TaxOnIncome = new(RateAndAmountFormat37Choice)
	return c.TaxOnIncome
}
