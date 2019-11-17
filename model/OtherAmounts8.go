package model

// Identifies other amounts pertaining to the transaction.
type OtherAmounts8 struct {

	// Interest amount that has accrued in between coupon payment periods.
	AccruedInterestAmount *AmountAndDirection23 `xml:"AcrdIntrstAmt,omitempty"`

	// Amount of money paid for the provision of financial services that cannot be categorised by another qualifier.
	ChargesFees *AmountAndDirection23 `xml:"ChrgsFees,omitempty"`

	// Principal amount of a trade (price multiplied by quantity).
	TradeAmount *AmountAndDirection23 `xml:"TradAmt,omitempty"`

	// Amount of money paid to an executing broker as a commission.
	ExecutingBrokerAmount *AmountAndDirection23 `xml:"ExctgBrkrAmt,omitempty"`

	// Tax charged by the jurisdiction in which the financial instrument settles.
	LocalTax *AmountAndDirection23 `xml:"LclTax,omitempty"`

	// Amount of commission paid to a local broker.
	LocalBrokerCommission *AmountAndDirection23 `xml:"LclBrkrComssn,omitempty"`

	// An amount that is not indicated by a known business denomination.
	Other *AmountAndDirection23 `xml:"Othr,omitempty"`

	// Amount of stamp duty.
	StampDuty *AmountAndDirection23 `xml:"StmpDty,omitempty"`

	// Amount of transaction tax.
	TransactionTax *AmountAndDirection23 `xml:"TxTax,omitempty"`

	// Amount of money that will be withheld by a tax authority.
	WithholdingTax *AmountAndDirection23 `xml:"WhldgTax,omitempty"`

	// Amount of consumption tax.
	ConsumptionTax *AmountAndDirection23 `xml:"CsmptnTax,omitempty"`
}

func (o *OtherAmounts8) AddAccruedInterestAmount() *AmountAndDirection23 {
	o.AccruedInterestAmount = new(AmountAndDirection23)
	return o.AccruedInterestAmount
}

func (o *OtherAmounts8) AddChargesFees() *AmountAndDirection23 {
	o.ChargesFees = new(AmountAndDirection23)
	return o.ChargesFees
}

func (o *OtherAmounts8) AddTradeAmount() *AmountAndDirection23 {
	o.TradeAmount = new(AmountAndDirection23)
	return o.TradeAmount
}

func (o *OtherAmounts8) AddExecutingBrokerAmount() *AmountAndDirection23 {
	o.ExecutingBrokerAmount = new(AmountAndDirection23)
	return o.ExecutingBrokerAmount
}

func (o *OtherAmounts8) AddLocalTax() *AmountAndDirection23 {
	o.LocalTax = new(AmountAndDirection23)
	return o.LocalTax
}

func (o *OtherAmounts8) AddLocalBrokerCommission() *AmountAndDirection23 {
	o.LocalBrokerCommission = new(AmountAndDirection23)
	return o.LocalBrokerCommission
}

func (o *OtherAmounts8) AddOther() *AmountAndDirection23 {
	o.Other = new(AmountAndDirection23)
	return o.Other
}

func (o *OtherAmounts8) AddStampDuty() *AmountAndDirection23 {
	o.StampDuty = new(AmountAndDirection23)
	return o.StampDuty
}

func (o *OtherAmounts8) AddTransactionTax() *AmountAndDirection23 {
	o.TransactionTax = new(AmountAndDirection23)
	return o.TransactionTax
}

func (o *OtherAmounts8) AddWithholdingTax() *AmountAndDirection23 {
	o.WithholdingTax = new(AmountAndDirection23)
	return o.WithholdingTax
}

func (o *OtherAmounts8) AddConsumptionTax() *AmountAndDirection23 {
	o.ConsumptionTax = new(AmountAndDirection23)
	return o.ConsumptionTax
}
