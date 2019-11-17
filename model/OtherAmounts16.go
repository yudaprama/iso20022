package model

// Identifies other amounts pertaining to the transaction.
type OtherAmounts16 struct {

	// Amount of money paid for the provision of financial services that cannot be categorised by another qualifier.
	ChargesFees *AmountAndDirection29 `xml:"ChrgsFees,omitempty"`

	// Amount of country, national or federal tax charged by the jurisdiction in which the account servicer is located.
	CountryNationalFederalTax *AmountAndDirection29 `xml:"CtryNtlFdrlTax,omitempty"`

	// Amount of money paid to an executing broker as a commission.
	ExecutingBrokerAmount *AmountAndDirection29 `xml:"ExctgBrkrAmt,omitempty"`

	// Amount of money defined as a discount on a new issue or on a tranche of an existing issue.
	IssueDiscountAllowance *AmountAndDirection29 `xml:"IsseDscntAllwnc,omitempty"`

	// Amount of payment levy tax.
	PaymentLevyTax *AmountAndDirection29 `xml:"PmtLevyTax,omitempty"`

	// Tax charged by the jurisdiction in which the financial instrument settles.
	LocalTax *AmountAndDirection29 `xml:"LclTax,omitempty"`

	// Amount of commission paid to a local broker.
	LocalBrokerCommission *AmountAndDirection29 `xml:"LclBrkrComssn,omitempty"`

	// Amount of money deposited by the trading party in a margin account.
	Margin *AmountAndDirection29 `xml:"Mrgn,omitempty"`

	// An amount that is not indicated by a known business denomination.
	Other *AmountAndDirection29 `xml:"Othr,omitempty"`

	// Amount of money charged by a regulatory authority, for example, Securities and Exchange fees.
	RegulatoryAmount *AmountAndDirection29 `xml:"RgltryAmt,omitempty"`

	// Amount of drawdown or other reduction from or in addition to the deal price.
	SpecialConcession *AmountAndDirection29 `xml:"SpclCncssn,omitempty"`

	// Amount of stamp duty.
	StampDuty *AmountAndDirection29 `xml:"StmpDty,omitempty"`

	// Amount of stock exchange tax.
	StockExchangeTax *AmountAndDirection29 `xml:"StockXchgTax,omitempty"`

	// Amount of tax levied on a transfer of ownership of financial instrument.
	TransferTax *AmountAndDirection29 `xml:"TrfTax,omitempty"`

	// Amount of transaction tax.
	TransactionTax *AmountAndDirection29 `xml:"TxTax,omitempty"`

	// Amount of value-added tax.
	ValueAddedTax *AmountAndDirection29 `xml:"ValAddedTax,omitempty"`

	// Amount of money that will be withheld by a tax authority.
	WithholdingTax *AmountAndDirection29 `xml:"WhldgTax,omitempty"`

	// Amount representing the difference between the cost and the current price of a security. In the context of securities settlement, it is the amount paid or received when the instructions are netted or paired off.
	NetGainLoss *AmountAndDirection29 `xml:"NetGnLoss,omitempty"`

	// A tax on spending on goods and services.
	ConsumptionTax *AmountAndDirection29 `xml:"CsmptnTax,omitempty"`

	// Amount of money charged for matching and/or confirmation.
	MatchingConfirmationFee *AmountAndDirection29 `xml:"MtchgConfFee,omitempty"`

	// Amount following a foreign exchange conversion.
	ConvertedAmount *AmountAndDirection29 `xml:"ConvtdAmt,omitempty"`

	// Posting/settlement amount in its original currency when conversion from/into another currency has occurred.
	OriginalCurrencyAmount *AmountAndDirection29 `xml:"OrgnlCcyAmt,omitempty"`

	// Cost of the securities. May be requested in some countries for tax purposes.
	BookValue *AmountAndDirection29 `xml:"BookVal,omitempty"`

	// Amount of unpaid interest (on bonds which have defaulted and have subsequently
	// restructured), which is capitalized and added to the original principal amount of the bond.
	AccruedCapitalisationAmount *AmountAndDirection29 `xml:"AcrdCptlstnAmt,omitempty"`

	// Local tax as defined by the country in its market practice document.
	LocalTaxCountrySpecific1 *AmountAndDirection29 `xml:"LclTaxCtrySpcfc1,omitempty"`

	// Local tax as defined by the country in its market practice document.
	LocalTaxCountrySpecific2 *AmountAndDirection29 `xml:"LclTaxCtrySpcfc2,omitempty"`

	// Local tax as defined by the country in its market practice document.
	LocalTaxCountrySpecific3 *AmountAndDirection29 `xml:"LclTaxCtrySpcfc3,omitempty"`

	// Local tax as defined by the country in its market practice document.
	LocalTaxCountrySpecific4 *AmountAndDirection29 `xml:"LclTaxCtrySpcfc4,omitempty"`

	// Amount paid as result of transactions subject to shared brokerage commissions.
	SharedBrokerageAmount *AmountAndDirection29 `xml:"ShrdBrkrgAmt,omitempty"`

	// Indicates the total fees related to the trade, with deduction of rebates (on brokerage, commission or differential) granted by the market member (fee amount with tax excluded).
	MarketMemberFeeAmount *AmountAndDirection29 `xml:"MktMmbFeeAmt,omitempty"`

	// Specifies that this information is available upon request.
	RemunerationAmountRequest *YesNoIndicator `xml:"RmnrtnAmtReq,omitempty"`

	// Specifies the source and amount of any other remuneration received or to be received by the broker in connection with the transaction.
	RemunerationAmount *AmountAndDirection29 `xml:"RmnrtnAmt,omitempty"`

	// Amount to be paid by the lender to the borrower calculated based on the interest rate.
	BorrowingInterestAmount *AmountAndDirection29 `xml:"BrrwgIntrstAmt,omitempty"`

	// Amount to be paid by the Borrower to the Lender for the securities borrowed calculated based on the bond loan rate.
	BorrowingFee *AmountAndDirection29 `xml:"BrrwgFee,omitempty"`

	// Net market value of the securities lent used to calculate the collateral amount.
	NetMarketValue *AmountAndDirection29 `xml:"NetMktVal,omitempty"`

	// Remaining nominal value of a security.
	RemainingFaceValue *AmountAndDirection29 `xml:"RmngFaceVal,omitempty"`

	// Remaining value at which an asset is carried on a balance sheet.
	RemainingBookValue *AmountAndDirection29 `xml:"RmngBookVal,omitempty"`

	// Amount of commission paid to a clearing broker.
	ClearingBrokerCommission *AmountAndDirection29 `xml:"ClrBrkrComssn,omitempty"`

	// Difference between the deal price and another reference price.
	DifferenceInPrice *AmountAndDirection29 `xml:"DiffInPric,omitempty"`

	// Specifies that the odd-lot differential or equivalent fee has been paid by such customer in connection with the execution of an order for an odd-lot number of shares or units (or principal amount) of a security and the fact that the amount of any such differential or fee will be furnished upon oral or written request.
	OddLotFee *YesNoIndicator `xml:"OddLotFee,omitempty"`
}

