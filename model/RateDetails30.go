package model

// Provides information about the rates related to securities movement.
type RateDetails30 struct {

	// Rate used for additional tax that cannot be categorised.
	AdditionalTax *RateAndAmountFormat43Choice `xml:"AddtlTax,omitempty"`

	// Rate used to calculate the amount of the charges/fees that cannot be categorised.
	ChargesFees *RateAndAmountFormat43Choice `xml:"ChrgsFees,omitempty"`

	// Percentage of fiscal tax to apply.
	FiscalStamp *PercentageRate `xml:"FsclStmp,omitempty"`

	// Cash dividend amount per equity before deductions or allowances have been made.
	GrossDividendRate []*GrossDividendRateFormat24Choice `xml:"GrssDvddRate,omitempty"`

	// Cash rate made available, as an incentive, in addition to the solicitation fee, in order to encourage early participation in an offer.
	EarlySolicitationFeeRate *SolicitationFeeRateFormat9Choice `xml:"EarlySlctnFeeRate,omitempty"`

	// Cash rate made available in an event in order to encourage participation in the offer. As information, Payment is made to a third party who has solicited an entity to take part in the offer.
	ThirdPartyIncentiveRate *RateAndAmountFormat43Choice `xml:"ThrdPtyIncntivRate,omitempty"`

	// Actual interest rate used for the payment of the interest for the specified interest period.
	InterestRateUsedForPayment []*InterestRateUsedForPaymentFormat9Choice `xml:"IntrstRateUsdForPmt,omitempty"`

	// Cash dividend amount per equity after deductions or allowances have been made.
	NetDividendRate []*NetDividendRateFormat26Choice `xml:"NetDvddRate,omitempty"`

	// Rate per share to which a non-resident is entitled.
	NonResidentRate *RateAndAmountFormat43Choice `xml:"NonResdtRate,omitempty"`

	// Rate applicable to the event announced, for example, redemption rate for a redemption event.
	ApplicableRate *PercentageRate `xml:"AplblRate,omitempty"`

	// Rate of the cash premium made available if the securities holder consents or participates to an event, for example consent fees or solicitation fee.
	SolicitationFeeRate *SolicitationFeeRateFormat9Choice `xml:"SlctnFeeRate,omitempty"`

	// Amount of money per equity allocated as the result of a tax credit.
	TaxCreditRate []*TaxCreditRateFormat9Choice `xml:"TaxCdtRate,omitempty"`

	// Percentage of a cash distribution that will be withheld by the tax authorities of the jurisdiction of the issuer, for which a relief at source and/or reclaim may be possible.
	WithholdingTaxRate []*RateAndAmountFormat45Choice `xml:"WhldgTaxRate,omitempty"`

	// Rate at which the income will be withheld by a jurisdiction other than the jurisdiction of the issuer’s country of tax incorporation, for which a relief at source and/or reclaim may be possible. It is levied in complement or offset of the withholding tax rate (TAXR) levied by the jurisdiction of the issuer’s tax domicile.
	SecondLevelTax []*RateAndAmountFormat45Choice `xml:"ScndLvlTax,omitempty"`

	// Taxation applied on an amount clearly identified as an income.
	TaxOnIncome *RateAndAmountFormat43Choice `xml:"TaxOnIncm,omitempty"`

	// Taxation applied on an amount clearly identified as capital profits, capital gains.
	TaxOnProfits *PercentageRate `xml:"TaxOnPrfts,omitempty"`

	// Percentage of cash that was paid in excess of actual tax obligation and was reclaimed.
	TaxReclaimRate *PercentageRate `xml:"TaxRclmRate,omitempty"`

	// Portion of the fund distribution which represents the average accrued income included in the purchase price for units bought during the account period.
	EqualisationRate *RestrictedFINActiveCurrencyAnd13DecimalAmount `xml:"EqulstnRate,omitempty"`
}

