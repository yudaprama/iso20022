package model

// Specifies rates.
type CorporateActionRate4 struct {

	// Rate used to calculate the amount of the charges/fees that cannot be categorised.
	ChargesFees *RateAndAmountFormat5Choice `xml:"ChrgsFees,omitempty"`

	// Dividend is final.
	FinalDividendRate *ActiveCurrencyAnd13DecimalAmount `xml:"FnlDvddRate,omitempty"`

	// Percentage of fiscal tax to apply.
	FiscalStamp *PercentageRate `xml:"FsclStmp,omitempty"`

	// Rate resulting from a fully franked dividend paid by a company; rate includes tax credit for companies that have made sufficient tax payments during fiscal period.
	FullyFrankedRate *RateAndAmountFormat5Choice `xml:"FullyFrnkdRate,omitempty"`

	// Cash dividend amount per equity before deductions or allowances have been made.
	GrossDividendRate []*GrossDividendRateFormat2Choice `xml:"GrssDvddRate,omitempty"`

	// Rate of the cash premium made available if the securities holder consents or participates to an event, for example, consent fees.
	CashIncentiveRate *PercentageRate `xml:"CshIncntivRate,omitempty"`

	// Public index rate applied to the amount paid to adjust it to inflation.
	IndexFactor *RateAndAmountFormat5Choice `xml:"IndxFctr,omitempty"`

	// Actual interest rate used for the payment of the interest for the specified interest period.
	InterestRateUsedForPayment []*InterestRateUsedForPaymentFormat2Choice `xml:"IntrstRateUsdForPmt,omitempty"`

	// Cash dividend amount per equity after deductions or allowances have been made.
	NetDividendRate []*NetDividendRateFormat2Choice `xml:"NetDvddRate,omitempty"`

	// Rate per share to which a non-resident is entitled.
	NonResidentRate *RateAndAmountFormat5Choice `xml:"NonResdtRate,omitempty"`

	// A maximum percentage of shares available through the over subscription privilege, usually a percentage of the basic subscription shares, for example, an account owner subscribing to 100 shares may over subscribe to a maximum of 50 additional shares when the over subscription maximum is 50 percent.
	MaximumAllowedOversubscriptionRate *PercentageRate `xml:"MaxAllwdOvrsbcptRate,omitempty"`

	// Dividend is provisional.
	ProvisionalDividendRate *ActiveCurrencyAnd13DecimalAmount `xml:"PrvsnlDvddRate,omitempty"`

	// Proportionate allocation used for the offer.
	ProrationRate *PercentageRate `xml:"PrratnRate,omitempty"`

	// Cash rate made available in an offer in order to encourage participation in the offer.
	SolicitationFeeRate *PercentageRate `xml:"SlctnFeeRate,omitempty"`

	// Amount of money per equity allocated as the result of a tax credit.
	TaxCreditRate []*TaxCreditRateFormat2Choice `xml:"TaxCdtRate,omitempty"`

	// Percentage of the gross dividend rate on which tax must be paid .
	TaxRelatedRate []*RateTypeAndAmountAndStatus6 `xml:"TaxRltdRate,omitempty"`

	// Percentage of a cash distribution that will be withheld by a tax authority.
	WithholdingTaxRate *PercentageRate `xml:"WhldgTaxRate,omitempty"`

	// Taxation applied on an amount clearly identified as an income.
	TaxOnIncome *PercentageRate `xml:"TaxOnIncm,omitempty"`

	// Taxation applied on an amount clearly identified as capital profits, capital gains.
	TaxOnProfits *PercentageRate `xml:"TaxOnPrfts,omitempty"`

	// Percentage of cash that was paid in excess of actual tax obligation and was reclaimed.
	TaxReclaimRate *PercentageRate `xml:"TaxRclmRate,omitempty"`

	// Rate at which the income will be withheld by the jurisdiction in which the income was originally paid, for which relief at source and/or reclaim may be possible.
	WithholdingOfForeignTax *RateAndAmountFormat5Choice `xml:"WhldgOfFrgnTax,omitempty"`

	// Rate at which the income will be withheld by the jurisdiction in which the account owner is located, for which relief at source and/or reclaim may be possible.
	WithholdingOfLocalTax *RateAndAmountFormat5Choice `xml:"WhldgOfLclTax,omitempty"`

	// Rate used for additional tax that cannot be categorised.
	AdditionalTax *RateAndAmountFormat5Choice `xml:"AddtlTax,omitempty"`

	// Rate applicable to the event announced, for example, redemption rate for a redemption event.
	ApplicableRate *PercentageRate `xml:"AplblRate,omitempty"`
}

func (c *CorporateActionRate4) AddChargesFees() *RateAndAmountFormat5Choice {
	c.ChargesFees = new(RateAndAmountFormat5Choice)
	return c.ChargesFees
}

