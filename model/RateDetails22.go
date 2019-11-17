package model

// Provides information about the rates related to securities movement.
type RateDetails22 struct {

	// Rate used for additional tax that cannot be categorised.
	AdditionalTax *RateAndAmountFormat37Choice `xml:"AddtlTax,omitempty"`

	// Cash dividend amount per equity before deductions or allowances have been made.
	GrossDividendRate []*GrossDividendRateFormat20Choice `xml:"GrssDvddRate,omitempty"`

	// The actual interest rate used for the payment of the interest for the specified interest period.
	// Usage guideline: It is used to provide the applicable rate for the current payment, after all calculations have been performed, that is, application of period and method of interest computation.
	InterestRateUsedForPayment []*InterestRateUsedForPaymentFormat8Choice `xml:"IntrstRateUsdForPmt,omitempty"`

	// Percentage of a cash distribution that will be withheld by the tax authorities of the jurisdiction of the issuer, for which a relief at source and/or reclaim may be possible.
	WithholdingTaxRate []*RateAndAmountFormat41Choice `xml:"WhldgTaxRate,omitempty"`

	// Rate at which the income will be withheld by a jurisdiction other than the jurisdiction of the issuer’s country of tax incorporation, for which a relief at source and/or reclaim may be possible. It is levied in complement or offset of the withholding tax rate (TAXR) levied by the jurisdiction of the issuer’s tax domicile.
	SecondLevelTax []*RateAndAmountFormat41Choice `xml:"ScndLvlTax,omitempty"`

	// Rate used to calculate the amount of the charges/fees that cannot be categorised.
	ChargesFees *RateAndAmountFormat37Choice `xml:"ChrgsFees,omitempty"`

	// Cash rate made available, as an incentive, in addition to the solicitation fee, in order to encourage early participation in an offer.
	EarlySolicitationFeeRate *SolicitationFeeRateFormat7Choice `xml:"EarlySlctnFeeRate,omitempty"`

	// Percentage of fiscal tax to apply.
	FiscalStamp *RateFormat3Choice `xml:"FsclStmp,omitempty"`

	// Rate resulting from a fully franked dividend paid by a company; rate includes tax credit for companies that have made sufficient tax payments during fiscal period.
	FullyFrankedRate *RateAndAmountFormat37Choice `xml:"FullyFrnkdRate,omitempty"`

	// Cash rate made available in an event in order to encourage participation in the offer. As information, payment is made to a third party who has solicited an entity to take part in the offer.
	ThirdPartyIncentiveRate *RateFormat20Choice `xml:"ThrdPtyIncntivRate,omitempty"`

	// Cash dividend amount per equity after deductions or allowances have been made.
	NetDividendRate []*NetDividendRateFormat22Choice `xml:"NetDvddRate,omitempty"`

	// Rate per share to which a non-resident is entitled.
	NonResidentRate *RateAndAmountFormat37Choice `xml:"NonResdtRate,omitempty"`

	// Rate applicable to the event announced, for example, redemption rate for a redemption event.
	ApplicableRate *RateFormat3Choice `xml:"AplblRate,omitempty"`

	// Rate of the cash premium made available if the securities holder consents or participates to an event, for example consent fees or solicitation fee.
	SolicitationFeeRate *SolicitationFeeRateFormat7Choice `xml:"SlctnFeeRate,omitempty"`

	// Amount of money per equity allocated as the result of a tax credit.
	TaxCreditRate []*TaxCreditRateFormat8Choice `xml:"TaxCdtRate,omitempty"`

	// Overall tax withheld at source by fund managers prior to considering the tax obligation of each unit holder.
	TaxOnIncome *RateAndAmountFormat37Choice `xml:"TaxOnIncm,omitempty"`

	// Taxation applied on an amount clearly identified as capital profits, capital gains.
	TaxOnProfits *RateFormat3Choice `xml:"TaxOnPrfts,omitempty"`

	// Percentage of cash that was paid in excess of actual tax obligation and was reclaimed.
	TaxReclaimRate *RateFormat3Choice `xml:"TaxRclmRate,omitempty"`

	// Portion of the fund distribution which represents the average accrued income included in the purchase price for units bought during the account period.
	EqualisationRate *RateAndAmountFormat42Choice `xml:"EqulstnRate,omitempty"`
}

