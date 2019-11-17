package model

// Identifies other amounts pertaining to the transaction.
type OtherAmounts36 struct {

	// Interest amount that has accrued in between coupon payment periods.
	AccruedInterestAmount *AmountAndDirection72 `xml:"AcrdIntrstAmt,omitempty"`

	// Amount of money paid for the provision of financial services that cannot be categorised by another qualifier.
	ChargesFees *AmountAndDirection72 `xml:"ChrgsFees,omitempty"`

	// Principal amount of a trade (price multiplied by quantity).
	TradeAmount *AmountAndDirection72 `xml:"TradAmt,omitempty"`

	// Amount of money paid to an executing broker as a commission.
	ExecutingBrokerAmount *AmountAndDirection72 `xml:"ExctgBrkrAmt,omitempty"`

	// Tax charged by the jurisdiction in which the financial instrument settles.
	LocalTax *AmountAndDirection72 `xml:"LclTax,omitempty"`

	// Amount of commission paid to a local broker.
	LocalBrokerCommission *AmountAndDirection72 `xml:"LclBrkrComssn,omitempty"`

	// An amount that is not indicated by a known business denomination.
	Other *AmountAndDirection72 `xml:"Othr,omitempty"`

	// Amount of stamp duty.
	StampDuty *AmountAndDirection72 `xml:"StmpDty,omitempty"`

	// Amount of transaction tax.
	TransactionTax *AmountAndDirection72 `xml:"TxTax,omitempty"`

	// Amount of money that will be withheld by a tax authority.
	WithholdingTax *AmountAndDirection72 `xml:"WhldgTax,omitempty"`

	// Amount of consumption tax.
	ConsumptionTax *AmountAndDirection72 `xml:"CsmptnTax,omitempty"`
}

func (o *OtherAmounts36) AddAccruedInterestAmount() *AmountAndDirection72 {
	o.AccruedInterestAmount = new(AmountAndDirection72)
	return o.AccruedInterestAmount
}

func (o *OtherAmounts36) AddChargesFees() *AmountAndDirection72 {
	o.ChargesFees = new(AmountAndDirection72)
	return o.ChargesFees
}

func (o *OtherAmounts36) AddTradeAmount() *AmountAndDirection72 {
	o.TradeAmount = new(AmountAndDirection72)
	return o.TradeAmount
}

func (o *OtherAmounts36) AddExecutingBrokerAmount() *AmountAndDirection72 {
	o.ExecutingBrokerAmount = new(AmountAndDirection72)
	return o.ExecutingBrokerAmount
}

func (o *OtherAmounts36) AddLocalTax() *AmountAndDirection72 {
	o.LocalTax = new(AmountAndDirection72)
	return o.LocalTax
}

func (o *OtherAmounts36) AddLocalBrokerCommission() *AmountAndDirection72 {
	o.LocalBrokerCommission = new(AmountAndDirection72)
	return o.LocalBrokerCommission
}

func (o *OtherAmounts36) AddOther() *AmountAndDirection72 {
	o.Other = new(AmountAndDirection72)
	return o.Other
}

func (o *OtherAmounts36) AddStampDuty() *AmountAndDirection72 {
	o.StampDuty = new(AmountAndDirection72)
	return o.StampDuty
}

func (o *OtherAmounts36) AddTransactionTax() *AmountAndDirection72 {
	o.TransactionTax = new(AmountAndDirection72)
	return o.TransactionTax
}

func (o *OtherAmounts36) AddWithholdingTax() *AmountAndDirection72 {
	o.WithholdingTax = new(AmountAndDirection72)
	return o.WithholdingTax
}

func (o *OtherAmounts36) AddConsumptionTax() *AmountAndDirection72 {
	o.ConsumptionTax = new(AmountAndDirection72)
	return o.ConsumptionTax
}