func (r *RateDetails30) AddAdditionalTax() *RateAndAmountFormat43Choice {
	r.AdditionalTax = new(RateAndAmountFormat43Choice)
	return r.AdditionalTax
}

func (r *RateDetails30) AddChargesFees() *RateAndAmountFormat43Choice {
	r.ChargesFees = new(RateAndAmountFormat43Choice)
	return r.ChargesFees
}

func (r *RateDetails30) SetFiscalStamp(value string) {
	r.FiscalStamp = (*PercentageRate)(&value)
}

func (r *RateDetails30) AddGrossDividendRate() *GrossDividendRateFormat24Choice {
	newValue := new(GrossDividendRateFormat24Choice)
	r.GrossDividendRate = append(r.GrossDividendRate, newValue)
	return newValue
}

func (r *RateDetails30) AddEarlySolicitationFeeRate() *SolicitationFeeRateFormat9Choice {
	r.EarlySolicitationFeeRate = new(SolicitationFeeRateFormat9Choice)
	return r.EarlySolicitationFeeRate
}

func (r *RateDetails30) AddThirdPartyIncentiveRate() *RateAndAmountFormat43Choice {
	r.ThirdPartyIncentiveRate = new(RateAndAmountFormat43Choice)
	return r.ThirdPartyIncentiveRate
}

func (r *RateDetails30) AddInterestRateUsedForPayment() *InterestRateUsedForPaymentFormat9Choice {
	newValue := new(InterestRateUsedForPaymentFormat9Choice)
	r.InterestRateUsedForPayment = append(r.InterestRateUsedForPayment, newValue)
	return newValue
}

func (r *RateDetails30) AddNetDividendRate() *NetDividendRateFormat26Choice {
	newValue := new(NetDividendRateFormat26Choice)
	r.NetDividendRate = append(r.NetDividendRate, newValue)
	return newValue
}

func (r *RateDetails30) AddNonResidentRate() *RateAndAmountFormat43Choice {
	r.NonResidentRate = new(RateAndAmountFormat43Choice)
	return r.NonResidentRate
}

func (r *RateDetails30) SetApplicableRate(value string) {
	r.ApplicableRate = (*PercentageRate)(&value)
}

func (r *RateDetails30) AddSolicitationFeeRate() *SolicitationFeeRateFormat9Choice {
	r.SolicitationFeeRate = new(SolicitationFeeRateFormat9Choice)
	return r.SolicitationFeeRate
}

func (r *RateDetails30) AddTaxCreditRate() *TaxCreditRateFormat9Choice {
	newValue := new(TaxCreditRateFormat9Choice)
	r.TaxCreditRate = append(r.TaxCreditRate, newValue)
	return newValue
}

func (r *RateDetails30) AddWithholdingTaxRate() *RateAndAmountFormat45Choice {
	newValue := new(RateAndAmountFormat45Choice)
	r.WithholdingTaxRate = append(r.WithholdingTaxRate, newValue)
	return newValue
}

func (r *RateDetails30) AddSecondLevelTax() *RateAndAmountFormat45Choice {
	newValue := new(RateAndAmountFormat45Choice)
	r.SecondLevelTax = append(r.SecondLevelTax, newValue)
	return newValue
}

func (r *RateDetails30) AddTaxOnIncome() *RateAndAmountFormat43Choice {
	r.TaxOnIncome = new(RateAndAmountFormat43Choice)
	return r.TaxOnIncome
}

func (r *RateDetails30) SetTaxOnProfits(value string) {
	r.TaxOnProfits = (*PercentageRate)(&value)
}

func (r *RateDetails30) SetTaxReclaimRate(value string) {
	r.TaxReclaimRate = (*PercentageRate)(&value)
}

func (r *RateDetails30) SetEqualisationRate(value, currency string) {
	r.EqualisationRate = NewRestrictedFINActiveCurrencyAnd13DecimalAmount(value, currency)
}
