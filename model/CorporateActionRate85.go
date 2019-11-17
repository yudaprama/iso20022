package model

// Specifies rates.
type CorporateActionRate85 struct {

	// Cash dividend amount per equity before deductions or allowances have been made.
	GrossDividendRate []*GrossDividendRateFormat23Choice `xml:"GrssDvddRate,omitempty"`

	// Cash dividend amount per equity after deductions or allowances have been made.
	NetDividendRate []*NetDividendRateFormat25Choice `xml:"NetDvddRate,omitempty"`

	// Public index rate applied to the amount paid to adjust it to inflation.
	IndexFactor *RateAndAmountFormat43Choice `xml:"IndxFctr,omitempty"`

	// Actual interest rate used for the payment of the interest for the specified interest period.
	InterestRateUsedForPayment []*InterestRateUsedForPaymentFormat9Choice `xml:"IntrstRateUsdForPmt,omitempty"`

	// A maximum percentage of shares available through the over subscription privilege, usually a percentage of the basic subscription shares, for example, an account owner subscribing to 100 shares may over subscribe to a maximum of 50 additional shares when the over subscription maximum is 50 percent.
	MaximumAllowedOversubscriptionRate *PercentageRate `xml:"MaxAllwdOvrsbcptRate,omitempty"`

	// Percentage of securities accepted by the offeror/issuer.
	ProrationRate *PercentageRate `xml:"PrratnRate,omitempty"`

	// Percentage of a cash distribution that will be withheld by the tax authorities of the jurisdiction of the issuer, for which a relief at source and/or reclaim may be possible.
	WithholdingTaxRate []*RateAndAmountFormat45Choice `xml:"WhldgTaxRate,omitempty"`

	// Rate at which the income will be withheld by a jurisdiction other than the jurisdiction of the issuer’s country of tax incorporation, for which a relief at source and/or reclaim may be possible. It is levied in complement or offset of the withholding tax rate (TAXR) levied by the jurisdiction of the issuer’s tax domicile.
	SecondLevelTax []*RateAndAmountFormat45Choice `xml:"ScndLvlTax,omitempty"`

	// Rate used for additional tax that cannot be categorised.
	AdditionalTax *RateAndAmountFormat43Choice `xml:"AddtlTax,omitempty"`

	// Amount included in the dividend/NAV that is identified as gains directly or indirectly derived from interest payments, for example, in the context of the EU Savings directive.
	TaxableIncomePerDividendShare []*RateTypeAndAmountAndStatus33 `xml:"TaxblIncmPerDvddShr,omitempty"`
}

func (c *CorporateActionRate85) AddGrossDividendRate() *GrossDividendRateFormat23Choice {
	newValue := new(GrossDividendRateFormat23Choice)
	c.GrossDividendRate = append(c.GrossDividendRate, newValue)
	return newValue
}

func (c *CorporateActionRate85) AddNetDividendRate() *NetDividendRateFormat25Choice {
	newValue := new(NetDividendRateFormat25Choice)
	c.NetDividendRate = append(c.NetDividendRate, newValue)
	return newValue
}

func (c *CorporateActionRate85) AddIndexFactor() *RateAndAmountFormat43Choice {
	c.IndexFactor = new(RateAndAmountFormat43Choice)
	return c.IndexFactor
}

func (c *CorporateActionRate85) AddInterestRateUsedForPayment() *InterestRateUsedForPaymentFormat9Choice {
	newValue := new(InterestRateUsedForPaymentFormat9Choice)
	c.InterestRateUsedForPayment = append(c.InterestRateUsedForPayment, newValue)
	return newValue
}

func (c *CorporateActionRate85) SetMaximumAllowedOversubscriptionRate(value string) {
	c.MaximumAllowedOversubscriptionRate = (*PercentageRate)(&value)
}

func (c *CorporateActionRate85) SetProrationRate(value string) {
	c.ProrationRate = (*PercentageRate)(&value)
}

func (c *CorporateActionRate85) AddWithholdingTaxRate() *RateAndAmountFormat45Choice {
	newValue := new(RateAndAmountFormat45Choice)
	c.WithholdingTaxRate = append(c.WithholdingTaxRate, newValue)
	return newValue
}

func (c *CorporateActionRate85) AddSecondLevelTax() *RateAndAmountFormat45Choice {
	newValue := new(RateAndAmountFormat45Choice)
	c.SecondLevelTax = append(c.SecondLevelTax, newValue)
	return newValue
}

func (c *CorporateActionRate85) AddAdditionalTax() *RateAndAmountFormat43Choice {
	c.AdditionalTax = new(RateAndAmountFormat43Choice)
	return c.AdditionalTax
}

func (c *CorporateActionRate85) AddTaxableIncomePerDividendShare() *RateTypeAndAmountAndStatus33 {
	newValue := new(RateTypeAndAmountAndStatus33)
	c.TaxableIncomePerDividendShare = append(c.TaxableIncomePerDividendShare, newValue)
	return newValue
}
