package model

// Specifies rates related to a corporate action option.
type CorporateActionRate83 struct {

	// Rate used for additional tax that cannot be categorised.
	AdditionalTax *RateAndAmountFormat46Choice `xml:"AddtlTax,omitempty"`

	// Cash dividend amount per equity before deductions or allowances have been made.
	GrossDividendRate []*GrossDividendRateFormat25Choice `xml:"GrssDvddRate,omitempty"`

	// Cash dividend amount per equity after deductions or allowances have been made.
	NetDividendRate []*NetDividendRateFormat27Choice `xml:"NetDvddRate,omitempty"`

	// Public index rate applied to the amount paid to adjust it to inflation.
	IndexFactor *RateAndAmountFormat46Choice `xml:"IndxFctr,omitempty"`

	// The actual interest rate used for the payment of the interest for the specified interest period.
	// Usage guideline: It is used to provide the applicable rate for the current payment, after all calculations have been performed, that is, application of period and method of interest computation.
	InterestRateUsedForPayment []*InterestRateUsedForPaymentFormat10Choice `xml:"IntrstRateUsdForPmt,omitempty"`

	// Maximum percentage of shares available through the over subscription privilege, usually a percentage of the basic subscription shares, for example, an account owner subscribing to 100 shares may over subscribe to a maximum of 50 additional shares when the over subscription maximum is 50 percent.
	MaximumAllowedOversubscriptionRate *RateFormat3Choice `xml:"MaxAllwdOvrsbcptRate,omitempty"`

	// Percentage of securities accepted by the offeror/issuer.
	ProrationRate *RateFormat3Choice `xml:"PrratnRate,omitempty"`

	// Percentage of a cash distribution that will be withheld by the tax authorities of the jurisdiction of the issuer, for which a relief at source and/or reclaim may be possible.
	WithholdingTaxRate []*RateAndAmountFormat47Choice `xml:"WhldgTaxRate,omitempty"`

	// Rate at which the income will be withheld by a jurisdiction other than the jurisdiction of the issuer’s country of tax incorporation, for which a relief at source and/or reclaim may be possible. It is levied in complement or offset of the withholding tax rate (TAXR) levied by the jurisdiction of the issuer’s tax domicile.
	SecondLevelTax []*RateAndAmountFormat47Choice `xml:"ScndLvlTax,omitempty"`

	// Amount included in the dividend/NAV that is identified as gains directly or indirectly derived from interest payments, for example, in the context of the EU Savings directive.
	TaxableIncomePerDividendShare []*RateTypeAndAmountAndStatus33 `xml:"TaxblIncmPerDvddShr,omitempty"`

	// Exchange rate (provided by the issuer) between the dividend or interest rate in the paid currency and the declared dividend or interest rate.
	IssuerDeclaredExchangeRate *ForeignExchangeTerms19 `xml:"IssrDclrdXchgRate,omitempty"`

	// Overall tax withheld at source by fund managers prior to considering the tax obligation of each unit holder.
	TaxOnIncome *RateAndAmountFormat46Choice `xml:"TaxOnIncm,omitempty"`
}

func (c *CorporateActionRate83) AddAdditionalTax() *RateAndAmountFormat46Choice {
	c.AdditionalTax = new(RateAndAmountFormat46Choice)
	return c.AdditionalTax
}

func (c *CorporateActionRate83) AddGrossDividendRate() *GrossDividendRateFormat25Choice {
	newValue := new(GrossDividendRateFormat25Choice)
	c.GrossDividendRate = append(c.GrossDividendRate, newValue)
	return newValue
}

func (c *CorporateActionRate83) AddNetDividendRate() *NetDividendRateFormat27Choice {
	newValue := new(NetDividendRateFormat27Choice)
	c.NetDividendRate = append(c.NetDividendRate, newValue)
	return newValue
}

func (c *CorporateActionRate83) AddIndexFactor() *RateAndAmountFormat46Choice {
	c.IndexFactor = new(RateAndAmountFormat46Choice)
	return c.IndexFactor
}

func (c *CorporateActionRate83) AddInterestRateUsedForPayment() *InterestRateUsedForPaymentFormat10Choice {
	newValue := new(InterestRateUsedForPaymentFormat10Choice)
	c.InterestRateUsedForPayment = append(c.InterestRateUsedForPayment, newValue)
	return newValue
}

func (c *CorporateActionRate83) AddMaximumAllowedOversubscriptionRate() *RateFormat3Choice {
	c.MaximumAllowedOversubscriptionRate = new(RateFormat3Choice)
	return c.MaximumAllowedOversubscriptionRate
}

func (c *CorporateActionRate83) AddProrationRate() *RateFormat3Choice {
	c.ProrationRate = new(RateFormat3Choice)
	return c.ProrationRate
}

func (c *CorporateActionRate83) AddWithholdingTaxRate() *RateAndAmountFormat47Choice {
	newValue := new(RateAndAmountFormat47Choice)
	c.WithholdingTaxRate = append(c.WithholdingTaxRate, newValue)
	return newValue
}

func (c *CorporateActionRate83) AddSecondLevelTax() *RateAndAmountFormat47Choice {
	newValue := new(RateAndAmountFormat47Choice)
	c.SecondLevelTax = append(c.SecondLevelTax, newValue)
	return newValue
}

func (c *CorporateActionRate83) AddTaxableIncomePerDividendShare() *RateTypeAndAmountAndStatus33 {
	newValue := new(RateTypeAndAmountAndStatus33)
	c.TaxableIncomePerDividendShare = append(c.TaxableIncomePerDividendShare, newValue)
	return newValue
}

func (c *CorporateActionRate83) AddIssuerDeclaredExchangeRate() *ForeignExchangeTerms19 {
	c.IssuerDeclaredExchangeRate = new(ForeignExchangeTerms19)
	return c.IssuerDeclaredExchangeRate
}

func (c *CorporateActionRate83) AddTaxOnIncome() *RateAndAmountFormat46Choice {
	c.TaxOnIncome = new(RateAndAmountFormat46Choice)
	return c.TaxOnIncome
}
