package model

// Provides information about the rates related to securities movement.
type RateDetails25 struct {

	// Rate used for additional tax that cannot be categorised.
	AdditionalTax *RateAndAmountFormat46Choice `xml:"AddtlTax,omitempty"`

	// Cash dividend amount per equity before deductions or allowances have been made.
	GrossDividendRate []*GrossDividendRateFormat26Choice `xml:"GrssDvddRate,omitempty"`

	// The actual interest rate used for the payment of the interest for the specified interest period.
	// Usage guideline: It is used to provide the applicable rate for the current payment, after all calculations have been performed, that is, application of period and method of interest computation.
	InterestRateUsedForPayment []*InterestRateUsedForPaymentFormat10Choice `xml:"IntrstRateUsdForPmt,omitempty"`

	// Percentage of a cash distribution that will be withheld by the tax authorities of the jurisdiction of the issuer, for which a relief at source and/or reclaim may be possible.
	WithholdingTaxRate []*RateAndAmountFormat47Choice `xml:"WhldgTaxRate,omitempty"`

	// Rate at which the income will be withheld by a jurisdiction other than the jurisdiction of the issuer’s country of tax incorporation, for which a relief at source and/or reclaim may be possible. It is levied in complement or offset of the withholding tax rate (TAXR) levied by the jurisdiction of the issuer’s tax domicile.
	SecondLevelTax []*RateAndAmountFormat47Choice `xml:"ScndLvlTax,omitempty"`

	// Rate used to calculate the amount of the charges/fees that cannot be categorised.
	ChargesFees *RateAndAmountFormat46Choice `xml:"ChrgsFees,omitempty"`

	// Cash rate made available, as an incentive, in addition to the solicitation fee, in order to encourage early participation in an offer.
	EarlySolicitationFeeRate *SolicitationFeeRateFormat10Choice `xml:"EarlySlctnFeeRate,omitempty"`

	// Percentage of fiscal tax to apply.
	FiscalStamp *RateFormat3Choice `xml:"FsclStmp,omitempty"`

	// Rate resulting from a fully franked dividend paid by a company; rate includes tax credit for companies that have made sufficient tax payments during fiscal period.
	FullyFrankedRate *RateAndAmountFormat46Choice `xml:"FullyFrnkdRate,omitempty"`

	// Cash rate made available in an event in order to encourage participation in the offer. As information, payment is made to a third party who has solicited an entity to take part in the offer.
	ThirdPartyIncentiveRate *RateFormat21Choice `xml:"ThrdPtyIncntivRate,omitempty"`

	// Cash dividend amount per equity after deductions or allowances have been made.
	NetDividendRate []*NetDividendRateFormat28Choice `xml:"NetDvddRate,omitempty"`

	// Rate per share to which a non-resident is entitled.
	NonResidentRate *RateAndAmountFormat46Choice `xml:"NonResdtRate,omitempty"`

	// Rate applicable to the event announced, for example, redemption rate for a redemption event.
	ApplicableRate *RateFormat3Choice `xml:"AplblRate,omitempty"`

	// Rate of the cash premium made available if the securities holder consents or participates to an event, for example consent fees or solicitation fee.
	SolicitationFeeRate *SolicitationFeeRateFormat10Choice `xml:"SlctnFeeRate,omitempty"`

	// Amount of money per equity allocated as the result of a tax credit.
	TaxCreditRate []*TaxCreditRateFormat10Choice `xml:"TaxCdtRate,omitempty"`

	// Overall tax withheld at source by fund managers prior to considering the tax obligation of each unit holder.
	TaxOnIncome *RateAndAmountFormat46Choice `xml:"TaxOnIncm,omitempty"`

	// Taxation applied on an amount clearly identified as capital profits, capital gains.
	TaxOnProfits *RateFormat3Choice `xml:"TaxOnPrfts,omitempty"`

	// Percentage of cash that was paid in excess of actual tax obligation and was reclaimed.
	TaxReclaimRate *RateFormat3Choice `xml:"TaxRclmRate,omitempty"`

	// Portion of the fund distribution which represents the average accrued income included in the purchase price for units bought during the account period.
	EqualisationRate *RateAndAmountFormat48Choice `xml:"EqulstnRate,omitempty"`
}

