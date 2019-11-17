package model

// Specifies amounts in the framework of a corporate action event.
type CorporateActionAmounts41 struct {

	// Amount of money before any deductions and allowances have been made.
	GrossCashAmount *RestrictedFINActiveCurrencyAndAmount `xml:"GrssCshAmt,omitempty"`

	// Amount of money after deductions and allowances have been made, if any, that is, the total amount +/- charges/fees.
	NetCashAmount *RestrictedFINActiveCurrencyAndAmount `xml:"NetCshAmt,omitempty"`

	// Cash premium made available if the securities holder consents or participates to an event, for example consent fees or solicitation fees.
	SolicitationFees *RestrictedFINActiveCurrencyAndAmount `xml:"SlctnFees,omitempty"`

	// Cash disbursement in lieu of a fractional quantity of, for example, equity.
	CashInLieuOfShare *RestrictedFINActiveCurrencyAndAmount `xml:"CshInLieuOfShr,omitempty"`

	// Amount of money distributed as the result of a capital gain.
	CapitalGain *RestrictedFINActiveCurrencyAndAmount `xml:"CptlGn,omitempty"`

	// Amount of money representing a coupon payment.
	InterestAmount *RestrictedFINActiveCurrencyAndAmount `xml:"IntrstAmt,omitempty"`

	// (Unique to France) Amount due to a buyer of securities dealt prior to ex date which may be subject to different rate of taxation.
	IndemnityAmount *RestrictedFINActiveCurrencyAndAmount `xml:"IndmntyAmt,omitempty"`

	// Amount of money that the borrower pays to the lender as a compensation. It does not entitle the lender to reclaim any tax credit and is sometimes treated differently by the local tax authorities of the lender. Also covers compensation/indemnity of missed dividend concerning early/late settlements if applicable to a market.
	ManufacturedDividendPaymentAmount *RestrictedFINActiveCurrencyAndAmount `xml:"ManfctrdDvddPmtAmt,omitempty"`

	// Amount of money reinvested in additional securities.
	ReinvestmentAmount *RestrictedFINActiveCurrencyAndAmount `xml:"RinvstmtAmt,omitempty"`

	// Amount resulting from a fully franked dividend paid by a company; amount includes tax credit for companies that have made sufficient tax payments during the fiscal period.
	FullyFrankedAmount *RestrictedFINActiveCurrencyAndAmount `xml:"FullyFrnkdAmt,omitempty"`

	// Amount resulting from an unfranked dividend paid by a company; the amount does not include tax credit and is subject to withholding tax.
	UnfrankedAmount *RestrictedFINActiveCurrencyAndAmount `xml:"UfrnkdAmt,omitempty"`

	// Amount of money related to taxable income that cannot be categorised.
	SundryOrOtherAmount *RestrictedFINActiveCurrencyAndAmount `xml:"SndryOrOthrAmt,omitempty"`

	// Amount of money that has not been subject to taxation.
	TaxFreeAmount *RestrictedFINActiveCurrencyAndAmount `xml:"TaxFreeAmt,omitempty"`

	// Amount of income eligible for deferred taxation.
	TaxDeferredAmount *RestrictedFINActiveCurrencyAndAmount `xml:"TaxDfrrdAmt,omitempty"`

	// Amount of value added tax.
	ValueAddedTaxAmount *RestrictedFINActiveCurrencyAndAmount `xml:"ValAddedTaxAmt,omitempty"`

	// Amount of stamp duty.
	StampDutyAmount *RestrictedFINActiveCurrencyAndAmount `xml:"StmpDtyAmt,omitempty"`

	// Amount that was paid in excess of actual tax obligation and was reclaimed.
	TaxReclaimAmount *RestrictedFINActiveCurrencyAndAmount `xml:"TaxRclmAmt,omitempty"`

	// Amount of taxes that have been previously paid in relation to the taxable event.
	TaxCreditAmount *RestrictedFINActiveCurrencyAndAmount `xml:"TaxCdtAmt,omitempty"`

	// Amount of additional taxes that cannot be categorised.
	AdditionalTaxAmount *RestrictedFINActiveCurrencyAndAmount `xml:"AddtlTaxAmt,omitempty"`

	// Amount of a cash distribution that will be withheld by the tax authorities of the jurisdiction of the issuer, for which a relief at source and/or reclaim may be possible.
	WithholdingTaxAmount *RestrictedFINActiveCurrencyAndAmount `xml:"WhldgTaxAmt,omitempty"`

	// Amount of money withheld by the jurisdiction other than the jurisdiction of the issuer’s country of tax incorporation, for which a relief at source and/or reclaim may be possible. It is levied in complement or offset of the withholding tax rate levied by the jurisdiction of the issuer’s tax domicile.
	SecondLevelTaxAmount *RestrictedFINActiveCurrencyAndAmount `xml:"ScndLvlTaxAmt,omitempty"`

	// Amount of fiscal tax to apply.
	FiscalStampAmount *RestrictedFINActiveCurrencyAndAmount `xml:"FsclStmpAmt,omitempty"`

	// Amount of money paid to an executing broker as a commission.
	ExecutingBrokerAmount *RestrictedFINActiveCurrencyAndAmount `xml:"ExctgBrkrAmt,omitempty"`

	// Amount of paying/sub-paying agent commission.
	PayingAgentCommissionAmount *RestrictedFINActiveCurrencyAndAmount `xml:"PngAgtComssnAmt,omitempty"`

	// Local broker's commission.
	LocalBrokerCommissionAmount *RestrictedFINActiveCurrencyAndAmount `xml:"LclBrkrComssnAmt,omitempty"`

	// Amount of money charged by a regulatory authority, for example, securities and exchange fees.
	RegulatoryFeesAmount *RestrictedFINActiveCurrencyAndAmount `xml:"RgltryFeesAmt,omitempty"`

	// All costs related to the physical delivery of documents such as stamps, postage, carrier fees, insurances or messenger services.
	ShippingFeesAmount *RestrictedFINActiveCurrencyAndAmount `xml:"ShppgFeesAmt,omitempty"`

	// Amount of money paid for the provision of financial services that cannot be categorised by another qualifier.
	ChargesAmount *RestrictedFINActiveCurrencyAndAmount `xml:"ChrgsAmt,omitempty"`

	// Cash amount based on terms of corporate action event and balance of underlying securities, entitled to/from account owner (which may be positive or negative).
	EntitledAmount *RestrictedFINActiveCurrencyAndAmount `xml:"EntitldAmt,omitempty"`

	// Posting/settlement amount in its original currency when conversion from/into another currency has occurred.
	OriginalAmount *RestrictedFINActiveCurrencyAndAmount `xml:"OrgnlAmt,omitempty"`

	// Amount of money representing a distribution of a bond's principal, for example, repayment of outstanding debt.
	PrincipalOrCorpus *RestrictedFINActiveCurrencyAndAmount `xml:"PrncplOrCrps,omitempty"`

	// Amount of money (not interest) in addition to the principal at the redemption of a bond.
	RedemptionPremiumAmount *RestrictedFINActiveCurrencyAndAmount `xml:"RedPrmAmt,omitempty"`

	// Amount relating to the underlying security for which income is distributed.
	IncomePortion *RestrictedFINActiveCurrencyAndAmount `xml:"IncmPrtn,omitempty"`

	// Amount of stock exchange tax.
	StockExchangeTax *RestrictedFINActiveCurrencyAndAmount `xml:"StockXchgTax,omitempty"`

	// Total amount of tax withheld at source in conformance with the EU Savings Directive.
	EUTaxRetentionAmount *RestrictedFINActiveCurrencyAndAmount `xml:"EUTaxRtntnAmt,omitempty"`

	// Amount of interest that has been accrued in between coupon payment periods
	AccruedInterestAmount *RestrictedFINActiveCurrencyAndAmount `xml:"AcrdIntrstAmt,omitempty"`

	// Portion of the fund distribution amount which represents the average accrued income included in the purchase price for units bought during the account period.
	EqualisationAmount *RestrictedFINActiveCurrencyAndAmount `xml:"EqulstnAmt,omitempty"`

	// FATCA (Foreign Account Tax Compliance Act) related tax amount.
	FATCATaxAmount *RestrictedFINActiveCurrencyAndAmount `xml:"FATCATaxAmt,omitempty"`

	// Amount of tax related income subject to NRA (Non Resident Alien).
	NRATaxAmount *RestrictedFINActiveCurrencyAndAmount `xml:"NRATaxAmt,omitempty"`

	// Amount of tax related to back up withholding.
	BackUpWithholdingTaxAmount *RestrictedFINActiveCurrencyAndAmount `xml:"BckUpWhldgTaxAmt,omitempty"`

	// Amount of overall tax withheld at source by fund managers prior to considering the tax obligation of each unit holder.
	TaxOnIncomeAmount *RestrictedFINActiveCurrencyAndAmount `xml:"TaxOnIncmAmt,omitempty"`

	// Amount of Transaction tax.
	TransactionTax *RestrictedFINActiveCurrencyAndAmount `xml:"TxTax,omitempty"`
}

