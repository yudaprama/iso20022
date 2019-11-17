package model

// Specifies rates.
type CorporateActionRate2 struct {

	// Percentage of a cash distribution that will be withheld by a tax authority.
	WithholdingTax *RateFormat1Choice `xml:"WhldgTax,omitempty"`

	// Rate at which the income will be withheld by the jurisdiction in which the income was originally paid, for which relief at source and/or reclaim may be possible.
	WithholdingOfForeignTax *RateAndAmountFormat1Choice `xml:"WhldgOfFrgnTax,omitempty"`

	// Rate at which the income will be withheld by the jurisdiction in which the account owner is located, for which relief at source and/or reclaim may be possible.
	WithholdingOfLocalTax *RateAndAmountFormat1Choice `xml:"WhldgOfLclTax,omitempty"`

	// Local tax (ZAS Anrechnungsbetrag) subject to interest down payment tax (proportion of interest liable for interest down payment tax/interim profit that is not covered by the tax exempt amount).
	GermanLocalTax1 *RateAndAmountFormat1Choice `xml:"GrmnLclTax1,omitempty"`

	// Local tax (ZAS Pflichtige Zinsen) interest liable for interest down payment tax (proportion of gross interest per unit/interim profits that is not covered by the credit in the interest pool).
	GermanLocalTax2 *RateAndAmountFormat1Choice `xml:"GrmnLclTax2,omitempty"`

	// Local tax (Zinstopf) offset interest per unit against tax exempt amount (variation to offset interest per unit in relation to tax exempt amount).
	GermanLocalTax3 *RateAndAmountFormat1Choice `xml:"GrmnLclTax3,omitempty"`

	// Local tax (Ertrag Besitzanteilig) yield liable for interest down payment tax.
	GermanLocalTax4 *RateAndAmountFormat1Choice `xml:"GrmnLclTax4,omitempty"`

	// Taxation applied on an amount clearly identified as an income.
	TaxOnIncome *RateFormat1Choice `xml:"TaxOnIncm,omitempty"`

	// Taxation applied on an amount clearly identified as capital profits, capital gains.
	TaxOnProfit *RateFormat1Choice `xml:"TaxOnPrft,omitempty"`

	// Percentage of cash that was paid in excess of actual tax obligation and was reclaimed.
	TaxReclaim *RateFormat1Choice `xml:"TaxRclm,omitempty"`

	// Percentage of fiscal tax to apply.
	FiscalStamp *RateFormat1Choice `xml:"FsclStmp,omitempty"`

	// Proportionate allocation used for the offer.
	Proration *RateFormat1Choice `xml:"Prratn,omitempty"`

	// Quantity of new securities for a given quantity of underlying securities, where the underlying securities will be exchanged or debited, eg, 2 for 1: 2 new equities credited for every 1 underlying equity debited = 2 resulting equities.
	NewToOld *RatioFormat2Choice `xml:"NewToOd,omitempty"`

	// Quantity of new equities that will be derived by the exercise of a given quantity of intermediate securities.
	NewSecuritiesToUnderlyingSecurities *RatioFormat2Choice `xml:"NewSctiesToUndrlygScties,omitempty"`

	// Quantity of additional securities for a given quantity of underlying securities where underlying securities are not exchanged or debited, eg, 1 for 1: 1 new
	// equity credited for every 1 underlying equity = 2 resulting equities.
	AdditionalQuantityForExistingSecurities *RatioFormat1Choice `xml:"AddtlQtyForExstgScties,omitempty"`

	// Quantity of additional intermediate securities/new equities awarded for a given quantity of securities derived from subscription.
	AdditionalQuantityForSubscribedResultantSecurities *RatioFormat1Choice `xml:"AddtlQtyForSbcbdRsltntScties,omitempty"`

	// Percentage of the gross dividend rate on which tax must be paid .
	RelatedTax *RelatedTaxType1 `xml:"RltdTax,omitempty"`

	// Rate per share to which a non-resident is entitled.
	NonResidentRate *RateAndAmountFormat1Choice `xml:"NonResdtRate,omitempty"`

	// Rate used to calculate the amount of the charges/fees that cannot be categorised.
	Charges *RateAndAmountFormat1Choice `xml:"Chrgs,omitempty"`

	// The actual interest rate used for the payment of the interest for the specified interest period.
	InterestForUsedPayment *RateAndAmountFormat1Choice `xml:"IntrstForUsdPmt,omitempty"`

	// Public index rate applied to the amount paid to adjust it to inflation.
	IndexFactor *RateAndAmountFormat1Choice `xml:"IndxFctr,omitempty"`

	// Rate resulting from a fully franked dividend paid by a company; rate includes tax credit for companies that have made sufficient tax payments during fiscal period.
	FullyFranked *RateAndAmountFormat1Choice `xml:"FullyFrnkd,omitempty"`

	// Cash dividend amount per equity before deductions or allowances have been made.
	GrossDividend *GrossDividendRate1Choice `xml:"GrssDvdd,omitempty"`

	// Cash dividend amount per equity after deductions or allowances have been made.
	NetDividend *NetDividendRate1Choice `xml:"NetDvdd,omitempty"`

	// Dividend is final.
	FinalDividend *AmountAndRateFormat2Choice `xml:"FnlDvdd,omitempty"`

	// Dividend is provisional.
	ProvisionalDividend *AmountAndRateFormat2Choice `xml:"PrvsnlDvdd,omitempty"`

	// Rate of the cash premium made available if the securities holder consents or participates to an event, e.g. consent fees.
	CashIncentive *RateFormat1Choice `xml:"CshIncntiv,omitempty"`

	// Cash rate made available in an offer in order to encourage participation in the offer.
	SolicitationFee *RateFormat1Choice `xml:"SlctnFee,omitempty"`

	// A maximum percentage of shares available through the over subscription privilege, usually a percentage of the basic subscription shares, eg, an account owner subscribing to 100 shares may over subscribe to a maximum of 50 additional shares when the over subscription maximum is 50%.
	MaximumAllowedOversubscription *RateFormat1Choice `xml:"MaxAllwdOvrsbcpt,omitempty"`

	// Rate used for additional tax that cannot be categorised.
	AdditionalTax *RateAndAmountFormat1Choice `xml:"AddtlTax,omitempty"`

	// Amount in its original currency when conversion from/into another currency has occurred.
	OriginalAmount *ActiveCurrencyAndAmount `xml:"OrgnlAmt,omitempty"`

	// Provides information about a foreign exchange.
	ExchangeRate *ForeignExchangeTerms8 `xml:"XchgRate,omitempty"`

	// Rate applicable to the event announced, eg, redemption rate for a redemption event.
	ApplicableRate *RateFormat1Choice `xml:"AplblRate,omitempty"`
}

