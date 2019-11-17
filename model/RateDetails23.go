package model

// Provides information about the rates related to securities movement.
type RateDetails23 struct {

	// Rate used for additional tax that cannot be categorised.
	AdditionalTax *RateAndAmountFormat39Choice `xml:"AddtlTax,omitempty"`

	// Rate used to calculate the amount of the charges/fees that cannot be categorised.
	ChargesFees *RateAndAmountFormat39Choice `xml:"ChrgsFees,omitempty"`

	// Percentage of fiscal tax to apply.
	FiscalStamp *PercentageRate `xml:"FsclStmp,omitempty"`

	// Rate resulting from a fully franked dividend paid by a company; rate includes tax credit for companies that have made sufficient tax payments during fiscal period.
	FullyFrankedRate *RateAndAmountFormat39Choice `xml:"FullyFrnkdRate,omitempty"`

	// Cash dividend amount per equity before deductions or allowances have been made.
	GrossDividendRate []*GrossDividendRateFormat22Choice `xml:"GrssDvddRate,omitempty"`

	// Cash rate made available, as an incentive, in addition to the solicitation fee, in order to encourage early participation in an offer.
	EarlySolicitationFeeRate *SolicitationFeeRateFormat8Choice `xml:"EarlySlctnFeeRate,omitempty"`

	// Cash rate made available in an event in order to encourage participation in the offer. As information, Payment is made to a third party who has solicited an entity to take part in the offer.
	ThirdPartyIncentiveRate *RateAndAmountFormat39Choice `xml:"ThrdPtyIncntivRate,omitempty"`

	// Actual interest rate used for the payment of the interest for the specified interest period.
	InterestRateUsedForPayment []*InterestRateUsedForPaymentFormat7Choice `xml:"IntrstRateUsdForPmt,omitempty"`

	// Cash dividend amount per equity after deductions or allowances have been made.
	NetDividendRate []*NetDividendRateFormat24Choice `xml:"NetDvddRate,omitempty"`

	// Rate per share to which a non-resident is entitled.
	NonResidentRate *RateAndAmountFormat39Choice `xml:"NonResdtRate,omitempty"`

	// Rate applicable to the event announced, for example, redemption rate for a redemption event.
	ApplicableRate *PercentageRate `xml:"AplblRate,omitempty"`

	// Rate of the cash premium made available if the securities holder consents or participates to an event, for example consent fees or solicitation fee.
	SolicitationFeeRate *SolicitationFeeRateFormat8Choice `xml:"SlctnFeeRate,omitempty"`

	// Amount of money per equity allocated as the result of a tax credit.
	TaxCreditRate []*TaxCreditRateFormat7Choice `xml:"TaxCdtRate,omitempty"`

	// Percentage of a cash distribution that will be withheld by the tax authorities of the jurisdiction of the issuer, for which a relief at source and/or reclaim may be possible.
	WithholdingTaxRate []*RateAndAmountFormat40Choice `xml:"WhldgTaxRate,omitempty"`

	// Rate at which the income will be withheld by a jurisdiction other than the jurisdiction of the issuer’s country of tax incorporation, for which a relief at source and/or reclaim may be possible. It is levied in complement or offset of the withholding tax rate (TAXR) levied by the jurisdiction of the issuer’s tax domicile.
	SecondLevelTax []*RateAndAmountFormat40Choice `xml:"ScndLvlTax,omitempty"`

	// Taxation applied on an amount clearly identified as an income.
	TaxOnIncome *RateAndAmountFormat39Choice `xml:"TaxOnIncm,omitempty"`

	// Taxation applied on an amount clearly identified as capital profits, capital gains.
	TaxOnProfits *PercentageRate `xml:"TaxOnPrfts,omitempty"`

	// Percentage of cash that was paid in excess of actual tax obligation and was reclaimed.
	TaxReclaimRate *PercentageRate `xml:"TaxRclmRate,omitempty"`

	// Portion of the fund distribution which represents the average accrued income included in the purchase price for units bought during the account period.
	EqualisationRate *ActiveCurrencyAnd13DecimalAmount `xml:"EqulstnRate,omitempty"`
}

