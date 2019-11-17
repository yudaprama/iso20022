package model

// Identifies other amounts pertaining to the transaction.
type OtherAmounts31 struct {

	// Interest amount that has accrued in between coupon payment periods.
	AccruedInterestAmount *AmountAndDirection44 `xml:"AcrdIntrstAmt,omitempty"`

	// Amount of money paid for the provision of financial services that cannot be categorised by another qualifier.
	ChargesFees *AmountAndDirection44 `xml:"ChrgsFees,omitempty"`

	// Amount of country, national or federal tax charged by the jurisdiction in which the account servicer is located.
	CountryNationalFederalTax *AmountAndDirection44 `xml:"CtryNtlFdrlTax,omitempty"`

	// Principal amount of a trade (price multiplied by quantity).
	TradeAmount *AmountAndDirection44 `xml:"TradAmt,omitempty"`

	// Amount of money paid to an executing broker as a commission.
	ExecutingBrokerAmount *AmountAndDirection44 `xml:"ExctgBrkrAmt,omitempty"`

	// Amount of money defined as a discount on a new issue or on a tranche of an existing issue.
	IssueDiscountAllowance *AmountAndDirection44 `xml:"IsseDscntAllwnc,omitempty"`

	// Amount of payment levy tax.
	PaymentLevyTax *AmountAndDirection44 `xml:"PmtLevyTax,omitempty"`

	// Tax charged by the jurisdiction in which the financial instrument settles.
	LocalTax *AmountAndDirection44 `xml:"LclTax,omitempty"`

	// Amount of commission paid to a local broker.
	LocalBrokerCommission *AmountAndDirection44 `xml:"LclBrkrComssn,omitempty"`

	// Amount of money deposited by the trading party in a margin account.
	Margin *AmountAndDirection44 `xml:"Mrgn,omitempty"`

	// An amount that is not indicated by a known business denomination.
	Other *AmountAndDirection44 `xml:"Othr,omitempty"`

	// Amount of money charged by a regulatory authority, for example, Securities and Exchange fees.
	RegulatoryAmount *AmountAndDirection44 `xml:"RgltryAmt,omitempty"`

	// All costs related to the physical delivery of documents such as stamps, postage, carrier fees, insurances or messenger services.
	ShippingAmount *AmountAndDirection44 `xml:"ShppgAmt,omitempty"`

	// Amount of drawdown or other reduction from or in addition to the deal price.
	SpecialConcession *AmountAndDirection44 `xml:"SpclCncssn,omitempty"`

	// Amount of stamp duty.
	StampDuty *AmountAndDirection44 `xml:"StmpDty,omitempty"`

	// Amount of stock exchange tax.
	StockExchangeTax *AmountAndDirection44 `xml:"StockXchgTax,omitempty"`

	// Amount of tax levied on a transfer of ownership of financial instrument.
	TransferTax *AmountAndDirection44 `xml:"TrfTax,omitempty"`

	// Amount of transaction tax.
	TransactionTax *AmountAndDirection44 `xml:"TxTax,omitempty"`

	// Amount of value-added tax.
	ValueAddedTax *AmountAndDirection44 `xml:"ValAddedTax,omitempty"`

	// Amount of money that will be withheld by a tax authority.
	WithholdingTax *AmountAndDirection44 `xml:"WhldgTax,omitempty"`

	// Amount representing the difference between the cost and the current price of a security. In the context of securities settlement, it is the amount paid or received when the instructions are netted or paired off.
	NetGainLoss *AmountAndDirection44 `xml:"NetGnLoss,omitempty"`

	// Amount of consumption tax.
	ConsumptionTax *AmountAndDirection44 `xml:"CsmptnTax,omitempty"`

	// Amount of unpaid interest (on bonds which have defaulted and have subsequently
	// restructured), which is capitalized and added to the original principal amount of the bond.
	AccruedCapitalisationAmount *AmountAndDirection44 `xml:"AcrdCptlstnAmt,omitempty"`

	// Cost of the securities. May be requested in some countries for tax purposes.
	BookValue *AmountAndDirection44 `xml:"BookVal,omitempty"`
}

func (o *OtherAmounts31) AddAccruedInterestAmount() *AmountAndDirection44 {
	o.AccruedInterestAmount = new(AmountAndDirection44)
	return o.AccruedInterestAmount
}

