package model

// Specifies amounts in the framework of a corporate action event.
type CorporateActionAmounts36 struct {

	// Amount of money before any deductions and allowances have been made.
	GrossCashAmount *ActiveCurrencyAndAmount `xml:"GrssCshAmt,omitempty"`

	// Amount of money after deductions and allowances have been made, if any, that is, the total amount +/- charges/fees.
	NetCashAmount *ActiveCurrencyAndAmount `xml:"NetCshAmt,omitempty"`

	// Cash premium made available if the securities holder consents or participates to an event, for example consent fees or solicitation fees.
	SolicitationFees *ActiveCurrencyAndAmount `xml:"SlctnFees,omitempty"`

	// Cash disbursement in lieu of a fractional quantity of, for example, equity.
	CashInLieuOfShare *ActiveCurrencyAndAmount `xml:"CshInLieuOfShr,omitempty"`

	// Amount of money distributed as the result of a capital gain.
	CapitalGain *ActiveCurrencyAndAmount `xml:"CptlGn,omitempty"`

	// Amount of money representing a coupon payment.
	InterestAmount *ActiveCurrencyAndAmount `xml:"IntrstAmt,omitempty"`

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

	// Amount of money that has not been subject to taxation.
	TaxFreeAmount *ActiveCurrencyAndAmount `xml:"TaxFreeAmt,omitempty"`

	// Amount of income eligible for deferred taxation.
	TaxDeferredAmount *ActiveCurrencyAndAmount `xml:"TaxDfrrdAmt,omitempty"`

	// Amount of value added tax.
	ValueAddedTaxAmount *ActiveCurrencyAndAmount `xml:"ValAddedTaxAmt,omitempty"`

	// Amount of stamp duty.
	StampDutyAmount *ActiveCurrencyAndAmount `xml:"StmpDtyAmt,omitempty"`

	// Amount that was paid in excess of actual tax obligation and was reclaimed.
	TaxReclaimAmount *ActiveCurrencyAndAmount `xml:"TaxRclmAmt,omitempty"`

	// Amount of taxes that have been previously paid in relation to the taxable event.
	TaxCreditAmount *ActiveCurrencyAndAmount `xml:"TaxCdtAmt,omitempty"`

	// Amount of additional taxes that cannot be categorised.
	AdditionalTaxAmount *ActiveCurrencyAndAmount `xml:"AddtlTaxAmt,omitempty"`

	// Amount of a cash distribution that will be withheld by the tax authorities of the jurisdiction of the issuer, for which a relief at source and/or reclaim may be possible.
	WithholdingTaxAmount *ActiveCurrencyAndAmount `xml:"WhldgTaxAmt,omitempty"`

	// Amount of money withheld by the jurisdiction other than the jurisdiction of the issuer’s country of tax incorporation, for which a relief at source and/or reclaim may be possible. It is levied in complement or offset of the withholding tax rate levied by the jurisdiction of the issuer’s tax domicile.
	SecondLevelTaxAmount *ActiveCurrencyAndAmount `xml:"ScndLvlTaxAmt,omitempty"`

	// Amount of fiscal tax to apply.
	FiscalStampAmount *ActiveCurrencyAndAmount `xml:"FsclStmpAmt,omitempty"`

	// Amount of money paid to an executing broker as a commission.
	ExecutingBrokerAmount *ActiveCurrencyAndAmount `xml:"ExctgBrkrAmt,omitempty"`

	// Amount of paying/sub-paying agent commission.
	PayingAgentCommissionAmount *ActiveCurrencyAndAmount `xml:"PngAgtComssnAmt,omitempty"`

	// Local broker's commission.
	LocalBrokerCommissionAmount *ActiveCurrencyAndAmount `xml:"LclBrkrComssnAmt,omitempty"`

	// Amount of money charged by a regulatory authority, for example, securities and exchange fees.
	RegulatoryFeesAmount *ActiveCurrencyAndAmount `xml:"RgltryFeesAmt,omitempty"`

	// All costs related to the physical delivery of documents such as stamps, postage, carrier fees, insurances or messenger services.
	ShippingFeesAmount *ActiveCurrencyAndAmount `xml:"ShppgFeesAmt,omitempty"`

	// Amount of money paid for the provision of financial services that cannot be categorised by another qualifier.
	ChargesAmount *ActiveCurrencyAndAmount `xml:"ChrgsAmt,omitempty"`

	// Cash amount based on terms of corporate action event and balance of underlying securities, entitled to/from account owner (which may be positive or negative).
	EntitledAmount *ActiveCurrencyAndAmount `xml:"EntitldAmt,omitempty"`

	// Posting/settlement amount in its original currency when conversion from/into another currency has occurred.
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

	// Portion of the fund distribution amount which represents the average accrued income included in the purchase price for units bought during the account period.
	EqualisationAmount *ActiveCurrencyAndAmount `xml:"EqulstnAmt,omitempty"`

	// FATCA (Foreign Account Tax Compliance Act) related tax amount.
	FATCATaxAmount *ActiveCurrencyAndAmount `xml:"FATCATaxAmt,omitempty"`

	// Amount of tax related income subject to NRA (Non Resident Alien).
	NRATaxAmount *ActiveCurrencyAndAmount `xml:"NRATaxAmt,omitempty"`

	// Amount of tax related to back up withholding.
	BackUpWithholdingTaxAmount *ActiveCurrencyAndAmount `xml:"BckUpWhldgTaxAmt,omitempty"`

	// Amount of overall tax withheld at source by fund managers prior to considering the tax obligation of each unit holder.
	TaxOnIncomeAmount *ActiveCurrencyAndAmount `xml:"TaxOnIncmAmt,omitempty"`

	// Amount of Transaction tax.
	TransactionTax *ActiveCurrencyAndAmount `xml:"TxTax,omitempty"`
}

