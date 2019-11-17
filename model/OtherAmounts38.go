package model

// Identifies other amounts pertaining to the transaction.
type OtherAmounts38 struct {

	// Interest amount that has accrued in between coupon payment periods.
	AccruedInterestAmount *AmountAndDirection58 `xml:"AcrdIntrstAmt,omitempty"`

	// Amount of money paid for the provision of financial services that cannot be categorised by another qualifier.
	ChargesFees *AmountAndDirection58 `xml:"ChrgsFees,omitempty"`

	// Amount of country, national or federal tax charged by the jurisdiction in which the account servicer is located.
	CountryNationalFederalTax *AmountAndDirection58 `xml:"CtryNtlFdrlTax,omitempty"`

	// Principal amount of a trade (price multiplied by quantity).
	TradeAmount *AmountAndDirection58 `xml:"TradAmt,omitempty"`

	// Amount of money paid to an executing broker as a commission.
	ExecutingBrokerAmount *AmountAndDirection58 `xml:"ExctgBrkrAmt,omitempty"`

	// Amount of money defined as a discount on a new issue or on a tranche of an existing issue.
	IssueDiscountAllowance *AmountAndDirection58 `xml:"IsseDscntAllwnc,omitempty"`

	// Amount of payment levy tax.
	PaymentLevyTax *AmountAndDirection58 `xml:"PmtLevyTax,omitempty"`

	// Tax charged by the jurisdiction in which the financial instrument settles.
	LocalTax *AmountAndDirection58 `xml:"LclTax,omitempty"`

	// Local tax country specific.
	LocalTaxCountrySpecific *AmountAndDirection58 `xml:"LclTaxCtrySpcfc,omitempty"`

	// Amount of commission paid to a local broker.
	LocalBrokerCommission *AmountAndDirection58 `xml:"LclBrkrComssn,omitempty"`

	// Amount of money deposited by the trading party in a margin account.
	Margin *AmountAndDirection58 `xml:"Mrgn,omitempty"`

	// An amount that is not indicated by a known business denomination.
	Other *AmountAndDirection58 `xml:"Othr,omitempty"`

	// Amount of money charged by a regulatory authority, for example, Securities and Exchange fees.
	RegulatoryAmount *AmountAndDirection58 `xml:"RgltryAmt,omitempty"`

	// All costs related to the physical delivery of documents such as stamps, postage, carrier fees, insurances or messenger services.
	ShippingAmount *AmountAndDirection58 `xml:"ShppgAmt,omitempty"`

	// Amount of drawdown or other reduction from or in addition to the deal price.
	SpecialConcession *AmountAndDirection58 `xml:"SpclCncssn,omitempty"`

	// Amount of stamp duty.
	StampDuty *AmountAndDirection58 `xml:"StmpDty,omitempty"`

	// Amount of stock exchange tax.
	StockExchangeTax *AmountAndDirection58 `xml:"StockXchgTax,omitempty"`

	// Amount of tax levied on a transfer of ownership of financial instrument.
	TransferTax *AmountAndDirection58 `xml:"TrfTax,omitempty"`

	// Amount of transaction tax.
	TransactionTax *AmountAndDirection58 `xml:"TxTax,omitempty"`

	// Amount of value-added tax.
	ValueAddedTax *AmountAndDirection58 `xml:"ValAddedTax,omitempty"`

	// Amount of money that will be withheld by a tax authority.
	WithholdingTax *AmountAndDirection58 `xml:"WhldgTax,omitempty"`

	// Amount representing the difference between the cost and the current price of a security. In the context of securities settlement, it is the amount paid or received when the instructions are netted or paired off.
	NetGainLoss *AmountAndDirection58 `xml:"NetGnLoss,omitempty"`

	// Amount of consumption tax.
	ConsumptionTax *AmountAndDirection58 `xml:"CsmptnTax,omitempty"`

	// Amount of unpaid interest (on bonds which have defaulted and have subsequently
	// restructured), which is capitalized and added to the original principal amount of the bond.
	AccruedCapitalisationAmount *AmountAndDirection58 `xml:"AcrdCptlstnAmt,omitempty"`

	// Cost of the securities. May be requested in some countries for tax purposes.
	BookValue *AmountAndDirection58 `xml:"BookVal,omitempty"`

	// Value of the collateral available for the delivery settlement process at the account level.
	CollateralMonitorAmount *AmountAndDirection58 `xml:"CollMntrAmt,omitempty"`
}

func (o *OtherAmounts38) AddAccruedInterestAmount() *AmountAndDirection58 {
	o.AccruedInterestAmount = new(AmountAndDirection58)
	return o.AccruedInterestAmount
}