func (c *CorporateActionAmounts41) SetGrossCashAmount(value, currency string) {
	c.GrossCashAmount = NewRestrictedFINActiveCurrencyAndAmount(value, currency)
}

func (c *CorporateActionAmounts41) SetNetCashAmount(value, currency string) {
	c.NetCashAmount = NewRestrictedFINActiveCurrencyAndAmount(value, currency)
}

func (c *CorporateActionAmounts41) SetSolicitationFees(value, currency string) {
	c.SolicitationFees = NewRestrictedFINActiveCurrencyAndAmount(value, currency)
}

func (c *CorporateActionAmounts41) SetCashInLieuOfShare(value, currency string) {
	c.CashInLieuOfShare = NewRestrictedFINActiveCurrencyAndAmount(value, currency)
}

func (c *CorporateActionAmounts41) SetCapitalGain(value, currency string) {
	c.CapitalGain = NewRestrictedFINActiveCurrencyAndAmount(value, currency)
}

func (c *CorporateActionAmounts41) SetInterestAmount(value, currency string) {
	c.InterestAmount = NewRestrictedFINActiveCurrencyAndAmount(value, currency)
}

func (c *CorporateActionAmounts41) SetIndemnityAmount(value, currency string) {
	c.IndemnityAmount = NewRestrictedFINActiveCurrencyAndAmount(value, currency)
}