func (r *RateDetails22) AddAdditionalTax() *RateAndAmountFormat37Choice {
	r.AdditionalTax = new(RateAndAmountFormat37Choice)
	return r.AdditionalTax
}

func (r *RateDetails22) AddGrossDividendRate() *GrossDividendRateFormat20Choice {
	newValue := new(GrossDividendRateFormat20Choice)
	r.GrossDividendRate = append(r.GrossDividendRate, newValue)
	return newValue
}

func (r *RateDetails22) AddInterestRateUsedForPayment() *InterestRateUsedForPaymentFormat8Choice {
	newValue := new(InterestRateUsedForPaymentFormat8Choice)
	r.InterestRateUsedForPayment = append(r.InterestRateUsedForPayment, newValue)
	return newValue
}

func (r *RateDetails22) AddWithholdingTaxRate() *RateAndAmountFormat41Choice {
	newValue := new(RateAndAmountFormat41Choice)
	r.WithholdingTaxRate = append(r.WithholdingTaxRate, newValue)
	return newValue
}

func (r *RateDetails22) AddSecondLevelTax() *RateAndAmountFormat41Choice {
	newValue := new(RateAndAmountFormat41Choice)
	r.SecondLevelTax = append(r.SecondLevelTax, newValue)
	return newValue
}

func (r *RateDetails22) AddChargesFees() *RateAndAmountFormat37Choice {
	r.ChargesFees = new(RateAndAmountFormat37Choice)
	return r.ChargesFees
}

func (r *RateDetails22) AddEarlySolicitationFeeRate() *SolicitationFeeRateFormat7Choice {
	r.EarlySolicitationFeeRate = new(SolicitationFeeRateFormat7Choice)
	return r.EarlySolicitationFeeRate
}

func (r *RateDetails22) AddFiscalStamp() *RateFormat3Choice {
	r.FiscalStamp = new(RateFormat3Choice)
	return r.FiscalStamp
}

func (r *RateDetails22) AddFullyFrankedRate() *RateAndAmountFormat37Choice {
	r.FullyFrankedRate = new(RateAndAmountFormat37Choice)
	return r.FullyFrankedRate
}

func (r *RateDetails22) AddThirdPartyIncentiveRate() *RateFormat20Choice {
	r.ThirdPartyIncentiveRate = new(RateFormat20Choice)
	return r.ThirdPartyIncentiveRate
}

func (r *RateDetails22) AddNetDividendRate() *NetDividendRateFormat22Choice {
	newValue := new(NetDividendRateFormat22Choice)
	r.NetDividendRate = append(r.NetDividendRate, newValue)
	return newValue
}

func (r *RateDetails22) AddNonResidentRate() *RateAndAmountFormat37Choice {
	r.NonResidentRate = new(RateAndAmountFormat37Choice)
	return r.NonResidentRate
}

func (r *RateDetails22) AddApplicableRate() *RateFormat3Choice {
	r.ApplicableRate = new(RateFormat3Choice)
	return r.ApplicableRate
}

func (r *RateDetails22) AddSolicitationFeeRate() *SolicitationFeeRateFormat7Choice {
	r.SolicitationFeeRate = new(SolicitationFeeRateFormat7Choice)
	return r.SolicitationFeeRate
}

func (r *RateDetails22) AddTaxCreditRate() *TaxCreditRateFormat8Choice {
	newValue := new(TaxCreditRateFormat8Choice)
	r.TaxCreditRate = append(r.TaxCreditRate, newValue)
	return newValue
}

func (r *RateDetails22) AddTaxOnIncome() *RateAndAmountFormat37Choice {
	r.TaxOnIncome = new(RateAndAmountFormat37Choice)
	return r.TaxOnIncome
}

func (r *RateDetails22) AddTaxOnProfits() *RateFormat3Choice {
	r.TaxOnProfits = new(RateFormat3Choice)
	return r.TaxOnProfits
}

func (r *RateDetails22) AddTaxReclaimRate() *RateFormat3Choice {
	r.TaxReclaimRate = new(RateFormat3Choice)
	return r.TaxReclaimRate
}

func (r *RateDetails22) AddEqualisationRate() *RateAndAmountFormat42Choice {
	r.EqualisationRate = new(RateAndAmountFormat42Choice)
	return r.EqualisationRate
}
