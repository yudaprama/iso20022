package model

// Specifies amounts in the framework of a corporate action event..
type CorporateActionAmounts1 struct {

	// Amount of money before any deductions and allowances have been made.
	GrossCashAmount *ActiveCurrencyAndAmount `xml:"GrssCshAmt,omitempty"`

	// Amount of money after deductions and allowances have been made, if any, ie, the total amount +/- charges/fees.
	NetCashAmount *ActiveCurrencyAndAmount `xml:"NetCshAmt,omitempty"`

	// Amount of money defined as a discount on a new issue or on a tranche of an existing issue.
	IssueDiscountAmount *ActiveCurrencyAndAmount `xml:"IsseDscntAmt,omitempty"`

	// Amount of cash premium made available in order to encourage participation in the offer. Payment is made to a third party who has solicited an entity to take part in the offer.
	SolicitationFees *ActiveCurrencyAndAmount `xml:"SlctnFees,omitempty"`

	// Cash disbursement in lieu of a fractional quantity of, eg, equity.
	CashInLieuOfShare *ActiveCurrencyAndAmount `xml:"CshInLieuOfShr,omitempty"`

	// Amount in its original currency when conversion from/into another currency has occurred.
	OriginalAmount *ActiveCurrencyAndAmount `xml:"OrgnlAmt,omitempty"`

	// Amount of money distributed as the result of a capital gain.
	CapitalGain *ActiveCurrencyAndAmount `xml:"CptlGn,omitempty"`

	// Amount of money representing a coupon payment.
	InterestAmount *ActiveCurrencyAndAmount `xml:"IntrstAmt,omitempty"`

	// (Unique to France) Amount due to a buyer of securities dealt prior to ex date which may be subject to different rate of taxation.
	IndemnityAmount *ActiveCurrencyAndAmount `xml:"IndmntyAmt,omitempty"`

	// Amount of money (not interest) in addition to the principal at the redemption of a bond.
	RedemptionPremiumAmount *ActiveCurrencyAndAmount `xml:"RedPrmAmt,omitempty"`

	// Amount of money that the borrower pays to the lender as a compensation. It does not entitle the lender to reclaim any tax credit and is sometimes treated differently by the local tax authorities of the lender.
	ManufacturedDividendAmount *ActiveCurrencyAndAmount `xml:"ManfctrdDvddAmt,omitempty"`

	// Amount of money representing a distribution of a bond's principal, eg, repayment of outstanding debt.
	PrincipalOrCorpus *ActiveCurrencyAndAmount `xml:"PrncplOrCrps,omitempty"`

	// Amount of money reinvested in additional securities.
	ReinvestmentAmount *ActiveCurrencyAndAmount `xml:"RinvstmtAmt,omitempty"`

	// Amount of money resulting from a market claim.
	MarketClaimAmount *ActiveCurrencyAndAmount `xml:"MktClmAmt,omitempty"`

	// Amount resulting from a fully franked dividend paid by a company; amount includes tax credit for companies that have made sufficient tax payments during the fiscal period.
	FullyFrankedAmount *ActiveCurrencyAndAmount `xml:"FullyFrnkdAmt,omitempty"`

	// Amount resulting from an unfranked dividend paid by a company; the amount does not include tax credit and is subject to withholding tax.
	UnfrankedAmount *ActiveCurrencyAndAmount `xml:"UfrnkdAmt,omitempty"`

	// Amount of money related to taxable income that cannot be categorised.
	SundryOrOtherAmount *ActiveCurrencyAndAmount `xml:"SndryOrOthrAmt,omitempty"`

	// Amount of drawdown or other reduction from or in addition to the deal price.
	SpecialConcessionAmount *ActiveCurrencyAndAmount `xml:"SpclCncssnAmt,omitempty"`

	// Cash amount based on terms of corporate action event and balance of underlying securities, entitled to/from account owner (which may be positive or negative).
	EntitledAmount *ActiveCurrencyAndAmount `xml:"EntitldAmt,omitempty"`

	// Rate of the cash premium made available if the securities holder consents or participates to an event, e.g. consent fees.
	CashIncentive *ActiveCurrencyAndAmount `xml:"CshIncntiv,omitempty"`

	// Additional costs - coming on top of the subscription costs - which the subscriber should pay as per the subscription process. Not to be used for the subscription cost itself.
	AdditionalSubscriptionCost *ActiveCurrencyAndAmount `xml:"AddtlSbcptCost,omitempty"`

	// Amount of money that has not been subject to taxation.
	TaxFreeAmount *ActiveCurrencyAndAmount `xml:"TaxFreeAmt,omitempty"`

	// Amount of income eligible for deferred taxation.
	TaxDeferredAmount *ActiveCurrencyAndAmount `xml:"TaxDfrrdAmt,omitempty"`

	// Local tax (ZAS Anrechnungsbetrag) subject to interest down payment tax (proportion of interest liable for interest down payment tax/interim profit that is not covered by the tax exempt amount).
	GermanLocalTax1Amount *ActiveCurrencyAndAmount `xml:"GrmnLclTax1Amt,omitempty"`

	// Local tax (ZAS Pflichtige Zinsen) interest liable for interest down payment tax (proportion of gross interest per unit/interim profits that is not covered by the credit in the interest pool).
	GermanLocalTax2Amount *ActiveCurrencyAndAmount `xml:"GrmnLclTax2Amt,omitempty"`

	// Local tax (Zinstopf) offset interest per unit against tax exempt amount (variation to offset interest per unit in relation to tax exempt amount).
	GermanLocalTax3Amount *ActiveCurrencyAndAmount `xml:"GrmnLclTax3Amt,omitempty"`

	// Local tax (Ertrag Besitzanteilig) yield liable for interest down payment tax.
	GermanLocalTax4Amount *ActiveCurrencyAndAmount `xml:"GrmnLclTax4Amt,omitempty"`

	// Amount of stock exchange tax.
	StockExchangeTaxAmount *ActiveCurrencyAndAmount `xml:"StockXchgTaxAmt,omitempty"`

	// Tax levied on a transfer of ownership of financial instrument.
	TransferTaxAmount *ActiveCurrencyAndAmount `xml:"TrfTaxAmt,omitempty"`

	// Amount of transaction tax.
	TransactionTaxAmount *ActiveCurrencyAndAmount `xml:"TxTaxAmt,omitempty"`

	// Tax on value added.
	ValueAddedTaxAmount *ActiveCurrencyAndAmount `xml:"ValAddedTaxAmt,omitempty"`

	// Total amount of tax withheld at source in conformance with the EU Savings Directive.
	EURetentionTaxAmount *ActiveCurrencyAndAmount `xml:"EURtntnTaxAmt,omitempty"`

	// Amount of tax charged by the jurisdiction in which the financial instrument settles.
	LocalTaxAmount *ActiveCurrencyAndAmount `xml:"LclTaxAmt,omitempty"`

	// Payment levy tax.
	PaymentLevyTaxAmount *ActiveCurrencyAndAmount `xml:"PmtLevyTaxAmt,omitempty"`

	// Amount of country, national or federal tax charged by the jurisdiction in which the account servicer is located.
	CountryNationalFederalTaxAmount *ActiveCurrencyAndAmount `xml:"CtryNtlFdrlTaxAmt,omitempty"`

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

	// Amount of money charged by a regulatory authority, eg, Securities and Exchange fees.
	RegulatoryFeesAmount *ActiveCurrencyAndAmount `xml:"RgltryFeesAmt,omitempty"`

	// Amount of money (including insurance) paid for delivery by carrier.
	ShippingFeesAmount *ActiveCurrencyAndAmount `xml:"ShppgFeesAmt,omitempty"`

	// Amount of money paid for the provision of financial services that cannot be categorised by another qualifier.
	ChargesAmount *ActiveCurrencyAndAmount `xml:"ChrgsAmt,omitempty"`
}