func (c *CorporateActionRate4) SetFinalDividendRate(value, currency string) {
	c.FinalDividendRate = NewActiveCurrencyAnd13DecimalAmount(value, currency)
}

func (c *CorporateActionRate4) SetFiscalStamp(value string) {
	c.FiscalStamp = (*PercentageRate)(&value)
}

func (c *CorporateActionRate4) AddFullyFrankedRate() *RateAndAmountFormat5Choice {
	c.FullyFrankedRate = new(RateAndAmountFormat5Choice)
	return c.FullyFrankedRate
}

func (c *CorporateActionRate4) AddGrossDividendRate() *GrossDividendRateFormat2Choice {
	newValue := new(GrossDividendRateFormat2Choice)
	c.GrossDividendRate = append(c.GrossDividendRate, newValue)
	return newValue
}

func (c *CorporateActionRate4) SetCashIncentiveRate(value string) {
	c.CashIncentiveRate = (*PercentageRate)(&value)
}

func (c *CorporateActionRate4) AddIndexFactor() *RateAndAmountFormat5Choice {
	c.IndexFactor = new(RateAndAmountFormat5Choice)
	return c.IndexFactor
}

func (c *CorporateActionRate4) AddInterestRateUsedForPayment() *InterestRateUsedForPaymentFormat2Choice {
	newValue := new(InterestRateUsedForPaymentFormat2Choice)
	c.InterestRateUsedForPayment = append(c.InterestRateUsedForPayment, newValue)
	return newValue
}

func (c *CorporateActionRate4) AddNetDividendRate() *NetDividendRateFormat2Choice {
	newValue := new(NetDividendRateFormat2Choice)
	c.NetDividendRate = append(c.NetDividendRate, newValue)
	return newValue
}

func (c *CorporateActionRate4) AddNonResidentRate() *RateAndAmountFormat5Choice {
	c.NonResidentRate = new(RateAndAmountFormat5Choice)
	return c.NonResidentRate
}

func (c *CorporateActionRate4) SetMaximumAllowedOversubscriptionRate(value string) {
	c.MaximumAllowedOversubscriptionRate = (*PercentageRate)(&value)
}

func (c *CorporateActionRate4) SetProvisionalDividendRate(value, currency string) {
	c.ProvisionalDividendRate = NewActiveCurrencyAnd13DecimalAmount(value, currency)
}

func (c *CorporateActionRate4) SetProrationRate(value string) {
	c.ProrationRate = (*PercentageRate)(&value)
}

func (c *CorporateActionRate4) SetSolicitationFeeRate(value string) {
	c.SolicitationFeeRate = (*PercentageRate)(&value)
}

func (c *CorporateActionRate4) AddTaxCreditRate() *TaxCreditRateFormat2Choice {
	newValue := new(TaxCreditRateFormat2Choice)
	c.TaxCreditRate = append(c.TaxCreditRate, newValue)
	return newValue
}

func (c *CorporateActionRate4) AddTaxRelatedRate() *RateTypeAndAmountAndStatus6 {
	newValue := new(RateTypeAndAmountAndStatus6)
	c.TaxRelatedRate = append(c.TaxRelatedRate, newValue)
	return newValue
}

func (c *CorporateActionRate4) SetWithholdingTaxRate(value string) {
	c.WithholdingTaxRate = (*PercentageRate)(&value)
}

func (c *CorporateActionRate4) SetTaxOnIncome(value string) {
	c.TaxOnIncome = (*PercentageRate)(&value)
}

func (c *CorporateActionRate4) SetTaxOnProfits(value string) {
	c.TaxOnProfits = (*PercentageRate)(&value)
}

func (c *CorporateActionRate4) SetTaxReclaimRate(value string) {
	c.TaxReclaimRate = (*PercentageRate)(&value)
}

func (c *CorporateActionRate4) AddWithholdingOfForeignTax() *RateAndAmountFormat5Choice {
	c.WithholdingOfForeignTax = new(RateAndAmountFormat5Choice)
	return c.WithholdingOfForeignTax
}

func (c *CorporateActionRate4) AddWithholdingOfLocalTax() *RateAndAmountFormat5Choice {
	c.WithholdingOfLocalTax = new(RateAndAmountFormat5Choice)
	return c.WithholdingOfLocalTax
}

func (c *CorporateActionRate4) AddAdditionalTax() *RateAndAmountFormat5Choice {
	c.AdditionalTax = new(RateAndAmountFormat5Choice)
	return c.AdditionalTax
}

func (c *CorporateActionRate4) SetApplicableRate(value string) {
	c.ApplicableRate = (*PercentageRate)(&value)
}