func (o *OtherAmounts38) AddChargesFees() *AmountAndDirection58 {
	o.ChargesFees = new(AmountAndDirection58)
	return o.ChargesFees
}

func (o *OtherAmounts38) AddCountryNationalFederalTax() *AmountAndDirection58 {
	o.CountryNationalFederalTax = new(AmountAndDirection58)
	return o.CountryNationalFederalTax
}

func (o *OtherAmounts38) AddTradeAmount() *AmountAndDirection58 {
	o.TradeAmount = new(AmountAndDirection58)
	return o.TradeAmount
}

func (o *OtherAmounts38) AddExecutingBrokerAmount() *AmountAndDirection58 {
	o.ExecutingBrokerAmount = new(AmountAndDirection58)
	return o.ExecutingBrokerAmount
}

func (o *OtherAmounts38) AddIssueDiscountAllowance() *AmountAndDirection58 {
	o.IssueDiscountAllowance = new(AmountAndDirection58)
	return o.IssueDiscountAllowance
}

func (o *OtherAmounts38) AddPaymentLevyTax() *AmountAndDirection58 {
	o.PaymentLevyTax = new(AmountAndDirection58)
	return o.PaymentLevyTax
}

func (o *OtherAmounts38) AddLocalTax() *AmountAndDirection58 {
	o.LocalTax = new(AmountAndDirection58)
	return o.LocalTax
}

func (o *OtherAmounts38) AddLocalTaxCountrySpecific() *AmountAndDirection58 {
	o.LocalTaxCountrySpecific = new(AmountAndDirection58)
	return o.LocalTaxCountrySpecific
}

func (o *OtherAmounts38) AddLocalBrokerCommission() *AmountAndDirection58 {
	o.LocalBrokerCommission = new(AmountAndDirection58)
	return o.LocalBrokerCommission
}

func (o *OtherAmounts38) AddMargin() *AmountAndDirection58 {
	o.Margin = new(AmountAndDirection58)
	return o.Margin
}

func (o *OtherAmounts38) AddOther() *AmountAndDirection58 {
	o.Other = new(AmountAndDirection58)
	return o.Other
}

func (o *OtherAmounts38) AddRegulatoryAmount() *AmountAndDirection58 {
	o.RegulatoryAmount = new(AmountAndDirection58)
	return o.RegulatoryAmount
}

func (o *OtherAmounts38) AddShippingAmount() *AmountAndDirection58 {
	o.ShippingAmount = new(AmountAndDirection58)
	return o.ShippingAmount
}

func (o *OtherAmounts38) AddSpecialConcession() *AmountAndDirection58 {
	o.SpecialConcession = new(AmountAndDirection58)
	return o.SpecialConcession
}

func (o *OtherAmounts38) AddStampDuty() *AmountAndDirection58 {
	o.StampDuty = new(AmountAndDirection58)
	return o.StampDuty
}

func (o *OtherAmounts38) AddStockExchangeTax() *AmountAndDirection58 {
	o.StockExchangeTax = new(AmountAndDirection58)
	return o.StockExchangeTax
}

func (o *OtherAmounts38) AddTransferTax() *AmountAndDirection58 {
	o.TransferTax = new(AmountAndDirection58)
	return o.TransferTax
}

func (o *OtherAmounts38) AddTransactionTax() *AmountAndDirection58 {
	o.TransactionTax = new(AmountAndDirection58)
	return o.TransactionTax
}

func (o *OtherAmounts38) AddValueAddedTax() *AmountAndDirection58 {
	o.ValueAddedTax = new(AmountAndDirection58)
	return o.ValueAddedTax
}

func (o *OtherAmounts38) AddWithholdingTax() *AmountAndDirection58 {
	o.WithholdingTax = new(AmountAndDirection58)
	return o.WithholdingTax
}

func (o *OtherAmounts38) AddNetGainLoss() *AmountAndDirection58 {
	o.NetGainLoss = new(AmountAndDirection58)
	return o.NetGainLoss
}

func (o *OtherAmounts38) AddConsumptionTax() *AmountAndDirection58 {
	o.ConsumptionTax = new(AmountAndDirection58)
	return o.ConsumptionTax
}

func (o *OtherAmounts38) AddAccruedCapitalisationAmount() *AmountAndDirection58 {
	o.AccruedCapitalisationAmount = new(AmountAndDirection58)
	return o.AccruedCapitalisationAmount
}

func (o *OtherAmounts38) AddBookValue() *AmountAndDirection58 {
	o.BookValue = new(AmountAndDirection58)
	return o.BookValue
}

func (o *OtherAmounts38) AddCollateralMonitorAmount() *AmountAndDirection58 {
	o.CollateralMonitorAmount = new(AmountAndDirection58)
	return o.CollateralMonitorAmount
}
