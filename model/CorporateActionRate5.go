package model

// Specifies rates related to a corporate action option.
type CorporateActionRate5 struct {

	// Rate used for additional tax that cannot be categorised.
	AdditionalTax *RateAndAmountFormat3Choice `xml:"AddtlTax,omitempty"`

	// Rate used to calculate the amount of the charges/fees that cannot be categorised.
	ChargesFees *RateAndAmountFormat3Choice `xml:"ChrgsFees,omitempty"`

	// Dividend is final.
	FinalDividendRate *RateAndAmountFormat4Choice `xml:"FnlDvddRate,omitempty"`

	// Percentage of fiscal tax to apply.
	FiscalStamp *RateFormat2Choice `xml:"FsclStmp,omitempty"`

	// Rate resulting from a fully franked dividend paid by a company; rate includes tax credit for companies that have made sufficient tax payments during fiscal period.
	FullyFrankedRate *RateAndAmountFormat3Choice `xml:"FullyFrnkdRate,omitempty"`

	// Cash dividend amount per equity before deductions or allowances have been made.
	GrossDividendRate []*GrossDividendRateFormat1Choice `xml:"GrssDvddRate,omitempty"`

	// Rate of the cash premium made available if the securities holder consents or participates to an event, for example, consent fees.
	CashIncentiveRate *RateFormat2Choice `xml:"CshIncntivRate,omitempty"`

	// Public index rate applied to the amount paid to adjust it to inflation.
	IndexFactor *RateAndAmountFormat3Choice `xml:"IndxFctr,omitempty"`

	// The actual interest rate used for the payment of the interest for the specified interest period.
	// Usage guideline: It is used to provide the applicable rate for the current payment, after all calculations have been performed, that is, application of period and method of interest computation.
	InterestRateUsedForPayment []*InterestRateUsedForPaymentFormat1Choice `xml:"IntrstRateUsdForPmt,omitempty"`

	// Cash dividend amount per equity after deductions or allowances have been made.
	NetDividendRate []*NetDividendRateFormat1Choice `xml:"NetDvddRate,omitempty"`

	// Rate per share to which a non-resident is entitled.
	NonResidentRate *RateAndAmountFormat3Choice `xml:"NonResdtRate,omitempty"`

	// Maximum percentage of shares available through the over subscription privilege, usually a percentage of the basic subscription shares, for example, an account owner subscribing to 100 shares may over subscribe to a maximum of 50 additional shares when the over subscription maximum is 50 percent.
	MaximumAllowedOversubscriptionRate *RateFormat2Choice `xml:"MaxAllwdOvrsbcptRate,omitempty"`

	// Dividend is provisional.
	ProvisionalDividendRate *RateAndAmountFormat4Choice `xml:"PrvsnlDvddRate,omitempty"`

	// Amount of money per equity allocated as the result of a tax credit.
	TaxCreditRate []*TaxCreditRateFormat1Choice `xml:"TaxCdtRate,omitempty"`

	// Proportionate allocation used for the offer.
	ProrationRate *RateFormat2Choice `xml:"PrratnRate,omitempty"`

	// Cash rate made available in an offer in order to encourage participation in the offer.
	SolicitationFeeRate *SolicitationFeeRateFormat1Choice `xml:"SlctnFeeRate,omitempty"`

	// Cash rate made available, as an incentive, in addition to the solicitation fee, in order to encourage early participation in an offer.
	EarlySolicitationFeeRate *SolicitationFeeRateFormat1Choice `xml:"EarlySlctnFeeRate,omitempty"`

	// Percentage of a cash distribution that will be withheld by a tax authority.
	WithholdingTaxRate *RateFormat2Choice `xml:"WhldgTaxRate,omitempty"`

	// Taxation applied on an amount clearly identified as an income.
	TaxOnIncome *RateFormat2Choice `xml:"TaxOnIncm,omitempty"`

	// Taxation applied on an amount clearly identified as capital profits, capital gains.
	TaxOnProfits *RateFormat2Choice `xml:"TaxOnPrfts,omitempty"`

	// Percentage of cash that was paid in excess of actual tax obligation and was reclaimed.
	TaxReclaimRate *RateFormat2Choice `xml:"TaxRclmRate,omitempty"`

	// Rate at which the income will be withheld by the jurisdiction in which the income was originally paid, for which relief at source and/or reclaim may be possible.
	WithholdingOfForeignTax *RateAndAmountFormat3Choice `xml:"WhldgOfFrgnTax,omitempty"`

	// Rate at which the income will be withheld by the jurisdiction in which the account owner is located, for which relief at source and/or reclaim may be possible.
	WithholdingOfLocalTax *RateAndAmountFormat3Choice `xml:"WhldgOfLclTax,omitempty"`

	// Percentage of the gross dividend rate on which tax must be paid .
	TaxRelatedRate []*RateTypeAndAmountAndStatus6 `xml:"TaxRltdRate,omitempty"`

	// Rate applicable to the event announced, for example, redemption rate for a redemption event.
	ApplicableRate *RateFormat2Choice `xml:"AplblRate,omitempty"`
}

func (c *CorporateActionRate5) AddAdditionalTax() *RateAndAmountFormat3Choice {
	c.AdditionalTax = new(RateAndAmountFormat3Choice)
	return c.AdditionalTax
}