func (r *RateDetails23) AddAdditionalTax() *RateAndAmountFormat39Choice {
	r.AdditionalTax = new(RateAndAmountFormat39Choice)
	return r.AdditionalTax
}

func (r *RateDetails23) AddChargesFees() *RateAndAmountFormat39Choice {
	r.ChargesFees = new(RateAndAmountFormat39Choice)
	return r.ChargesFees
}

func (r *RateDetails23) SetFiscalStamp(value string) {
	r.FiscalStamp = (*PercentageRate)(&value)
}

func (r *RateDetails23) AddFullyFrankedRate() *RateAndAmountFormat39Choice {
	r.FullyFrankedRate = new(RateAndAmountFormat39Choice)
	return r.FullyFrankedRate
}

func (r *RateDetails23) AddGrossDividendRate() *GrossDividendRateFormat22Choice {
	newValue := new(GrossDividendRateFormat22Choice)
	r.GrossDividendRate = append(r.GrossDividendRate, newValue)
	return newValue
}

func (r *RateDetails23) AddEarlySolicitationFeeRate() *SolicitationFeeRateFormat8Choice {
	r.EarlySolicitationFeeRate = new(SolicitationFeeRateFormat8Choice)
	return r.EarlySolicitationFeeRate
}

func (r *RateDetails23) AddThirdPartyIncentiveRate() *RateAndAmountFormat39Choice {
	r.ThirdPartyIncentiveRate = new(RateAndAmountFormat39Choice)
	return r.ThirdPartyIncentiveRate
}

func (r *RateDetails23) AddInterestRateUsedForPayment() *InterestRateUsedForPaymentFormat7Choice {
	newValue := new(InterestRateUsedForPaymentFormat7Choice)
	r.InterestRateUsedForPayment = append(r.InterestRateUsedForPayment, newValue)
	return newValue
}

func (r *RateDetails23) AddNetDividendRate() *NetDividendRateFormat24Choice {
	newValue := new(NetDividendRateFormat24Choice)
	r.NetDividendRate = append(r.NetDividendRate, newValue)
	return newValue
}

func (r *RateDetails23) AddNonResidentRate() *RateAndAmountFormat39Choice {
	r.NonResidentRate = new(RateAndAmountFormat39Choice)
	return r.NonResidentRate
}

func (r *RateDetails23) SetApplicableRate(value string) {
	r.ApplicableRate = (*PercentageRate)(&value)
}

func (r *RateDetails23) AddSolicitationFeeRate() *SolicitationFeeRateFormat8Choice {
	r.SolicitationFeeRate = new(SolicitationFeeRateFormat8Choice)
	return r.SolicitationFeeRate
}

func (r *RateDetails23) AddTaxCreditRate() *TaxCreditRateFormat7Choice {
	newValue := new(TaxCreditRateFormat7Choice)
	r.TaxCreditRate = append(r.TaxCreditRate, newValue)
	return newValue
}

func (r *RateDetails23) AddWithholdingTaxRate() *RateAndAmountFormat40Choice {
	newValue := new(RateAndAmountFormat40Choice)
	r.WithholdingTaxRate = append(r.WithholdingTaxRate, newValue)
	return newValue
}

func (r *RateDetails23) AddSecondLevelTax() *RateAndAmountFormat40Choice {
	newValue := new(RateAndAmountFormat40Choice)
	r.SecondLevelTax = append(r.SecondLevelTax, newValue)
	return newValue
}

func (r *RateDetails23) AddTaxOnIncome() *RateAndAmountFormat39Choice {
	r.TaxOnIncome = new(RateAndAmountFormat39Choice)
	return r.TaxOnIncome
}

func (r *RateDetails23) SetTaxOnProfits(value string) {
	r.TaxOnProfits = (*PercentageRate)(&value)
}

func (r *RateDetails23) SetTaxReclaimRate(value string) {
	r.TaxReclaimRate = (*PercentageRate)(&value)
}

func (r *RateDetails23) SetEqualisationRate(value, currency string) {
	r.EqualisationRate = NewActiveCurrencyAnd13DecimalAmount(value, currency)
}