func (o *OtherAmounts31) AddChargesFees() *AmountAndDirection44 {
	o.ChargesFees = new(AmountAndDirection44)
	return o.ChargesFees
}

func (o *OtherAmounts31) AddCountryNationalFederalTax() *AmountAndDirection44 {
	o.CountryNationalFederalTax = new(AmountAndDirection44)
	return o.CountryNationalFederalTax
}

func (o *OtherAmounts31) AddTradeAmount() *AmountAndDirection44 {
	o.TradeAmount = new(AmountAndDirection44)
	return o.TradeAmount
}

func (o *OtherAmounts31) AddExecutingBrokerAmount() *AmountAndDirection44 {
	o.ExecutingBrokerAmount = new(AmountAndDirection44)
	return o.ExecutingBrokerAmount
}

func (o *OtherAmounts31) AddIssueDiscountAllowance() *AmountAndDirection44 {
	o.IssueDiscountAllowance = new(AmountAndDirection44)
	return o.IssueDiscountAllowance
}

func (o *OtherAmounts31) AddPaymentLevyTax() *AmountAndDirection44 {
	o.PaymentLevyTax = new(AmountAndDirection44)
	return o.PaymentLevyTax
}

func (o *OtherAmounts31) AddLocalTax() *AmountAndDirection44 {
	o.LocalTax = new(AmountAndDirection44)
	return o.LocalTax
}

func (o *OtherAmounts31) AddLocalBrokerCommission() *AmountAndDirection44 {
	o.LocalBrokerCommission = new(AmountAndDirection44)
	return o.LocalBrokerCommission
}

func (o *OtherAmounts31) AddMargin() *AmountAndDirection44 {
	o.Margin = new(AmountAndDirection44)
	return o.Margin
}

func (o *OtherAmounts31) AddOther() *AmountAndDirection44 {
	o.Other = new(AmountAndDirection44)
	return o.Other
}

func (o *OtherAmounts31) AddRegulatoryAmount() *AmountAndDirection44 {
	o.RegulatoryAmount = new(AmountAndDirection44)
	return o.RegulatoryAmount
}

func (o *OtherAmounts31) AddShippingAmount() *AmountAndDirection44 {
	o.ShippingAmount = new(AmountAndDirection44)
	return o.ShippingAmount
}

func (o *OtherAmounts31) AddSpecialConcession() *AmountAndDirection44 {
	o.SpecialConcession = new(AmountAndDirection44)
	return o.SpecialConcession
}

func (o *OtherAmounts31) AddStampDuty() *AmountAndDirection44 {
	o.StampDuty = new(AmountAndDirection44)
	return o.StampDuty
}

func (o *OtherAmounts31) AddStockExchangeTax() *AmountAndDirection44 {
	o.StockExchangeTax = new(AmountAndDirection44)
	return o.StockExchangeTax
}

func (o *OtherAmounts31) AddTransferTax() *AmountAndDirection44 {
	o.TransferTax = new(AmountAndDirection44)
	return o.TransferTax
}

func (o *OtherAmounts31) AddTransactionTax() *AmountAndDirection44 {
	o.TransactionTax = new(AmountAndDirection44)
	return o.TransactionTax
}

func (o *OtherAmounts31) AddValueAddedTax() *AmountAndDirection44 {
	o.ValueAddedTax = new(AmountAndDirection44)
	return o.ValueAddedTax
}

func (o *OtherAmounts31) AddWithholdingTax() *AmountAndDirection44 {
	o.WithholdingTax = new(AmountAndDirection44)
	return o.WithholdingTax
}

func (o *OtherAmounts31) AddNetGainLoss() *AmountAndDirection44 {
	o.NetGainLoss = new(AmountAndDirection44)
	return o.NetGainLoss
}

func (o *OtherAmounts31) AddConsumptionTax() *AmountAndDirection44 {
	o.ConsumptionTax = new(AmountAndDirection44)
	return o.ConsumptionTax
}

func (o *OtherAmounts31) AddAccruedCapitalisationAmount() *AmountAndDirection44 {
	o.AccruedCapitalisationAmount = new(AmountAndDirection44)
	return o.AccruedCapitalisationAmount
}

func (o *OtherAmounts31) AddBookValue() *AmountAndDirection44 {
	o.BookValue = new(AmountAndDirection44)
	return o.BookValue
}