func (c *CorporateActionRate2) AddWithholdingTax() *RateFormat1Choice {
	c.WithholdingTax = new(RateFormat1Choice)
	return c.WithholdingTax
}

func (c *CorporateActionRate2) AddWithholdingOfForeignTax() *RateAndAmountFormat1Choice {
	c.WithholdingOfForeignTax = new(RateAndAmountFormat1Choice)
	return c.WithholdingOfForeignTax
}

func (c *CorporateActionRate2) AddWithholdingOfLocalTax() *RateAndAmountFormat1Choice {
	c.WithholdingOfLocalTax = new(RateAndAmountFormat1Choice)
	return c.WithholdingOfLocalTax
}

func (c *CorporateActionRate2) AddGermanLocalTax1() *RateAndAmountFormat1Choice {
	c.GermanLocalTax1 = new(RateAndAmountFormat1Choice)
	return c.GermanLocalTax1
}

func (c *CorporateActionRate2) AddGermanLocalTax2() *RateAndAmountFormat1Choice {
	c.GermanLocalTax2 = new(RateAndAmountFormat1Choice)
	return c.GermanLocalTax2
}

func (c *CorporateActionRate2) AddGermanLocalTax3() *RateAndAmountFormat1Choice {
	c.GermanLocalTax3 = new(RateAndAmountFormat1Choice)
	return c.GermanLocalTax3
}

func (c *CorporateActionRate2) AddGermanLocalTax4() *RateAndAmountFormat1Choice {
	c.GermanLocalTax4 = new(RateAndAmountFormat1Choice)
	return c.GermanLocalTax4
}

func (c *CorporateActionRate2) AddTaxOnIncome() *RateFormat1Choice {
	c.TaxOnIncome = new(RateFormat1Choice)
	return c.TaxOnIncome
}

func (c *CorporateActionRate2) AddTaxOnProfit() *RateFormat1Choice {
	c.TaxOnProfit = new(RateFormat1Choice)
	return c.TaxOnProfit
}

func (c *CorporateActionRate2) AddTaxReclaim() *RateFormat1Choice {
	c.TaxReclaim = new(RateFormat1Choice)
	return c.TaxReclaim
}

func (c *CorporateActionRate2) AddFiscalStamp() *RateFormat1Choice {
	c.FiscalStamp = new(RateFormat1Choice)
	return c.FiscalStamp
}