func (c *CorporateActionAmounts41) SetManufacturedDividendPaymentAmount(value, currency string) {
	c.ManufacturedDividendPaymentAmount = NewRestrictedFINActiveCurrencyAndAmount(value, currency)
}

func (c *CorporateActionAmounts41) SetReinvestmentAmount(value, currency string) {
	c.ReinvestmentAmount = NewRestrictedFINActiveCurrencyAndAmount(value, currency)
}

func (c *CorporateActionAmounts41) SetFullyFrankedAmount(value, currency string) {
	c.FullyFrankedAmount = NewRestrictedFINActiveCurrencyAndAmount(value, currency)
}

func (c *CorporateActionAmounts41) SetUnfrankedAmount(value, currency string) {
	c.UnfrankedAmount = NewRestrictedFINActiveCurrencyAndAmount(value, currency)
}

func (c *CorporateActionAmounts41) SetSundryOrOtherAmount(value, currency string) {
	c.SundryOrOtherAmount = NewRestrictedFINActiveCurrencyAndAmount(value, currency)
}

func (c *CorporateActionAmounts41) SetTaxFreeAmount(value, currency string) {
	c.TaxFreeAmount = NewRestrictedFINActiveCurrencyAndAmount(value, currency)
}

func (c *CorporateActionAmounts41) SetTaxDeferredAmount(value, currency string) {
	c.TaxDeferredAmount = NewRestrictedFINActiveCurrencyAndAmount(value, currency)
}

func (c *CorporateActionAmounts41) SetValueAddedTaxAmount(value, currency string) {
	c.ValueAddedTaxAmount = NewRestrictedFINActiveCurrencyAndAmount(value, currency)
}

func (c *CorporateActionAmounts41) SetStampDutyAmount(value, currency string) {
	c.StampDutyAmount = NewRestrictedFINActiveCurrencyAndAmount(value, currency)
}

func (c *CorporateActionAmounts41) SetTaxReclaimAmount(value, currency string) {
	c.TaxReclaimAmount = NewRestrictedFINActiveCurrencyAndAmount(value, currency)
}

func (c *CorporateActionAmounts41) SetTaxCreditAmount(value, currency string) {
	c.TaxCreditAmount = NewRestrictedFINActiveCurrencyAndAmount(value, currency)
}

func (c *CorporateActionAmounts41) SetAdditionalTaxAmount(value, currency string) {
	c.AdditionalTaxAmount = NewRestrictedFINActiveCurrencyAndAmount(value, currency)
}

func (c *CorporateActionAmounts41) SetWithholdingTaxAmount(value, currency string) {
	c.WithholdingTaxAmount = NewRestrictedFINActiveCurrencyAndAmount(value, currency)
}

