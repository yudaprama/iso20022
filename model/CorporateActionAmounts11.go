package model

// Specifies amounts in the framework of a corporate action event.
type CorporateActionAmounts11 struct {

	// Amount of money that is to be/was posted to the account.
	PostingAmount *ActiveCurrencyAndAmount `xml:"PstngAmt"`

	// Amount of money before any deductions and allowances have been made.
	GrossCashAmount *ActiveCurrencyAndAmount `xml:"GrssCshAmt,omitempty"`

	// Amount of money after deductions and allowances have been made, if any, that is, the total amount +/- charges/fees.
	NetCashAmount *ActiveCurrencyAndAmount `xml:"NetCshAmt,omitempty"`

	// Amount of cash premium made available in order to encourage participation in the offer. Payment is made to a third party who has solicited an entity to take part in the offer.
	SolicitationFees *ActiveCurrencyAndAmount `xml:"SlctnFees,omitempty"`

	// Cash disbursement in lieu of a fractional quantity of, for example, equity.
	CashInLieuOfShare *ActiveCurrencyAndAmount `xml:"CshInLieuOfShr,omitempty"`

	// Amount of money distributed as the result of a capital gain.
	CapitalGain *ActiveCurrencyAndAmount `xml:"CptlGn,omitempty"`

	// Amount of money representing a coupon payment.
	InterestAmount *ActiveCurrencyAndAmount `xml:"IntrstAmt,omitempty"`

	// Amount of money resulting from a market claim.
	MarketClaimAmount *ActiveCurrencyAndAmount `xml:"MktClmAmt,omitempty"`

	// (Unique to France) Amount due to a buyer of securities dealt prior to ex date which may be subject to different rate of taxation.
	IndemnityAmount *ActiveCurrencyAndAmount `xml:"IndmntyAmt,omitempty"`

	// Amount of money that the borrower pays to the lender as a compensation. It does not entitle the lender to reclaim any tax credit and is sometimes treated differently by the local tax authorities of the lender. Also covers compensation/indemnity of missed dividend concerning early/late settlements if applicable to a market.
	ManufacturedDividendPaymentAmount *ActiveCurrencyAndAmount `xml:"ManfctrdDvddPmtAmt,omitempty"`

	// Amount of money reinvested in additional securities.
	ReinvestmentAmount *ActiveCurrencyAndAmount `xml:"RinvstmtAmt,omitempty"`

	// Amount resulting from a fully franked dividend paid by a company; amount includes tax credit for companies that have made sufficient tax payments during the fiscal period.
	FullyFrankedAmount *ActiveCurrencyAndAmount `xml:"FullyFrnkdAmt,omitempty"`

	// Amount resulting from an unfranked dividend paid by a company; the amount does not include tax credit and is subject to withholding tax.
	UnfrankedAmount *ActiveCurrencyAndAmount `xml:"UfrnkdAmt,omitempty"`

	// Amount of money related to taxable income that cannot be categorised.
	SundryOrOtherAmount *ActiveCurrencyAndAmount `xml:"SndryOrOthrAmt,omitempty"`

	// Cash premium made available if the securities holder consents or participates to an event, for example, consent fees or solicitation fees.
	CashIncentive *ActiveCurrencyAndAmount `xml:"CshIncntiv,omitempty"`

	// Amount of money that has not been subject to taxation.
	TaxFreeAmount *ActiveCurrencyAndAmount `xml:"TaxFreeAmt,omitempty"`

	// Amount of income eligible for deferred taxation.
	TaxDeferredAmount *ActiveCurrencyAndAmount `xml:"TaxDfrrdAmt,omitempty"`

	// Tax on value added.
	ValueAddedTaxAmount *ActiveCurrencyAndAmount `xml:"ValAddedTaxAmt,omitempty"`

	// Amount of stamp duty.
	StampDutyAmount *ActiveCurrencyAndAmount `xml:"StmpDtyAmt,omitempty"`

	// Amount that was paid in excess of actual tax obligation and was reclaimed.
	TaxReclaimAmount *ActiveCurrencyAndAmount `xml:"TaxRclmAmt,omitempty"`

	// Amount of taxes that have been previously paid in relation to the taxable event.
	TaxCreditAmount *ActiveCurrencyAndAmount `xml:"TaxCdtAmt,omitempty"`

	// Amount of money withheld by the jurisdiction in which the income was originally paid, for which relief at source and/or reclaim may be possible.
	WithholdingOfForeignTaxAmount *ActiveCurrencyAndAmount `xml:"WhldgOfFrgnTaxAmt,omitempty"`

	// Amount of money withheld by the jurisdiction in which the account owner is located, for which relief at source and/or reclaim may be possible.
	WithholdingOfLocalTaxAmount *ActiveCurrencyAndAmount `xml:"WhldgOfLclTaxAmt,omitempty"`

	// Amount of additional taxes that cannot be categorised.
	AdditionalTaxAmount *ActiveCurrencyAndAmount `xml:"AddtlTaxAmt,omitempty"`

	// Amount of a cash distribution that will be withheld by a tax authority.
	WithholdingTaxAmount *ActiveCurrencyAndAmount `xml:"WhldgTaxAmt,omitempty"`

	// Amount of fiscal tax to apply.
	FiscalStampAmount *ActiveCurrencyAndAmount `xml:"FsclStmpAmt,omitempty"`

	// Amount of money paid to an executing broker as a commission.
	ExecutingBrokerAmount *ActiveCurrencyAndAmount `xml:"ExctgBrkrAmt,omitempty"`

	// Amount of paying/sub-paying agent commission.
	PayingAgentCommissionAmount *ActiveCurrencyAndAmount `xml:"PngAgtComssnAmt,omitempty"`

	// Local broker's commission.
	LocalBrokerCommissionAmount *ActiveCurrencyAndAmount `xml:"LclBrkrComssnAmt,omitempty"`

	// Amount of money paid for delivery by regular post mail.
	PostageFeeAmount *ActiveCurrencyAndAmount `xml:"PstgFeeAmt,omitempty"`

	// Amount of money charged by a regulatory authority, for example, Securities and Exchange fees.
	RegulatoryFeesAmount *ActiveCurrencyAndAmount `xml:"RgltryFeesAmt,omitempty"`

	// Amount of money (including insurance) paid for delivery by carrier.
	ShippingFeesAmount *ActiveCurrencyAndAmount `xml:"ShppgFeesAmt,omitempty"`

	// Amount of money paid for the provision of financial services that cannot be categorised by another qualifier.
	ChargesAmount *ActiveCurrencyAndAmount `xml:"ChrgsAmt,omitempty"`

	// Indicates cash retained from previous dividend.
	CashAmountBroughtForward *ActiveCurrencyAndAmount `xml:"CshAmtBrghtFwd,omitempty"`

	// Indicates the balance carried forward to next dividend.
	CashAmountCarriedForward *ActiveCurrencyAndAmount `xml:"CshAmtCrrdFwd,omitempty"`

	// Amount of cash that would have been payable if the dividend had been taken in the form of cash rather than shares.
	NotionalDividendPayableAmount *ActiveCurrencyAndAmount `xml:"NtnlDvddPyblAmt,omitempty"`

	// Where new securities are issued in lieu of a cash dividend, the notional tax is the tax on the amount of cash that would have been paid.
	NotionalTaxAmount *ActiveCurrencyAndAmount `xml:"NtnlTaxAmt,omitempty"`

	// Amount of money paid by the Tax Authorities in addition to the payment of the tax refund itself.
	TaxArrearsAmount *ActiveCurrencyAndAmount `xml:"TaxArrearsAmt,omitempty"`

	// Posting/settlement Amount in its original currency when conversion from/into another currency has occurred.
	OriginalAmount *ActiveCurrencyAndAmount `xml:"OrgnlAmt,omitempty"`

	// Amount of money representing a distribution of a bond's principal, for example, repayment of outstanding debt.
	PrincipalOrCorpus *ActiveCurrencyAndAmount `xml:"PrncplOrCrps,omitempty"`

	// Amount of money (not interest) in addition to the principal at the redemption of a bond.
	RedemptionPremiumAmount *ActiveCurrencyAndAmount `xml:"RedPrmAmt,omitempty"`

	// Amount relating to the underlying security for which income is distributed.
	IncomePortion *ActiveCurrencyAndAmount `xml:"IncmPrtn,omitempty"`

	// Amount of stock exchange tax.
	StockExchangeTax *ActiveCurrencyAndAmount `xml:"StockXchgTax,omitempty"`

	// Total amount of tax withheld at source in conformance with the EU Savings Directive.
	EUTaxRetentionAmount *ActiveCurrencyAndAmount `xml:"EUTaxRtntnAmt,omitempty"`

	// Amount of interest that has been accrued in between coupon payment periods
	AccruedInterestAmount *ActiveCurrencyAndAmount `xml:"AcrdIntrstAmt,omitempty"`
}