func (c *CorporateActionAmounts36) SetGrossCashAmount(value, currency string) {
	c.GrossCashAmount = NewActiveCurrencyAndAmount(value, currency)
}

func (c *CorporateActionAmounts36) SetNetCashAmount(value, currency string) {
	c.NetCashAmount = NewActiveCurrencyAndAmount(value, currency)
}

func (c *CorporateActionAmounts36) SetSolicitationFees(value, currency string) {
	c.SolicitationFees = NewActiveCurrencyAndAmount(value, currency)
}

func (c *CorporateActionAmounts36) SetCashInLieuOfShare(value, currency string) {
	c.CashInLieuOfShare = NewActiveCurrencyAndAmount(value, currency)
}

func (c *CorporateActionAmounts36) SetCapitalGain(value, currency string) {
	c.CapitalGain = NewActiveCurrencyAndAmount(value, currency)
}

func (c *CorporateActionAmounts36) SetInterestAmount(value, currency string) {
	c.InterestAmount = NewActiveCurrencyAndAmount(value, currency)
}

func (c *CorporateActionAmounts36) SetIndemnityAmount(value, currency string) {
	c.IndemnityAmount = NewActiveCurrencyAndAmount(value, currency)
}

func (c *CorporateActionAmounts36) SetManufacturedDividendPaymentAmount(value, currency string) {
	c.ManufacturedDividendPaymentAmount = NewActiveCurrencyAndAmount(value, currency)
}

func (c *CorporateActionAmounts36) SetReinvestmentAmount(value, currency string) {
	c.ReinvestmentAmount = NewActiveCurrencyAndAmount(value, currency)
}

func (c *CorporateActionAmounts36) SetFullyFrankedAmount(value, currency string) {
	c.FullyFrankedAmount = NewActiveCurrencyAndAmount(value, currency)
}

func (c *CorporateActionAmounts36) SetUnfrankedAmount(value, currency string) {
	c.UnfrankedAmount = NewActiveCurrencyAndAmount(value, currency)
}

func (c *CorporateActionAmounts36) SetSundryOrOtherAmount(value, currency string) {
	c.SundryOrOtherAmount = NewActiveCurrencyAndAmount(value, currency)
}

func (c *CorporateActionAmounts36) SetTaxFreeAmount(value, currency string) {
	c.TaxFreeAmount = NewActiveCurrencyAndAmount(value, currency)
}

func (c *CorporateActionAmounts36) SetTaxDeferredAmount(value, currency string) {
	c.TaxDeferredAmount = NewActiveCurrencyAndAmount(value, currency)
}

func (c *CorporateActionAmounts36) SetValueAddedTaxAmount(value, currency string) {
	c.ValueAddedTaxAmount = NewActiveCurrencyAndAmount(value, currency)
}

func (c *CorporateActionAmounts36) SetStampDutyAmount(value, currency string) {
	c.StampDutyAmount = NewActiveCurrencyAndAmount(value, currency)
}

func (c *CorporateActionAmounts36) SetTaxReclaimAmount(value, currency string) {
	c.TaxReclaimAmount = NewActiveCurrencyAndAmount(value, currency)
}

func (c *CorporateActionAmounts36) SetTaxCreditAmount(value, currency string) {
	c.TaxCreditAmount = NewActiveCurrencyAndAmount(value, currency)
}