func (c *CorporateActionAmounts1) SetGrossCashAmount(value, currency string) {
	c.GrossCashAmount = NewActiveCurrencyAndAmount(value, currency)
}

func (c *CorporateActionAmounts1) SetNetCashAmount(value, currency string) {
	c.NetCashAmount = NewActiveCurrencyAndAmount(value, currency)
}

func (c *CorporateActionAmounts1) SetIssueDiscountAmount(value, currency string) {
	c.IssueDiscountAmount = NewActiveCurrencyAndAmount(value, currency)
}

func (c *CorporateActionAmounts1) SetSolicitationFees(value, currency string) {
	c.SolicitationFees = NewActiveCurrencyAndAmount(value, currency)
}

func (c *CorporateActionAmounts1) SetCashInLieuOfShare(value, currency string) {
	c.CashInLieuOfShare = NewActiveCurrencyAndAmount(value, currency)
}

func (c *CorporateActionAmounts1) SetOriginalAmount(value, currency string) {
	c.OriginalAmount = NewActiveCurrencyAndAmount(value, currency)
}

func (c *CorporateActionAmounts1) SetCapitalGain(value, currency string) {
	c.CapitalGain = NewActiveCurrencyAndAmount(value, currency)
}

func (c *CorporateActionAmounts1) SetInterestAmount(value, currency string) {
	c.InterestAmount = NewActiveCurrencyAndAmount(value, currency)
}