func (c *CorporateActionAmounts41) SetSecondLevelTaxAmount(value, currency string) {
	c.SecondLevelTaxAmount = NewRestrictedFINActiveCurrencyAndAmount(value, currency)
}

func (c *CorporateActionAmounts41) SetFiscalStampAmount(value, currency string) {
	c.FiscalStampAmount = NewRestrictedFINActiveCurrencyAndAmount(value, currency)
}

func (c *CorporateActionAmounts41) SetExecutingBrokerAmount(value, currency string) {
	c.ExecutingBrokerAmount = NewRestrictedFINActiveCurrencyAndAmount(value, currency)
}

func (c *CorporateActionAmounts41) SetPayingAgentCommissionAmount(value, currency string) {
	c.PayingAgentCommissionAmount = NewRestrictedFINActiveCurrencyAndAmount(value, currency)
}

func (c *CorporateActionAmounts41) SetLocalBrokerCommissionAmount(value, currency string) {
	c.LocalBrokerCommissionAmount = NewRestrictedFINActiveCurrencyAndAmount(value, currency)
}

func (c *CorporateActionAmounts41) SetRegulatoryFeesAmount(value, currency string) {
	c.RegulatoryFeesAmount = NewRestrictedFINActiveCurrencyAndAmount(value, currency)
}

func (c *CorporateActionAmounts41) SetShippingFeesAmount(value, currency string) {
	c.ShippingFeesAmount = NewRestrictedFINActiveCurrencyAndAmount(value, currency)
}

func (c *CorporateActionAmounts41) SetChargesAmount(value, currency string) {
	c.ChargesAmount = NewRestrictedFINActiveCurrencyAndAmount(value, currency)
}

func (c *CorporateActionAmounts41) SetEntitledAmount(value, currency string) {
	c.EntitledAmount = NewRestrictedFINActiveCurrencyAndAmount(value, currency)
}

func (c *CorporateActionAmounts41) SetOriginalAmount(value, currency string) {
	c.OriginalAmount = NewRestrictedFINActiveCurrencyAndAmount(value, currency)
}

func (c *CorporateActionAmounts41) SetPrincipalOrCorpus(value, currency string) {
	c.PrincipalOrCorpus = NewRestrictedFINActiveCurrencyAndAmount(value, currency)
}

func (c *CorporateActionAmounts41) SetRedemptionPremiumAmount(value, currency string) {
	c.RedemptionPremiumAmount = NewRestrictedFINActiveCurrencyAndAmount(value, currency)
}

func (c *CorporateActionAmounts41) SetIncomePortion(value, currency string) {
	c.IncomePortion = NewRestrictedFINActiveCurrencyAndAmount(value, currency)
}

func (c *CorporateActionAmounts41) SetStockExchangeTax(value, currency string) {
	c.StockExchangeTax = NewRestrictedFINActiveCurrencyAndAmount(value, currency)
}

func (c *CorporateActionAmounts41) SetEUTaxRetentionAmount(value, currency string) {
	c.EUTaxRetentionAmount = NewRestrictedFINActiveCurrencyAndAmount(value, currency)
}

func (c *CorporateActionAmounts41) SetAccruedInterestAmount(value, currency string) {
	c.AccruedInterestAmount = NewRestrictedFINActiveCurrencyAndAmount(value, currency)
}

func (c *CorporateActionAmounts41) SetEqualisationAmount(value, currency string) {
	c.EqualisationAmount = NewRestrictedFINActiveCurrencyAndAmount(value, currency)
}

func (c *CorporateActionAmounts41) SetFATCATaxAmount(value, currency string) {
	c.FATCATaxAmount = NewRestrictedFINActiveCurrencyAndAmount(value, currency)
}

func (c *CorporateActionAmounts41) SetNRATaxAmount(value, currency string) {
	c.NRATaxAmount = NewRestrictedFINActiveCurrencyAndAmount(value, currency)
}

func (c *CorporateActionAmounts41) SetBackUpWithholdingTaxAmount(value, currency string) {
	c.BackUpWithholdingTaxAmount = NewRestrictedFINActiveCurrencyAndAmount(value, currency)
}

func (c *CorporateActionAmounts41) SetTaxOnIncomeAmount(value, currency string) {
	c.TaxOnIncomeAmount = NewRestrictedFINActiveCurrencyAndAmount(value, currency)
}

func (c *CorporateActionAmounts41) SetTransactionTax(value, currency string) {
	c.TransactionTax = NewRestrictedFINActiveCurrencyAndAmount(value, currency)
}
