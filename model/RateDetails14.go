package model

// Provides information about the rates related to securities movement.
type RateDetails14 struct {

	// Rate used for additional tax that cannot be categorised.
	AdditionalTax *RateAndAmountFormat14Choice `xml:"AddtlTax,omitempty"`

	// Cash dividend amount per equity before deductions or allowances have been made.
	GrossDividendRate []*GrossDividendRateFormat7Choice `xml:"GrssDvddRate,omitempty"`

	// The actual interest rate used for the payment of the interest for the specified interest period.
	// Usage guideline: It is used to provide the applicable rate for the current payment, after all calculations have been performed, that is, application of period and method of interest computation.
	InterestRateUsedForPayment []*InterestRateUsedForPaymentFormat5Choice `xml:"IntrstRateUsdForPmt,omitempty"`

	// Percentage of the gross dividend rate on which tax must be paid .
	TaxRelatedRate []*RateTypeAndAmountAndStatus6 `xml:"TaxRltdRate,omitempty"`

	// Percentage of a cash distribution that will be withheld by a tax authority.
	WithholdingTaxRate []*RateFormat10Choice `xml:"WhldgTaxRate,omitempty"`

	// Rate used to calculate the amount of the charges/fees that cannot be categorised.
	ChargesFees *RateAndAmountFormat14Choice `xml:"ChrgsFees,omitempty"`

	// Cash rate made available, as an incentive, in addition to the solicitation fee, in order to encourage early participation in an offer.
	EarlySolicitationFeeRate *SolicitationFeeRateFormat5Choice `xml:"EarlySlctnFeeRate,omitempty"`

	// Dividend is final.
	FinalDividendRate *RateAndAmountFormat15Choice `xml:"FnlDvddRate,omitempty"`

	// Percentage of fiscal tax to apply.
	FiscalStamp *RateFormat3Choice `xml:"FsclStmp,omitempty"`

	// Rate resulting from a fully franked dividend paid by a company; rate includes tax credit for companies that have made sufficient tax payments during fiscal period.
	FullyFrankedRate *RateAndAmountFormat14Choice `xml:"FullyFrnkdRate,omitempty"`

	// Cash rate made available in an event in order to encourage participation in the offer. As information, payment is made to a third party who has solicited an entity to take part in the offer.
	ThirdPartyIncentiveRate *RateFormat8Choice `xml:"ThrdPtyIncntivRate,omitempty"`

	// Cash dividend amount per equity after deductions or allowances have been made.
	NetDividendRate []*NetDividendRateFormat9Choice `xml:"NetDvddRate,omitempty"`

	// Rate per share to which a non-resident is entitled.
	NonResidentRate *RateAndAmountFormat14Choice `xml:"NonResdtRate,omitempty"`

	// Dividend is provisional.
	ProvisionalDividendRate *RateAndAmountFormat15Choice `xml:"PrvsnlDvddRate,omitempty"`

	// Rate applicable to the event announced, for example, redemption rate for a redemption event.
	ApplicableRate *RateFormat3Choice `xml:"AplblRate,omitempty"`

	// Rate of the cash premium made available if the securities holder consents or participates to an event, for example consent fees or solicitation fee.
	SolicitationFeeRate *SolicitationFeeRateFormat5Choice `xml:"SlctnFeeRate,omitempty"`

	// Amount of money per equity allocated as the result of a tax credit.
	TaxCreditRate []*TaxCreditRateFormat5Choice `xml:"TaxCdtRate,omitempty"`

	// Overall tax withheld at source by fund managers prior to considering the tax obligation of each unit holder.
	TaxOnIncome *RateAndAmountFormat14Choice `xml:"TaxOnIncm,omitempty"`

	// Taxation applied on an amount clearly identified as capital profits, capital gains.
	TaxOnProfits *RateFormat3Choice `xml:"TaxOnPrfts,omitempty"`

	// Percentage of cash that was paid in excess of actual tax obligation and was reclaimed.
	TaxReclaimRate *RateFormat3Choice `xml:"TaxRclmRate,omitempty"`

	// Rate at which the income will be withheld by the jurisdiction in which the income was originally paid, for which relief at source and/or reclaim may be possible.
	WithholdingOfForeignTax []*RateAndAmountFormat20Choice `xml:"WhldgOfFrgnTax,omitempty"`

	// Rate at which the income will be withheld by the jurisdiction in which the account owner is located, for which relief at source and/or reclaim may be possible.
	WithholdingOfLocalTax []*RateAndAmountFormat20Choice `xml:"WhldgOfLclTax,omitempty"`

	// Portion of the fund distribution which represents the average accrued income included in the purchase price for units bought during the account period.
	EqualisationRate *RateAndAmountFormat15Choice `xml:"EqulstnRate,omitempty"`
}

func (r *RateDetails14) AddAdditionalTax() *RateAndAmountFormat14Choice {
	r.AdditionalTax = new(RateAndAmountFormat14Choice)
	return r.AdditionalTax
}

