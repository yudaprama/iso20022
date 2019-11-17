package model

// Provides information about the rates related to securities movement.
type RateDetails2 struct {

	// Rate used for additional tax that cannot be categorised.
	AdditionalTax *RateAndAmountFormat5Choice `xml:"AddtlTax,omitempty"`

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

	// Actual interest rate used for the payment of the interest for the specified interest period.
	InterestRateUsedForPayment []*InterestRateUsedForPaymentFormat2Choice `xml:"IntrstRateUsdForPmt,omitempty"`

	// Cash dividend amount per equity after deductions or allowances have been made.
	NetDividendRate []*NetDividendRateFormat2Choice `xml:"NetDvddRate,omitempty"`

	// Rate per share to which a non-resident is entitled.
	NonResidentRate *RateAndAmountFormat5Choice `xml:"NonResdtRate,omitempty"`

	// Dividend is provisional.
	ProvisionalDividendRate *ActiveCurrencyAnd13DecimalAmount `xml:"PrvsnlDvddRate,omitempty"`

	// Rate applicable to the event announced, for example, redemption rate for a redemption event.
	ApplicableRate *PercentageRate `xml:"AplblRate,omitempty"`

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
}

func (r *RateDetails2) AddAdditionalTax() *RateAndAmountFormat5Choice {
	r.AdditionalTax = new(RateAndAmountFormat5Choice)
	return r.AdditionalTax
}

func (r *RateDetails2) AddChargesFees() *RateAndAmountFormat5Choice {
	r.ChargesFees = new(RateAndAmountFormat5Choice)
	return r.ChargesFees
}

func (r *RateDetails2) SetFinalDividendRate(value, currency string) {
	r.FinalDividendRate = NewActiveCurrencyAnd13DecimalAmount(value, currency)
}

func (r *RateDetails2) SetFiscalStamp(value string) {
	r.FiscalStamp = (*PercentageRate)(&value)
}

func (r *RateDetails2) AddFullyFrankedRate() *RateAndAmountFormat5Choice {
	r.FullyFrankedRate = new(RateAndAmountFormat5Choice)
	return r.FullyFrankedRate
}

func (r *RateDetails2) AddGrossDividendRate() *GrossDividendRateFormat2Choice {
	newValue := new(GrossDividendRateFormat2Choice)
	r.GrossDividendRate = append(r.GrossDividendRate, newValue)
	return newValue
}

func (r *RateDetails2) SetCashIncentiveRate(value string) {
	r.CashIncentiveRate = (*PercentageRate)(&value)
}

func (r *RateDetails2) AddInterestRateUsedForPayment() *InterestRateUsedForPaymentFormat2Choice {
	newValue := new(InterestRateUsedForPaymentFormat2Choice)
	r.InterestRateUsedForPayment = append(r.InterestRateUsedForPayment, newValue)
	return newValue
}

func (r *RateDetails2) AddNetDividendRate() *NetDividendRateFormat2Choice {
	newValue := new(NetDividendRateFormat2Choice)
	r.NetDividendRate = append(r.NetDividendRate, newValue)
	return newValue
}

func (r *RateDetails2) AddNonResidentRate() *RateAndAmountFormat5Choice {
	r.NonResidentRate = new(RateAndAmountFormat5Choice)
	return r.NonResidentRate
}

func (r *RateDetails2) SetProvisionalDividendRate(value, currency string) {
	r.ProvisionalDividendRate = NewActiveCurrencyAnd13DecimalAmount(value, currency)
}

func (r *RateDetails2) SetApplicableRate(value string) {
	r.ApplicableRate = (*PercentageRate)(&value)
}

func (r *RateDetails2) SetSolicitationFeeRate(value string) {
	r.SolicitationFeeRate = (*PercentageRate)(&value)
}

func (r *RateDetails2) AddTaxCreditRate() *TaxCreditRateFormat2Choice {
	newValue := new(TaxCreditRateFormat2Choice)
	r.TaxCreditRate = append(r.TaxCreditRate, newValue)
	return newValue
}

func (r *RateDetails2) AddTaxRelatedRate() *RateTypeAndAmountAndStatus6 {
	newValue := new(RateTypeAndAmountAndStatus6)
	r.TaxRelatedRate = append(r.TaxRelatedRate, newValue)
	return newValue
}

func (r *RateDetails2) SetWithholdingTaxRate(value string) {
	r.WithholdingTaxRate = (*PercentageRate)(&value)
}

func (r *RateDetails2) SetTaxOnIncome(value string) {
	r.TaxOnIncome = (*PercentageRate)(&value)
}

func (r *RateDetails2) SetTaxOnProfits(value string) {
	r.TaxOnProfits = (*PercentageRate)(&value)
}

func (r *RateDetails2) SetTaxReclaimRate(value string) {
	r.TaxReclaimRate = (*PercentageRate)(&value)
}

func (r *RateDetails2) AddWithholdingOfForeignTax() *RateAndAmountFormat5Choice {
	r.WithholdingOfForeignTax = new(RateAndAmountFormat5Choice)
	return r.WithholdingOfForeignTax
}

func (r *RateDetails2) AddWithholdingOfLocalTax() *RateAndAmountFormat5Choice {
	r.WithholdingOfLocalTax = new(RateAndAmountFormat5Choice)
	return r.WithholdingOfLocalTax
}