func (c *CorporateActionRate2) AddProration() *RateFormat1Choice {
	c.Proration = new(RateFormat1Choice)
	return c.Proration
}

func (c *CorporateActionRate2) AddNewToOld() *RatioFormat2Choice {
	c.NewToOld = new(RatioFormat2Choice)
	return c.NewToOld
}

func (c *CorporateActionRate2) AddNewSecuritiesToUnderlyingSecurities() *RatioFormat2Choice {
	c.NewSecuritiesToUnderlyingSecurities = new(RatioFormat2Choice)
	return c.NewSecuritiesToUnderlyingSecurities
}

func (c *CorporateActionRate2) AddAdditionalQuantityForExistingSecurities() *RatioFormat1Choice {
	c.AdditionalQuantityForExistingSecurities = new(RatioFormat1Choice)
	return c.AdditionalQuantityForExistingSecurities
}

func (c *CorporateActionRate2) AddAdditionalQuantityForSubscribedResultantSecurities() *RatioFormat1Choice {
	c.AdditionalQuantityForSubscribedResultantSecurities = new(RatioFormat1Choice)
	return c.AdditionalQuantityForSubscribedResultantSecurities
}

func (c *CorporateActionRate2) AddRelatedTax() *RelatedTaxType1 {
	c.RelatedTax = new(RelatedTaxType1)
	return c.RelatedTax
}

func (c *CorporateActionRate2) AddNonResidentRate() *RateAndAmountFormat1Choice {
	c.NonResidentRate = new(RateAndAmountFormat1Choice)
	return c.NonResidentRate
}

func (c *CorporateActionRate2) AddCharges() *RateAndAmountFormat1Choice {
	c.Charges = new(RateAndAmountFormat1Choice)
	return c.Charges
}

func (c *CorporateActionRate2) AddInterestForUsedPayment() *RateAndAmountFormat1Choice {
	c.InterestForUsedPayment = new(RateAndAmountFormat1Choice)
	return c.InterestForUsedPayment
}

func (c *CorporateActionRate2) AddIndexFactor() *RateAndAmountFormat1Choice {
	c.IndexFactor = new(RateAndAmountFormat1Choice)
	return c.IndexFactor
}

func (c *CorporateActionRate2) AddFullyFranked() *RateAndAmountFormat1Choice {
	c.FullyFranked = new(RateAndAmountFormat1Choice)
	return c.FullyFranked
}

func (c *CorporateActionRate2) AddGrossDividend() *GrossDividendRate1Choice {
	c.GrossDividend = new(GrossDividendRate1Choice)
	return c.GrossDividend
}

func (c *CorporateActionRate2) AddNetDividend() *NetDividendRate1Choice {
	c.NetDividend = new(NetDividendRate1Choice)
	return c.NetDividend
}

func (c *CorporateActionRate2) AddFinalDividend() *AmountAndRateFormat2Choice {
	c.FinalDividend = new(AmountAndRateFormat2Choice)
	return c.FinalDividend
}

func (c *CorporateActionRate2) AddProvisionalDividend() *AmountAndRateFormat2Choice {
	c.ProvisionalDividend = new(AmountAndRateFormat2Choice)
	return c.ProvisionalDividend
}

func (c *CorporateActionRate2) AddCashIncentive() *RateFormat1Choice {
	c.CashIncentive = new(RateFormat1Choice)
	return c.CashIncentive
}

func (c *CorporateActionRate2) AddSolicitationFee() *RateFormat1Choice {
	c.SolicitationFee = new(RateFormat1Choice)
	return c.SolicitationFee
}

func (c *CorporateActionRate2) AddMaximumAllowedOversubscription() *RateFormat1Choice {
	c.MaximumAllowedOversubscription = new(RateFormat1Choice)
	return c.MaximumAllowedOversubscription
}

func (c *CorporateActionRate2) AddAdditionalTax() *RateAndAmountFormat1Choice {
	c.AdditionalTax = new(RateAndAmountFormat1Choice)
	return c.AdditionalTax
}

func (c *CorporateActionRate2) SetOriginalAmount(value, currency string) {
	c.OriginalAmount = NewActiveCurrencyAndAmount(value, currency)
}

func (c *CorporateActionRate2) AddExchangeRate() *ForeignExchangeTerms8 {
	c.ExchangeRate = new(ForeignExchangeTerms8)
	return c.ExchangeRate
}

func (c *CorporateActionRate2) AddApplicableRate() *RateFormat1Choice {
	c.ApplicableRate = new(RateFormat1Choice)
	return c.ApplicableRate
}
