package model

// Provides information about the rates related to securities movement.
type RateDetails15 struct {

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
	GrossDividendRate []*GrossDividendRateFormat8Choice `xml:"GrssDvddRate,omitempty"`

	// Cash rate made available, as an incentive, in addition to the solicitation fee, in order to encourage early participation in an offer.
	EarlySolicitationFeeRate *RateAndAmountFormat5Choice `xml:"EarlySlctnFeeRate,omitempty"`

	// Cash rate made available in an event in order to encourage participation in the offer. As information, Payment is made to a third party who has solicited an entity to take part in the offer.
	ThirdPartyIncentiveRate *RateAndAmountFormat5Choice `xml:"ThrdPtyIncntivRate,omitempty"`

	// Actual interest rate used for the payment of the interest for the specified interest period.
	InterestRateUsedForPayment []*InterestRateUsedForPaymentFormat2Choice `xml:"IntrstRateUsdForPmt,omitempty"`

	// Cash dividend amount per equity after deductions or allowances have been made.
	NetDividendRate []*NetDividendRateFormat10Choice `xml:"NetDvddRate,omitempty"`

	// Rate per share to which a non-resident is entitled.
	NonResidentRate *RateAndAmountFormat5Choice `xml:"NonResdtRate,omitempty"`

	// Dividend is provisional.
	ProvisionalDividendRate *ActiveCurrencyAnd13DecimalAmount `xml:"PrvsnlDvddRate,omitempty"`

	// Rate applicable to the event announced, for example, redemption rate for a redemption event.
	ApplicableRate *PercentageRate `xml:"AplblRate,omitempty"`

	// Rate of the cash premium made available if the securities holder consents or participates to an event, for example consent fees or solicitation fee.
	SolicitationFeeRate *RateAndAmountFormat5Choice `xml:"SlctnFeeRate,omitempty"`

	// Amount of money per equity allocated as the result of a tax credit.
	TaxCreditRate []*TaxCreditRateFormat2Choice `xml:"TaxCdtRate,omitempty"`

	// Percentage of the gross dividend rate on which tax must be paid .
	TaxRelatedRate []*RateTypeAndAmountAndStatus6 `xml:"TaxRltdRate,omitempty"`

	// Percentage of a cash distribution that will be withheld by a tax authority.
	WithholdingTaxRate []*RateFormat11Choice `xml:"WhldgTaxRate,omitempty"`

	// Taxation applied on an amount clearly identified as an income.
	TaxOnIncome *RateAndAmountFormat5Choice `xml:"TaxOnIncm,omitempty"`

	// Taxation applied on an amount clearly identified as capital profits, capital gains.
	TaxOnProfits *PercentageRate `xml:"TaxOnPrfts,omitempty"`

	// Percentage of cash that was paid in excess of actual tax obligation and was reclaimed.
	TaxReclaimRate *PercentageRate `xml:"TaxRclmRate,omitempty"`

	// Rate at which the income will be withheld by the jurisdiction in which the income was originally paid, for which relief at source and/or reclaim may be possible.
	WithholdingOfForeignTax []*RateAndAmountFormat21Choice `xml:"WhldgOfFrgnTax,omitempty"`

	// Rate at which the income will be withheld by the jurisdiction in which the account owner is located, for which relief at source and/or reclaim may be possible.
	WithholdingOfLocalTax []*RateAndAmountFormat21Choice `xml:"WhldgOfLclTax,omitempty"`

	// Portion of the fund distribution which represents the average accrued income included in the purchase price for units bought during the account period.
	EqualisationRate *ActiveCurrencyAnd13DecimalAmount `xml:"EqulstnRate,omitempty"`
}

func (r *RateDetails15) AddAdditionalTax() *RateAndAmountFormat5Choice {
	r.AdditionalTax = new(RateAndAmountFormat5Choice)
	return r.AdditionalTax
}

func (r *RateDetails15) AddChargesFees() *RateAndAmountFormat5Choice {
	r.ChargesFees = new(RateAndAmountFormat5Choice)
	return r.ChargesFees
}

func (r *RateDetails15) SetFinalDividendRate(value, currency string) {
	r.FinalDividendRate = NewActiveCurrencyAnd13DecimalAmount(value, currency)
}

func (r *RateDetails15) SetFiscalStamp(value string) {
	r.FiscalStamp = (*PercentageRate)(&value)
}