func (c *CorporateActionAmounts11) SetPostingAmount(value, currency string) {
	c.PostingAmount = NewActiveCurrencyAndAmount(value, currency)
}

func (c *CorporateActionAmounts11) SetGrossCashAmount(value, currency string) {
	c.GrossCashAmount = NewActiveCurrencyAndAmount(value, currency)
}

func (c *CorporateActionAmounts11) SetNetCashAmount(value, currency string) {
	c.NetCashAmount = NewActiveCurrencyAndAmount(value, currency)
}

func (c *CorporateActionAmounts11) SetSolicitationFees(value, currency string) {
	c.SolicitationFees = NewActiveCurrencyAndAmount(value, currency)
}

func (c *CorporateActionAmounts11) SetCashInLieuOfShare(value, currency string) {
	c.CashInLieuOfShare = NewActiveCurrencyAndAmount(value, currency)
}

func (c *CorporateActionAmounts11) SetCapitalGain(value, currency string) {
	c.CapitalGain = NewActiveCurrencyAndAmount(value, currency)
}

func (c *CorporateActionAmounts11) SetInterestAmount(value, currency string) {
	c.InterestAmount = NewActiveCurrencyAndAmount(value, currency)
}

func (c *CorporateActionAmounts11) SetMarketClaimAmount(value, currency string) {
	c.MarketClaimAmount = NewActiveCurrencyAndAmount(value, currency)
}

