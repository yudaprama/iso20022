package model

// Provides information about the rates related to securities movement.
type RateDetails3 struct {

	// Rate used for additional tax that cannot be categorised.
	AdditionalTax *RateAndAmountFormat14Choice `xml:"AddtlTax,omitempty"`

	// Cash dividend amount per equity before deductions or allowances have been made.
	GrossDividendRate []*GrossDividendRateFormat5Choice `xml:"GrssDvddRate,omitempty"`

	// The actual interest rate used for the payment of the interest for the specified interest period.
	// Usage guideline: It is used to provide the applicable rate for the current payment, after all calculations have been performed, that is, application of period and method of interest computation.
	InterestRateUsedForPayment []*InterestRateUsedForPaymentFormat5Choice `xml:"IntrstRateUsdForPmt,omitempty"`

	// Percentage of the gross dividend rate on which tax must be paid .
	TaxRelatedRate []*RateTypeAndAmountAndStatus6 `xml:"TaxRltdRate,omitempty"`

	// Percentage of a cash distribution that will be withheld by a tax authority.
	WithholdingTaxRate *RateFormat6Choice `xml:"WhldgTaxRate,omitempty"`

	// Rate used to calculate the amount of the charges/fees that cannot be categorised.
	ChargesFees *RateAndAmountFormat14Choice `xml:"ChrgsFees,omitempty"`

	// Cash rate made available, as an incentive, in addition to the solicitation fee, in order to encourage early participation in an offer.
	EarlySolicitationFeeRate *SolicitationFeeRateFormat3Choice `xml:"EarlySlctnFeeRate,omitempty"`

	// Dividend is final.
	FinalDividendRate *RateAndAmountFormat15Choice `xml:"FnlDvddRate,omitempty"`

	// Percentage of fiscal tax to apply.
	FiscalStamp *RateFormat6Choice `xml:"FsclStmp,omitempty"`

	// Rate resulting from a fully franked dividend paid by a company; rate includes tax credit for companies that have made sufficient tax payments during fiscal period.
	FullyFrankedRate *RateAndAmountFormat14Choice `xml:"FullyFrnkdRate,omitempty"`

	// Rate of the cash premium made available if the securities holder consents or participates to an event, for example, consent fees.
	CashIncentiveRate *RateFormat6Choice `xml:"CshIncntivRate,omitempty"`

	// Cash dividend amount per equity after deductions or allowances have been made.
	NetDividendRate []*NetDividendRateFormat5Choice `xml:"NetDvddRate,omitempty"`

	// Rate per share to which a non-resident is entitled.
	NonResidentRate *RateAndAmountFormat14Choice `xml:"NonResdtRate,omitempty"`

	// Dividend is provisional.
	ProvisionalDividendRate *RateAndAmountFormat15Choice `xml:"PrvsnlDvddRate,omitempty"`

	// Rate applicable to the event announced, for example, redemption rate for a redemption event.
	ApplicableRate *RateFormat6Choice `xml:"AplblRate,omitempty"`

	// Cash rate made available in an offer in order to encourage participation in the offer.
	SolicitationFeeRate *SolicitationFeeRateFormat3Choice `xml:"SlctnFeeRate,omitempty"`

	// Amount of money per equity allocated as the result of a tax credit.
	TaxCreditRate []*TaxCreditRateFormat5Choice `xml:"TaxCdtRate,omitempty"`

	// Taxation applied on an amount clearly identified as an income.
	TaxOnIncome *RateFormat6Choice `xml:"TaxOnIncm,omitempty"`

	// Taxation applied on an amount clearly identified as capital profits, capital gains.
	TaxOnProfits *RateFormat6Choice `xml:"TaxOnPrfts,omitempty"`

	// Percentage of cash that was paid in excess of actual tax obligation and was reclaimed.
	TaxReclaimRate *RateFormat6Choice `xml:"TaxRclmRate,omitempty"`

	// Rate at which the income will be withheld by the jurisdiction in which the income was originally paid, for which relief at source and/or reclaim may be possible.
	WithholdingOfForeignTax *RateAndAmountFormat14Choice `xml:"WhldgOfFrgnTax,omitempty"`

	// Rate at which the income will be withheld by the jurisdiction in which the account owner is located, for which relief at source and/or reclaim may be possible.
	WithholdingOfLocalTax *RateAndAmountFormat14Choice `xml:"WhldgOfLclTax,omitempty"`
}

func (r *RateDetails3) AddAdditionalTax() *RateAndAmountFormat14Choice {
	r.AdditionalTax = new(RateAndAmountFormat14Choice)
	return r.AdditionalTax
}