func (c *CorporateActionAmounts1) SetIndemnityAmount(value, currency string) {
	c.IndemnityAmount = NewActiveCurrencyAndAmount(value, currency)
}

func (c *CorporateActionAmounts1) SetRedemptionPremiumAmount(value, currency string) {
	c.RedemptionPremiumAmount = NewActiveCurrencyAndAmount(value, currency)
}

func (c *CorporateActionAmounts1) SetManufacturedDividendAmount(value, currency string) {
	c.ManufacturedDividendAmount = NewActiveCurrencyAndAmount(value, currency)
}

func (c *CorporateActionAmounts1) SetPrincipalOrCorpus(value, currency string) {
	c.PrincipalOrCorpus = NewActiveCurrencyAndAmount(value, currency)
}

func (c *CorporateActionAmounts1) SetReinvestmentAmount(value, currency string) {
	c.ReinvestmentAmount = NewActiveCurrencyAndAmount(value, currency)
}

func (c *CorporateActionAmounts1) SetMarketClaimAmount(value, currency string) {
	c.MarketClaimAmount = NewActiveCurrencyAndAmount(value, currency)
}

func (c *CorporateActionAmounts1) SetFullyFrankedAmount(value, currency string) {
	c.FullyFrankedAmount = NewActiveCurrencyAndAmount(value, currency)
}

func (c *CorporateActionAmounts1) SetUnfrankedAmount(value, currency string) {
	c.UnfrankedAmount = NewActiveCurrencyAndAmount(value, currency)
}

func (c *CorporateActionAmounts1) SetSundryOrOtherAmount(value, currency string) {
	c.SundryOrOtherAmount = NewActiveCurrencyAndAmount(value, currency)
}

func (c *CorporateActionAmounts1) SetSpecialConcessionAmount(value, currency string) {
	c.SpecialConcessionAmount = NewActiveCurrencyAndAmount(value, currency)
}

func (c *CorporateActionAmounts1) SetEntitledAmount(value, currency string) {
	c.EntitledAmount = NewActiveCurrencyAndAmount(value, currency)
}

func (c *CorporateActionAmounts1) SetCashIncentive(value, currency string) {
	c.CashIncentive = NewActiveCurrencyAndAmount(value, currency)
}

func (c *CorporateActionAmounts1) SetAdditionalSubscriptionCost(value, currency string) {
	c.AdditionalSubscriptionCost = NewActiveCurrencyAndAmount(value, currency)
}

func (c *CorporateActionAmounts1) SetTaxFreeAmount(value, currency string) {
	c.TaxFreeAmount = NewActiveCurrencyAndAmount(value, currency)
}

func (c *CorporateActionAmounts1) SetTaxDeferredAmount(value, currency string) {
	c.TaxDeferredAmount = NewActiveCurrencyAndAmount(value, currency)
}

func (c *CorporateActionAmounts1) SetGermanLocalTax1Amount(value, currency string) {
	c.GermanLocalTax1Amount = NewActiveCurrencyAndAmount(value, currency)
}

func (c *CorporateActionAmounts1) SetGermanLocalTax2Amount(value, currency string) {
	c.GermanLocalTax2Amount = NewActiveCurrencyAndAmount(value, currency)
}

func (c *CorporateActionAmounts1) SetGermanLocalTax3Amount(value, currency string) {
	c.GermanLocalTax3Amount = NewActiveCurrencyAndAmount(value, currency)
}

func (c *CorporateActionAmounts1) SetGermanLocalTax4Amount(value, currency string) {
	c.GermanLocalTax4Amount = NewActiveCurrencyAndAmount(value, currency)
}

func (c *CorporateActionAmounts1) SetStockExchangeTaxAmount(value, currency string) {
	c.StockExchangeTaxAmount = NewActiveCurrencyAndAmount(value, currency)
}