func (c *CorporateActionRate5) AddChargesFees() *RateAndAmountFormat3Choice {
	c.ChargesFees = new(RateAndAmountFormat3Choice)
	return c.ChargesFees
}

func (c *CorporateActionRate5) AddFinalDividendRate() *RateAndAmountFormat4Choice {
	c.FinalDividendRate = new(RateAndAmountFormat4Choice)
	return c.FinalDividendRate
}

func (c *CorporateActionRate5) AddFiscalStamp() *RateFormat2Choice {
	c.FiscalStamp = new(RateFormat2Choice)
	return c.FiscalStamp
}

func (c *CorporateActionRate5) AddFullyFrankedRate() *RateAndAmountFormat3Choice {
	c.FullyFrankedRate = new(RateAndAmountFormat3Choice)
	return c.FullyFrankedRate
}

func (c *CorporateActionRate5) AddGrossDividendRate() *GrossDividendRateFormat1Choice {
	newValue := new(GrossDividendRateFormat1Choice)
	c.GrossDividendRate = append(c.GrossDividendRate, newValue)
	return newValue
}

func (c *CorporateActionRate5) AddCashIncentiveRate() *RateFormat2Choice {
	c.CashIncentiveRate = new(RateFormat2Choice)
	return c.CashIncentiveRate
}

func (c *CorporateActionRate5) AddIndexFactor() *RateAndAmountFormat3Choice {
	c.IndexFactor = new(RateAndAmountFormat3Choice)
	return c.IndexFactor
}

func (c *CorporateActionRate5) AddInterestRateUsedForPayment() *InterestRateUsedForPaymentFormat1Choice {
	newValue := new(InterestRateUsedForPaymentFormat1Choice)
	c.InterestRateUsedForPayment = append(c.InterestRateUsedForPayment, newValue)
	return newValue
}

func (c *CorporateActionRate5) AddNetDividendRate() *NetDividendRateFormat1Choice {
	newValue := new(NetDividendRateFormat1Choice)
	c.NetDividendRate = append(c.NetDividendRate, newValue)
	return newValue
}

func (c *CorporateActionRate5) AddNonResidentRate() *RateAndAmountFormat3Choice {
	c.NonResidentRate = new(RateAndAmountFormat3Choice)
	return c.NonResidentRate
}

func (c *CorporateActionRate5) AddMaximumAllowedOversubscriptionRate() *RateFormat2Choice {
	c.MaximumAllowedOversubscriptionRate = new(RateFormat2Choice)
	return c.MaximumAllowedOversubscriptionRate
}

func (c *CorporateActionRate5) AddProvisionalDividendRate() *RateAndAmountFormat4Choice {
	c.ProvisionalDividendRate = new(RateAndAmountFormat4Choice)
	return c.ProvisionalDividendRate
}

func (c *CorporateActionRate5) AddTaxCreditRate() *TaxCreditRateFormat1Choice {
	newValue := new(TaxCreditRateFormat1Choice)
	c.TaxCreditRate = append(c.TaxCreditRate, newValue)
	return newValue
}

func (c *CorporateActionRate5) AddProrationRate() *RateFormat2Choice {
	c.ProrationRate = new(RateFormat2Choice)
	return c.ProrationRate
}

func (c *CorporateActionRate5) AddSolicitationFeeRate() *SolicitationFeeRateFormat1Choice {
	c.SolicitationFeeRate = new(SolicitationFeeRateFormat1Choice)
	return c.SolicitationFeeRate
}

func (c *CorporateActionRate5) AddEarlySolicitationFeeRate() *SolicitationFeeRateFormat1Choice {
	c.EarlySolicitationFeeRate = new(SolicitationFeeRateFormat1Choice)
	return c.EarlySolicitationFeeRate
}

func (c *CorporateActionRate5) AddWithholdingTaxRate() *RateFormat2Choice {
	c.WithholdingTaxRate = new(RateFormat2Choice)
	return c.WithholdingTaxRate
}

func (c *CorporateActionRate5) AddTaxOnIncome() *RateFormat2Choice {
	c.TaxOnIncome = new(RateFormat2Choice)
	return c.TaxOnIncome
}

func (c *CorporateActionRate5) AddTaxOnProfits() *RateFormat2Choice {
	c.TaxOnProfits = new(RateFormat2Choice)
	return c.TaxOnProfits
}

func (c *CorporateActionRate5) AddTaxReclaimRate() *RateFormat2Choice {
	c.TaxReclaimRate = new(RateFormat2Choice)
	return c.TaxReclaimRate
}

func (c *CorporateActionRate5) AddWithholdingOfForeignTax() *RateAndAmountFormat3Choice {
	c.WithholdingOfForeignTax = new(RateAndAmountFormat3Choice)
	return c.WithholdingOfForeignTax
}

func (c *CorporateActionRate5) AddWithholdingOfLocalTax() *RateAndAmountFormat3Choice {
	c.WithholdingOfLocalTax = new(RateAndAmountFormat3Choice)
	return c.WithholdingOfLocalTax
}

func (c *CorporateActionRate5) AddTaxRelatedRate() *RateTypeAndAmountAndStatus6 {
	newValue := new(RateTypeAndAmountAndStatus6)
	c.TaxRelatedRate = append(c.TaxRelatedRate, newValue)
	return newValue
}

func (c *CorporateActionRate5) AddApplicableRate() *RateFormat2Choice {
	c.ApplicableRate = new(RateFormat2Choice)
	return c.ApplicableRate
}