func (o *OtherAmounts16) AddChargesFees() *AmountAndDirection29 {
	o.ChargesFees = new(AmountAndDirection29)
	return o.ChargesFees
}

func (o *OtherAmounts16) AddCountryNationalFederalTax() *AmountAndDirection29 {
	o.CountryNationalFederalTax = new(AmountAndDirection29)
	return o.CountryNationalFederalTax
}

func (o *OtherAmounts16) AddExecutingBrokerAmount() *AmountAndDirection29 {
	o.ExecutingBrokerAmount = new(AmountAndDirection29)
	return o.ExecutingBrokerAmount
}

func (o *OtherAmounts16) AddIssueDiscountAllowance() *AmountAndDirection29 {
	o.IssueDiscountAllowance = new(AmountAndDirection29)
	return o.IssueDiscountAllowance
}

func (o *OtherAmounts16) AddPaymentLevyTax() *AmountAndDirection29 {
	o.PaymentLevyTax = new(AmountAndDirection29)
	return o.PaymentLevyTax
}

func (o *OtherAmounts16) AddLocalTax() *AmountAndDirection29 {
	o.LocalTax = new(AmountAndDirection29)
	return o.LocalTax
}

func (o *OtherAmounts16) AddLocalBrokerCommission() *AmountAndDirection29 {
	o.LocalBrokerCommission = new(AmountAndDirection29)
	return o.LocalBrokerCommission
}

func (o *OtherAmounts16) AddMargin() *AmountAndDirection29 {
	o.Margin = new(AmountAndDirection29)
	return o.Margin
}

func (o *OtherAmounts16) AddOther() *AmountAndDirection29 {
	o.Other = new(AmountAndDirection29)
	return o.Other
}

func (o *OtherAmounts16) AddRegulatoryAmount() *AmountAndDirection29 {
	o.RegulatoryAmount = new(AmountAndDirection29)
	return o.RegulatoryAmount
}

func (o *OtherAmounts16) AddSpecialConcession() *AmountAndDirection29 {
	o.SpecialConcession = new(AmountAndDirection29)
	return o.SpecialConcession
}

func (o *OtherAmounts16) AddStampDuty() *AmountAndDirection29 {
	o.StampDuty = new(AmountAndDirection29)
	return o.StampDuty
}

func (o *OtherAmounts16) AddStockExchangeTax() *AmountAndDirection29 {
	o.StockExchangeTax = new(AmountAndDirection29)
	return o.StockExchangeTax
}

func (o *OtherAmounts16) AddTransferTax() *AmountAndDirection29 {
	o.TransferTax = new(AmountAndDirection29)
	return o.TransferTax
}

func (o *OtherAmounts16) AddTransactionTax() *AmountAndDirection29 {
	o.TransactionTax = new(AmountAndDirection29)
	return o.TransactionTax
}