func (r *RateDetails25) AddAdditionalTax() *RateAndAmountFormat46Choice {
	r.AdditionalTax = new(RateAndAmountFormat46Choice)
	return r.AdditionalTax
}

func (r *RateDetails25) AddGrossDividendRate() *GrossDividendRateFormat26Choice {
	newValue := new(GrossDividendRateFormat26Choice)
	r.GrossDividendRate = append(r.GrossDividendRate, newValue)
	return newValue
}

func (r *RateDetails25) AddInterestRateUsedForPayment() *InterestRateUsedForPaymentFormat10Choice {
	newValue := new(InterestRateUsedForPaymentFormat10Choice)
	r.InterestRateUsedForPayment = append(r.InterestRateUsedForPayment, newValue)
	return newValue
}

func (r *RateDetails25) AddWithholdingTaxRate() *RateAndAmountFormat47Choice {
	newValue := new(RateAndAmountFormat47Choice)
	r.WithholdingTaxRate = append(r.WithholdingTaxRate, newValue)
	return newValue
}

func (r *RateDetails25) AddSecondLevelTax() *RateAndAmountFormat47Choice {
	newValue := new(RateAndAmountFormat47Choice)
	r.SecondLevelTax = append(r.SecondLevelTax, newValue)
	return newValue
}

func (r *RateDetails25) AddChargesFees() *RateAndAmountFormat46Choice {
	r.ChargesFees = new(RateAndAmountFormat46Choice)
	return r.ChargesFees
}

func (r *RateDetails25) AddEarlySolicitationFeeRate() *SolicitationFeeRateFormat10Choice {
	r.EarlySolicitationFeeRate = new(SolicitationFeeRateFormat10Choice)
	return r.EarlySolicitationFeeRate
}

func (r *RateDetails25) AddFiscalStamp() *RateFormat3Choice {
	r.FiscalStamp = new(RateFormat3Choice)
	return r.FiscalStamp
}

func (r *RateDetails25) AddFullyFrankedRate() *RateAndAmountFormat46Choice {
	r.FullyFrankedRate = new(RateAndAmountFormat46Choice)
	return r.FullyFrankedRate
}

func (r *RateDetails25) AddThirdPartyIncentiveRate() *RateFormat21Choice {
	r.ThirdPartyIncentiveRate = new(RateFormat21Choice)
	return r.ThirdPartyIncentiveRate
}

func (r *RateDetails25) AddNetDividendRate() *NetDividendRateFormat28Choice {
	newValue := new(NetDividendRateFormat28Choice)
	r.NetDividendRate = append(r.NetDividendRate, newValue)
	return newValue
}

func (r *RateDetails25) AddNonResidentRate() *RateAndAmountFormat46Choice {
	r.NonResidentRate = new(RateAndAmountFormat46Choice)
	return r.NonResidentRate
}

func (r *RateDetails25) AddApplicableRate() *RateFormat3Choice {
	r.ApplicableRate = new(RateFormat3Choice)
	return r.ApplicableRate
}

func (r *RateDetails25) AddSolicitationFeeRate() *SolicitationFeeRateFormat10Choice {
	r.SolicitationFeeRate = new(SolicitationFeeRateFormat10Choice)
	return r.SolicitationFeeRate
}

func (r *RateDetails25) AddTaxCreditRate() *TaxCreditRateFormat10Choice {
	newValue := new(TaxCreditRateFormat10Choice)
	r.TaxCreditRate = append(r.TaxCreditRate, newValue)
	return newValue
}

func (r *RateDetails25) AddTaxOnIncome() *RateAndAmountFormat46Choice {
	r.TaxOnIncome = new(RateAndAmountFormat46Choice)
	return r.TaxOnIncome
}

func (r *RateDetails25) AddTaxOnProfits() *RateFormat3Choice {
	r.TaxOnProfits = new(RateFormat3Choice)
	return r.TaxOnProfits
}

func (r *RateDetails25) AddTaxReclaimRate() *RateFormat3Choice {
	r.TaxReclaimRate = new(RateFormat3Choice)
	return r.TaxReclaimRate
}

func (r *RateDetails25) AddEqualisationRate() *RateAndAmountFormat48Choice {
	r.EqualisationRate = new(RateAndAmountFormat48Choice)
	return r.EqualisationRate
}