func (c *CorporateActionAmounts1) SetTransferTaxAmount(value, currency string) {
	c.TransferTaxAmount = NewActiveCurrencyAndAmount(value, currency)
}

func (c *CorporateActionAmounts1) SetTransactionTaxAmount(value, currency string) {
	c.TransactionTaxAmount = NewActiveCurrencyAndAmount(value, currency)
}

func (c *CorporateActionAmounts1) SetValueAddedTaxAmount(value, currency string) {
	c.ValueAddedTaxAmount = NewActiveCurrencyAndAmount(value, currency)
}

func (c *CorporateActionAmounts1) SetEURetentionTaxAmount(value, currency string) {
	c.EURetentionTaxAmount = NewActiveCurrencyAndAmount(value, currency)
}

func (c *CorporateActionAmounts1) SetLocalTaxAmount(value, currency string) {
	c.LocalTaxAmount = NewActiveCurrencyAndAmount(value, currency)
}

func (c *CorporateActionAmounts1) SetPaymentLevyTaxAmount(value, currency string) {
	c.PaymentLevyTaxAmount = NewActiveCurrencyAndAmount(value, currency)
}

func (c *CorporateActionAmounts1) SetCountryNationalFederalTaxAmount(value, currency string) {
	c.CountryNationalFederalTaxAmount = NewActiveCurrencyAndAmount(value, currency)
}

func (c *CorporateActionAmounts1) SetStampDutyAmount(value, currency string) {
	c.StampDutyAmount = NewActiveCurrencyAndAmount(value, currency)
}

func (c *CorporateActionAmounts1) SetTaxReclaimAmount(value, currency string) {
	c.TaxReclaimAmount = NewActiveCurrencyAndAmount(value, currency)
}

func (c *CorporateActionAmounts1) SetTaxCreditAmount(value, currency string) {
	c.TaxCreditAmount = NewActiveCurrencyAndAmount(value, currency)
}

func (c *CorporateActionAmounts1) SetWithholdingOfForeignTaxAmount(value, currency string) {
	c.WithholdingOfForeignTaxAmount = NewActiveCurrencyAndAmount(value, currency)
}

func (c *CorporateActionAmounts1) SetWithholdingOfLocalTaxAmount(value, currency string) {
	c.WithholdingOfLocalTaxAmount = NewActiveCurrencyAndAmount(value, currency)
}

func (c *CorporateActionAmounts1) SetAdditionalTaxAmount(value, currency string) {
	c.AdditionalTaxAmount = NewActiveCurrencyAndAmount(value, currency)
}

func (c *CorporateActionAmounts1) SetWithholdingTaxAmount(value, currency string) {
	c.WithholdingTaxAmount = NewActiveCurrencyAndAmount(value, currency)
}

func (c *CorporateActionAmounts1) SetFiscalStampAmount(value, currency string) {
	c.FiscalStampAmount = NewActiveCurrencyAndAmount(value, currency)
}

func (c *CorporateActionAmounts1) SetExecutingBrokerAmount(value, currency string) {
	c.ExecutingBrokerAmount = NewActiveCurrencyAndAmount(value, currency)
}

func (c *CorporateActionAmounts1) SetPayingAgentCommissionAmount(value, currency string) {
	c.PayingAgentCommissionAmount = NewActiveCurrencyAndAmount(value, currency)
}

func (c *CorporateActionAmounts1) SetLocalBrokerCommissionAmount(value, currency string) {
	c.LocalBrokerCommissionAmount = NewActiveCurrencyAndAmount(value, currency)
}

func (c *CorporateActionAmounts1) SetPostageFeeAmount(value, currency string) {
	c.PostageFeeAmount = NewActiveCurrencyAndAmount(value, currency)
}

func (c *CorporateActionAmounts1) SetRegulatoryFeesAmount(value, currency string) {
	c.RegulatoryFeesAmount = NewActiveCurrencyAndAmount(value, currency)
}

func (c *CorporateActionAmounts1) SetShippingFeesAmount(value, currency string) {
	c.ShippingFeesAmount = NewActiveCurrencyAndAmount(value, currency)
}

func (c *CorporateActionAmounts1) SetChargesAmount(value, currency string) {
	c.ChargesAmount = NewActiveCurrencyAndAmount(value, currency)
}