func (r *RateDetails3) AddGrossDividendRate() *GrossDividendRateFormat5Choice {
	newValue := new(GrossDividendRateFormat5Choice)
	r.GrossDividendRate = append(r.GrossDividendRate, newValue)
	return newValue
}

func (r *RateDetails3) AddInterestRateUsedForPayment() *InterestRateUsedForPaymentFormat5Choice {
	newValue := new(InterestRateUsedForPaymentFormat5Choice)
	r.InterestRateUsedForPayment = append(r.InterestRateUsedForPayment, newValue)
	return newValue
}

func (r *RateDetails3) AddTaxRelatedRate() *RateTypeAndAmountAndStatus6 {
	newValue := new(RateTypeAndAmountAndStatus6)
	r.TaxRelatedRate = append(r.TaxRelatedRate, newValue)
	return newValue
}

func (r *RateDetails3) AddWithholdingTaxRate() *RateFormat6Choice {
	r.WithholdingTaxRate = new(RateFormat6Choice)
	return r.WithholdingTaxRate
}

func (r *RateDetails3) AddChargesFees() *RateAndAmountFormat14Choice {
	r.ChargesFees = new(RateAndAmountFormat14Choice)
	return r.ChargesFees
}

func (r *RateDetails3) AddEarlySolicitationFeeRate() *SolicitationFeeRateFormat3Choice {
	r.EarlySolicitationFeeRate = new(SolicitationFeeRateFormat3Choice)
	return r.EarlySolicitationFeeRate
}

func (r *RateDetails3) AddFinalDividendRate() *RateAndAmountFormat15Choice {
	r.FinalDividendRate = new(RateAndAmountFormat15Choice)
	return r.FinalDividendRate
}

func (r *RateDetails3) AddFiscalStamp() *RateFormat6Choice {
	r.FiscalStamp = new(RateFormat6Choice)
	return r.FiscalStamp
}

func (r *RateDetails3) AddFullyFrankedRate() *RateAndAmountFormat14Choice {
	r.FullyFrankedRate = new(RateAndAmountFormat14Choice)
	return r.FullyFrankedRate
}

func (r *RateDetails3) AddCashIncentiveRate() *RateFormat6Choice {
	r.CashIncentiveRate = new(RateFormat6Choice)
	return r.CashIncentiveRate
}

func (r *RateDetails3) AddNetDividendRate() *NetDividendRateFormat5Choice {
	newValue := new(NetDividendRateFormat5Choice)
	r.NetDividendRate = append(r.NetDividendRate, newValue)
	return newValue
}

func (r *RateDetails3) AddNonResidentRate() *RateAndAmountFormat14Choice {
	r.NonResidentRate = new(RateAndAmountFormat14Choice)
	return r.NonResidentRate
}

func (r *RateDetails3) AddProvisionalDividendRate() *RateAndAmountFormat15Choice {
	r.ProvisionalDividendRate = new(RateAndAmountFormat15Choice)
	return r.ProvisionalDividendRate
}

func (r *RateDetails3) AddApplicableRate() *RateFormat6Choice {
	r.ApplicableRate = new(RateFormat6Choice)
	return r.ApplicableRate
}

func (r *RateDetails3) AddSolicitationFeeRate() *SolicitationFeeRateFormat3Choice {
	r.SolicitationFeeRate = new(SolicitationFeeRateFormat3Choice)
	return r.SolicitationFeeRate
}

func (r *RateDetails3) AddTaxCreditRate() *TaxCreditRateFormat5Choice {
	newValue := new(TaxCreditRateFormat5Choice)
	r.TaxCreditRate = append(r.TaxCreditRate, newValue)
	return newValue
}

func (r *RateDetails3) AddTaxOnIncome() *RateFormat6Choice {
	r.TaxOnIncome = new(RateFormat6Choice)
	return r.TaxOnIncome
}

func (r *RateDetails3) AddTaxOnProfits() *RateFormat6Choice {
	r.TaxOnProfits = new(RateFormat6Choice)
	return r.TaxOnProfits
}

func (r *RateDetails3) AddTaxReclaimRate() *RateFormat6Choice {
	r.TaxReclaimRate = new(RateFormat6Choice)
	return r.TaxReclaimRate
}

func (r *RateDetails3) AddWithholdingOfForeignTax() *RateAndAmountFormat14Choice {
	r.WithholdingOfForeignTax = new(RateAndAmountFormat14Choice)
	return r.WithholdingOfForeignTax
}

func (r *RateDetails3) AddWithholdingOfLocalTax() *RateAndAmountFormat14Choice {
	r.WithholdingOfLocalTax = new(RateAndAmountFormat14Choice)
	return r.WithholdingOfLocalTax
}