func (c *CorporateActionAmounts36) SetAdditionalTaxAmount(value, currency string) {
	c.AdditionalTaxAmount = NewActiveCurrencyAndAmount(value, currency)
}

func (c *CorporateActionAmounts36) SetWithholdingTaxAmount(value, currency string) {
	c.WithholdingTaxAmount = NewActiveCurrencyAndAmount(value, currency)
}

func (c *CorporateActionAmounts36) SetSecondLevelTaxAmount(value, currency string) {
	c.SecondLevelTaxAmount = NewActiveCurrencyAndAmount(value, currency)
}

func (c *CorporateActionAmounts36) SetFiscalStampAmount(value, currency string) {
	c.FiscalStampAmount = NewActiveCurrencyAndAmount(value, currency)
}

func (c *CorporateActionAmounts36) SetExecutingBrokerAmount(value, currency string) {
	c.ExecutingBrokerAmount = NewActiveCurrencyAndAmount(value, currency)
}

func (c *CorporateActionAmounts36) SetPayingAgentCommissionAmount(value, currency string) {
	c.PayingAgentCommissionAmount = NewActiveCurrencyAndAmount(value, currency)
}

func (c *CorporateActionAmounts36) SetLocalBrokerCommissionAmount(value, currency string) {
	c.LocalBrokerCommissionAmount = NewActiveCurrencyAndAmount(value, currency)
}

func (c *CorporateActionAmounts36) SetRegulatoryFeesAmount(value, currency string) {
	c.RegulatoryFeesAmount = NewActiveCurrencyAndAmount(value, currency)
}

func (c *CorporateActionAmounts36) SetShippingFeesAmount(value, currency string) {
	c.ShippingFeesAmount = NewActiveCurrencyAndAmount(value, currency)
}

func (c *CorporateActionAmounts36) SetChargesAmount(value, currency string) {
	c.ChargesAmount = NewActiveCurrencyAndAmount(value, currency)
}

func (c *CorporateActionAmounts36) SetEntitledAmount(value, currency string) {
	c.EntitledAmount = NewActiveCurrencyAndAmount(value, currency)
}

func (c *CorporateActionAmounts36) SetOriginalAmount(value, currency string) {
	c.OriginalAmount = NewActiveCurrencyAndAmount(value, currency)
}

func (c *CorporateActionAmounts36) SetPrincipalOrCorpus(value, currency string) {
	c.PrincipalOrCorpus = NewActiveCurrencyAndAmount(value, currency)
}

func (c *CorporateActionAmounts36) SetRedemptionPremiumAmount(value, currency string) {
	c.RedemptionPremiumAmount = NewActiveCurrencyAndAmount(value, currency)
}

func (c *CorporateActionAmounts36) SetIncomePortion(value, currency string) {
	c.IncomePortion = NewActiveCurrencyAndAmount(value, currency)
}

func (c *CorporateActionAmounts36) SetStockExchangeTax(value, currency string) {
	c.StockExchangeTax = NewActiveCurrencyAndAmount(value, currency)
}

func (c *CorporateActionAmounts36) SetEUTaxRetentionAmount(value, currency string) {
	c.EUTaxRetentionAmount = NewActiveCurrencyAndAmount(value, currency)
}

func (c *CorporateActionAmounts36) SetAccruedInterestAmount(value, currency string) {
	c.AccruedInterestAmount = NewActiveCurrencyAndAmount(value, currency)
}

func (c *CorporateActionAmounts36) SetEqualisationAmount(value, currency string) {
	c.EqualisationAmount = NewActiveCurrencyAndAmount(value, currency)
}

func (c *CorporateActionAmounts36) SetFATCATaxAmount(value, currency string) {
	c.FATCATaxAmount = NewActiveCurrencyAndAmount(value, currency)
}

func (c *CorporateActionAmounts36) SetNRATaxAmount(value, currency string) {
	c.NRATaxAmount = NewActiveCurrencyAndAmount(value, currency)
}

func (c *CorporateActionAmounts36) SetBackUpWithholdingTaxAmount(value, currency string) {
	c.BackUpWithholdingTaxAmount = NewActiveCurrencyAndAmount(value, currency)
}

func (c *CorporateActionAmounts36) SetTaxOnIncomeAmount(value, currency string) {
	c.TaxOnIncomeAmount = NewActiveCurrencyAndAmount(value, currency)
}

func (c *CorporateActionAmounts36) SetTransactionTax(value, currency string) {
	c.TransactionTax = NewActiveCurrencyAndAmount(value, currency)
}