func (r *RateDetails14) AddGrossDividendRate() *GrossDividendRateFormat7Choice {
	newValue := new(GrossDividendRateFormat7Choice)
	r.GrossDividendRate = append(r.GrossDividendRate, newValue)
	return newValue
}

func (r *RateDetails14) AddInterestRateUsedForPayment() *InterestRateUsedForPaymentFormat5Choice {
	newValue := new(InterestRateUsedForPaymentFormat5Choice)
	r.InterestRateUsedForPayment = append(r.InterestRateUsedForPayment, newValue)
	return newValue
}

func (r *RateDetails14) AddTaxRelatedRate() *RateTypeAndAmountAndStatus6 {
	newValue := new(RateTypeAndAmountAndStatus6)
	r.TaxRelatedRate = append(r.TaxRelatedRate, newValue)
	return newValue
}

func (r *RateDetails14) AddWithholdingTaxRate() *RateFormat10Choice {
	newValue := new(RateFormat10Choice)
	r.WithholdingTaxRate = append(r.WithholdingTaxRate, newValue)
	return newValue
}

func (r *RateDetails14) AddChargesFees() *RateAndAmountFormat14Choice {
	r.ChargesFees = new(RateAndAmountFormat14Choice)
	return r.ChargesFees
}

func (r *RateDetails14) AddEarlySolicitationFeeRate() *SolicitationFeeRateFormat5Choice {
	r.EarlySolicitationFeeRate = new(SolicitationFeeRateFormat5Choice)
	return r.EarlySolicitationFeeRate
}

func (r *RateDetails14) AddFinalDividendRate() *RateAndAmountFormat15Choice {
	r.FinalDividendRate = new(RateAndAmountFormat15Choice)
	return r.FinalDividendRate
}

func (r *RateDetails14) AddFiscalStamp() *RateFormat3Choice {
	r.FiscalStamp = new(RateFormat3Choice)
	return r.FiscalStamp
}

func (r *RateDetails14) AddFullyFrankedRate() *RateAndAmountFormat14Choice {
	r.FullyFrankedRate = new(RateAndAmountFormat14Choice)
	return r.FullyFrankedRate
}

func (r *RateDetails14) AddThirdPartyIncentiveRate() *RateFormat8Choice {
	r.ThirdPartyIncentiveRate = new(RateFormat8Choice)
	return r.ThirdPartyIncentiveRate
}

func (r *RateDetails14) AddNetDividendRate() *NetDividendRateFormat9Choice {
	newValue := new(NetDividendRateFormat9Choice)
	r.NetDividendRate = append(r.NetDividendRate, newValue)
	return newValue
}

func (r *RateDetails14) AddNonResidentRate() *RateAndAmountFormat14Choice {
	r.NonResidentRate = new(RateAndAmountFormat14Choice)
	return r.NonResidentRate
}

func (r *RateDetails14) AddProvisionalDividendRate() *RateAndAmountFormat15Choice {
	r.ProvisionalDividendRate = new(RateAndAmountFormat15Choice)
	return r.ProvisionalDividendRate
}

func (r *RateDetails14) AddApplicableRate() *RateFormat3Choice {
	r.ApplicableRate = new(RateFormat3Choice)
	return r.ApplicableRate
}

func (r *RateDetails14) AddSolicitationFeeRate() *SolicitationFeeRateFormat5Choice {
	r.SolicitationFeeRate = new(SolicitationFeeRateFormat5Choice)
	return r.SolicitationFeeRate
}

func (r *RateDetails14) AddTaxCreditRate() *TaxCreditRateFormat5Choice {
	newValue := new(TaxCreditRateFormat5Choice)
	r.TaxCreditRate = append(r.TaxCreditRate, newValue)
	return newValue
}

func (r *RateDetails14) AddTaxOnIncome() *RateAndAmountFormat14Choice {
	r.TaxOnIncome = new(RateAndAmountFormat14Choice)
	return r.TaxOnIncome
}

func (r *RateDetails14) AddTaxOnProfits() *RateFormat3Choice {
	r.TaxOnProfits = new(RateFormat3Choice)
	return r.TaxOnProfits
}

func (r *RateDetails14) AddTaxReclaimRate() *RateFormat3Choice {
	r.TaxReclaimRate = new(RateFormat3Choice)
	return r.TaxReclaimRate
}

func (r *RateDetails14) AddWithholdingOfForeignTax() *RateAndAmountFormat20Choice {
	newValue := new(RateAndAmountFormat20Choice)
	r.WithholdingOfForeignTax = append(r.WithholdingOfForeignTax, newValue)
	return newValue
}

func (r *RateDetails14) AddWithholdingOfLocalTax() *RateAndAmountFormat20Choice {
	newValue := new(RateAndAmountFormat20Choice)
	r.WithholdingOfLocalTax = append(r.WithholdingOfLocalTax, newValue)
	return newValue
}

func (r *RateDetails14) AddEqualisationRate() *RateAndAmountFormat15Choice {
	r.EqualisationRate = new(RateAndAmountFormat15Choice)
	return r.EqualisationRate
}