func (r *RateDetails15) AddFullyFrankedRate() *RateAndAmountFormat5Choice {
	r.FullyFrankedRate = new(RateAndAmountFormat5Choice)
	return r.FullyFrankedRate
}

func (r *RateDetails15) AddGrossDividendRate() *GrossDividendRateFormat8Choice {
	newValue := new(GrossDividendRateFormat8Choice)
	r.GrossDividendRate = append(r.GrossDividendRate, newValue)
	return newValue
}

func (r *RateDetails15) AddEarlySolicitationFeeRate() *RateAndAmountFormat5Choice {
	r.EarlySolicitationFeeRate = new(RateAndAmountFormat5Choice)
	return r.EarlySolicitationFeeRate
}

func (r *RateDetails15) AddThirdPartyIncentiveRate() *RateAndAmountFormat5Choice {
	r.ThirdPartyIncentiveRate = new(RateAndAmountFormat5Choice)
	return r.ThirdPartyIncentiveRate
}

func (r *RateDetails15) AddInterestRateUsedForPayment() *InterestRateUsedForPaymentFormat2Choice {
	newValue := new(InterestRateUsedForPaymentFormat2Choice)
	r.InterestRateUsedForPayment = append(r.InterestRateUsedForPayment, newValue)
	return newValue
}

func (r *RateDetails15) AddNetDividendRate() *NetDividendRateFormat10Choice {
	newValue := new(NetDividendRateFormat10Choice)
	r.NetDividendRate = append(r.NetDividendRate, newValue)
	return newValue
}

func (r *RateDetails15) AddNonResidentRate() *RateAndAmountFormat5Choice {
	r.NonResidentRate = new(RateAndAmountFormat5Choice)
	return r.NonResidentRate
}

func (r *RateDetails15) SetProvisionalDividendRate(value, currency string) {
	r.ProvisionalDividendRate = NewActiveCurrencyAnd13DecimalAmount(value, currency)
}

func (r *RateDetails15) SetApplicableRate(value string) {
	r.ApplicableRate = (*PercentageRate)(&value)
}

func (r *RateDetails15) AddSolicitationFeeRate() *RateAndAmountFormat5Choice {
	r.SolicitationFeeRate = new(RateAndAmountFormat5Choice)
	return r.SolicitationFeeRate
}

func (r *RateDetails15) AddTaxCreditRate() *TaxCreditRateFormat2Choice {
	newValue := new(TaxCreditRateFormat2Choice)
	r.TaxCreditRate = append(r.TaxCreditRate, newValue)
	return newValue
}

func (r *RateDetails15) AddTaxRelatedRate() *RateTypeAndAmountAndStatus6 {
	newValue := new(RateTypeAndAmountAndStatus6)
	r.TaxRelatedRate = append(r.TaxRelatedRate, newValue)
	return newValue
}

func (r *RateDetails15) AddWithholdingTaxRate() *RateFormat11Choice {
	newValue := new(RateFormat11Choice)
	r.WithholdingTaxRate = append(r.WithholdingTaxRate, newValue)
	return newValue
}

func (r *RateDetails15) AddTaxOnIncome() *RateAndAmountFormat5Choice {
	r.TaxOnIncome = new(RateAndAmountFormat5Choice)
	return r.TaxOnIncome
}

func (r *RateDetails15) SetTaxOnProfits(value string) {
	r.TaxOnProfits = (*PercentageRate)(&value)
}

func (r *RateDetails15) SetTaxReclaimRate(value string) {
	r.TaxReclaimRate = (*PercentageRate)(&value)
}

func (r *RateDetails15) AddWithholdingOfForeignTax() *RateAndAmountFormat21Choice {
	newValue := new(RateAndAmountFormat21Choice)
	r.WithholdingOfForeignTax = append(r.WithholdingOfForeignTax, newValue)
	return newValue
}

func (r *RateDetails15) AddWithholdingOfLocalTax() *RateAndAmountFormat21Choice {
	newValue := new(RateAndAmountFormat21Choice)
	r.WithholdingOfLocalTax = append(r.WithholdingOfLocalTax, newValue)
	return newValue
}

func (r *RateDetails15) SetEqualisationRate(value, currency string) {
	r.EqualisationRate = NewActiveCurrencyAnd13DecimalAmount(value, currency)
}