func (c *CorporateActionAmounts11) SetIndemnityAmount(value, currency string) {
	c.IndemnityAmount = NewActiveCurrencyAndAmount(value, currency)
}

func (c *CorporateActionAmounts11) SetManufacturedDividendPaymentAmount(value, currency string) {
	c.ManufacturedDividendPaymentAmount = NewActiveCurrencyAndAmount(value, currency)
}

func (c *CorporateActionAmounts11) SetReinvestmentAmount(value, currency string) {
	c.ReinvestmentAmount = NewActiveCurrencyAndAmount(value, currency)
}

func (c *CorporateActionAmounts11) SetFullyFrankedAmount(value, currency string) {
	c.FullyFrankedAmount = NewActiveCurrencyAndAmount(value, currency)
}

func (c *CorporateActionAmounts11) SetUnfrankedAmount(value, currency string) {
	c.UnfrankedAmount = NewActiveCurrencyAndAmount(value, currency)
}

func (c *CorporateActionAmounts11) SetSundryOrOtherAmount(value, currency string) {
	c.SundryOrOtherAmount = NewActiveCurrencyAndAmount(value, currency)
}

func (c *CorporateActionAmounts11) SetCashIncentive(value, currency string) {
	c.CashIncentive = NewActiveCurrencyAndAmount(value, currency)
}

func (c *CorporateActionAmounts11) SetTaxFreeAmount(value, currency string) {
	c.TaxFreeAmount = NewActiveCurrencyAndAmount(value, currency)
}

func (c *CorporateActionAmounts11) SetTaxDeferredAmount(value, currency string) {
	c.TaxDeferredAmount = NewActiveCurrencyAndAmount(value, currency)
}

func (c *CorporateActionAmounts11) SetValueAddedTaxAmount(value, currency string) {
	c.ValueAddedTaxAmount = NewActiveCurrencyAndAmount(value, currency)
}

func (c *CorporateActionAmounts11) SetStampDutyAmount(value, currency string) {
	c.StampDutyAmount = NewActiveCurrencyAndAmount(value, currency)
}

func (c *CorporateActionAmounts11) SetTaxReclaimAmount(value, currency string) {
	c.TaxReclaimAmount = NewActiveCurrencyAndAmount(value, currency)
}

func (c *CorporateActionAmounts11) SetTaxCreditAmount(value, currency string) {
	c.TaxCreditAmount = NewActiveCurrencyAndAmount(value, currency)
}

func (c *CorporateActionAmounts11) SetWithholdingOfForeignTaxAmount(value, currency string) {
	c.WithholdingOfForeignTaxAmount = NewActiveCurrencyAndAmount(value, currency)
}