func (o *OtherAmounts16) AddValueAddedTax() *AmountAndDirection29 {
	o.ValueAddedTax = new(AmountAndDirection29)
	return o.ValueAddedTax
}

func (o *OtherAmounts16) AddWithholdingTax() *AmountAndDirection29 {
	o.WithholdingTax = new(AmountAndDirection29)
	return o.WithholdingTax
}

func (o *OtherAmounts16) AddNetGainLoss() *AmountAndDirection29 {
	o.NetGainLoss = new(AmountAndDirection29)
	return o.NetGainLoss
}

func (o *OtherAmounts16) AddConsumptionTax() *AmountAndDirection29 {
	o.ConsumptionTax = new(AmountAndDirection29)
	return o.ConsumptionTax
}

func (o *OtherAmounts16) AddMatchingConfirmationFee() *AmountAndDirection29 {
	o.MatchingConfirmationFee = new(AmountAndDirection29)
	return o.MatchingConfirmationFee
}

func (o *OtherAmounts16) AddConvertedAmount() *AmountAndDirection29 {
	o.ConvertedAmount = new(AmountAndDirection29)
	return o.ConvertedAmount
}

func (o *OtherAmounts16) AddOriginalCurrencyAmount() *AmountAndDirection29 {
	o.OriginalCurrencyAmount = new(AmountAndDirection29)
	return o.OriginalCurrencyAmount
}

func (o *OtherAmounts16) AddBookValue() *AmountAndDirection29 {
	o.BookValue = new(AmountAndDirection29)
	return o.BookValue
}

func (o *OtherAmounts16) AddAccruedCapitalisationAmount() *AmountAndDirection29 {
	o.AccruedCapitalisationAmount = new(AmountAndDirection29)
	return o.AccruedCapitalisationAmount
}

func (o *OtherAmounts16) AddLocalTaxCountrySpecific1() *AmountAndDirection29 {
	o.LocalTaxCountrySpecific1 = new(AmountAndDirection29)
	return o.LocalTaxCountrySpecific1
}

func (o *OtherAmounts16) AddLocalTaxCountrySpecific2() *AmountAndDirection29 {
	o.LocalTaxCountrySpecific2 = new(AmountAndDirection29)
	return o.LocalTaxCountrySpecific2
}

func (o *OtherAmounts16) AddLocalTaxCountrySpecific3() *AmountAndDirection29 {
	o.LocalTaxCountrySpecific3 = new(AmountAndDirection29)
	return o.LocalTaxCountrySpecific3
}

func (o *OtherAmounts16) AddLocalTaxCountrySpecific4() *AmountAndDirection29 {
	o.LocalTaxCountrySpecific4 = new(AmountAndDirection29)
	return o.LocalTaxCountrySpecific4
}

func (o *OtherAmounts16) AddSharedBrokerageAmount() *AmountAndDirection29 {
	o.SharedBrokerageAmount = new(AmountAndDirection29)
	return o.SharedBrokerageAmount
}

func (o *OtherAmounts16) AddMarketMemberFeeAmount() *AmountAndDirection29 {
	o.MarketMemberFeeAmount = new(AmountAndDirection29)
	return o.MarketMemberFeeAmount
}

func (o *OtherAmounts16) SetRemunerationAmountRequest(value string) {
	o.RemunerationAmountRequest = (*YesNoIndicator)(&value)
}

func (o *OtherAmounts16) AddRemunerationAmount() *AmountAndDirection29 {
	o.RemunerationAmount = new(AmountAndDirection29)
	return o.RemunerationAmount
}

func (o *OtherAmounts16) AddBorrowingInterestAmount() *AmountAndDirection29 {
	o.BorrowingInterestAmount = new(AmountAndDirection29)
	return o.BorrowingInterestAmount
}

func (o *OtherAmounts16) AddBorrowingFee() *AmountAndDirection29 {
	o.BorrowingFee = new(AmountAndDirection29)
	return o.BorrowingFee
}

func (o *OtherAmounts16) AddNetMarketValue() *AmountAndDirection29 {
	o.NetMarketValue = new(AmountAndDirection29)
	return o.NetMarketValue
}

func (o *OtherAmounts16) AddRemainingFaceValue() *AmountAndDirection29 {
	o.RemainingFaceValue = new(AmountAndDirection29)
	return o.RemainingFaceValue
}

func (o *OtherAmounts16) AddRemainingBookValue() *AmountAndDirection29 {
	o.RemainingBookValue = new(AmountAndDirection29)
	return o.RemainingBookValue
}

func (o *OtherAmounts16) AddClearingBrokerCommission() *AmountAndDirection29 {
	o.ClearingBrokerCommission = new(AmountAndDirection29)
	return o.ClearingBrokerCommission
}

func (o *OtherAmounts16) AddDifferenceInPrice() *AmountAndDirection29 {
	o.DifferenceInPrice = new(AmountAndDirection29)
	return o.DifferenceInPrice
}

func (o *OtherAmounts16) SetOddLotFee(value string) {
	o.OddLotFee = (*YesNoIndicator)(&value)
}