func (c *CorporateActionAmounts11) SetWithholdingOfLocalTaxAmount(value, currency string) {
	c.WithholdingOfLocalTaxAmount = NewActiveCurrencyAndAmount(value, currency)
}

func (c *CorporateActionAmounts11) SetAdditionalTaxAmount(value, currency string) {
	c.AdditionalTaxAmount = NewActiveCurrencyAndAmount(value, currency)
}

func (c *CorporateActionAmounts11) SetWithholdingTaxAmount(value, currency string) {
	c.WithholdingTaxAmount = NewActiveCurrencyAndAmount(value, currency)
}

func (c *CorporateActionAmounts11) SetFiscalStampAmount(value, currency string) {
	c.FiscalStampAmount = NewActiveCurrencyAndAmount(value, currency)
}

func (c *CorporateActionAmounts11) SetExecutingBrokerAmount(value, currency string) {
	c.ExecutingBrokerAmount = NewActiveCurrencyAndAmount(value, currency)
}

func (c *CorporateActionAmounts11) SetPayingAgentCommissionAmount(value, currency string) {
	c.PayingAgentCommissionAmount = NewActiveCurrencyAndAmount(value, currency)
}

func (c *CorporateActionAmounts11) SetLocalBrokerCommissionAmount(value, currency string) {
	c.LocalBrokerCommissionAmount = NewActiveCurrencyAndAmount(value, currency)
}

func (c *CorporateActionAmounts11) SetPostageFeeAmount(value, currency string) {
	c.PostageFeeAmount = NewActiveCurrencyAndAmount(value, currency)
}

func (c *CorporateActionAmounts11) SetRegulatoryFeesAmount(value, currency string) {
	c.RegulatoryFeesAmount = NewActiveCurrencyAndAmount(value, currency)
}

func (c *CorporateActionAmounts11) SetShippingFeesAmount(value, currency string) {
	c.ShippingFeesAmount = NewActiveCurrencyAndAmount(value, currency)
}

func (c *CorporateActionAmounts11) SetChargesAmount(value, currency string) {
	c.ChargesAmount = NewActiveCurrencyAndAmount(value, currency)
}

func (c *CorporateActionAmounts11) SetCashAmountBroughtForward(value, currency string) {
	c.CashAmountBroughtForward = NewActiveCurrencyAndAmount(value, currency)
}

func (c *CorporateActionAmounts11) SetCashAmountCarriedForward(value, currency string) {
	c.CashAmountCarriedForward = NewActiveCurrencyAndAmount(value, currency)
}

func (c *CorporateActionAmounts11) SetNotionalDividendPayableAmount(value, currency string) {
	c.NotionalDividendPayableAmount = NewActiveCurrencyAndAmount(value, currency)
}

func (c *CorporateActionAmounts11) SetNotionalTaxAmount(value, currency string) {
	c.NotionalTaxAmount = NewActiveCurrencyAndAmount(value, currency)
}

func (c *CorporateActionAmounts11) SetTaxArrearsAmount(value, currency string) {
	c.TaxArrearsAmount = NewActiveCurrencyAndAmount(value, currency)
}

func (c *CorporateActionAmounts11) SetOriginalAmount(value, currency string) {
	c.OriginalAmount = NewActiveCurrencyAndAmount(value, currency)
}

func (c *CorporateActionAmounts11) SetPrincipalOrCorpus(value, currency string) {
	c.PrincipalOrCorpus = NewActiveCurrencyAndAmount(value, currency)
}

func (c *CorporateActionAmounts11) SetRedemptionPremiumAmount(value, currency string) {
	c.RedemptionPremiumAmount = NewActiveCurrencyAndAmount(value, currency)
}

func (c *CorporateActionAmounts11) SetIncomePortion(value, currency string) {
	c.IncomePortion = NewActiveCurrencyAndAmount(value, currency)
}

func (c *CorporateActionAmounts11) SetStockExchangeTax(value, currency string) {
	c.StockExchangeTax = NewActiveCurrencyAndAmount(value, currency)
}

func (c *CorporateActionAmounts11) SetEUTaxRetentionAmount(value, currency string) {
	c.EUTaxRetentionAmount = NewActiveCurrencyAndAmount(value, currency)
}

func (c *CorporateActionAmounts11) SetAccruedInterestAmount(value, currency string) {
	c.AccruedInterestAmount = NewActiveCurrencyAndAmount(value, currency)
}
