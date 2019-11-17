package model

// Identifies other amounts pertaining to the transaction.
type OtherAmounts32 struct {

	// Interest amount that has accrued in between coupon payment periods.
	AccruedInterestAmount *AmountAndDirection47 `xml:"AcrdIntrstAmt,omitempty"`

	// Amount of money paid for the provision of financial services that cannot be categorised by another qualifier.
	ChargesFees *AmountAndDirection47 `xml:"ChrgsFees,omitempty"`

	// Principal amount of a trade (price multiplied by quantity).
	TradeAmount *AmountAndDirection47 `xml:"TradAmt,omitempty"`

	// Amount of money paid to an executing broker as a commission.
	ExecutingBrokerAmount *AmountAndDirection47 `xml:"ExctgBrkrAmt,omitempty"`

	// Tax charged by the jurisdiction in which the financial instrument settles.
	LocalTax *AmountAndDirection47 `xml:"LclTax,omitempty"`

	// Amount of commission paid to a local broker.
	LocalBrokerCommission *AmountAndDirection47 `xml:"LclBrkrComssn,omitempty"`

	// An amount that is not indicated by a known business denomination.
	Other *AmountAndDirection47 `xml:"Othr,omitempty"`

	// Amount of stamp duty.
	StampDuty *AmountAndDirection47 `xml:"StmpDty,omitempty"`

	// Amount of transaction tax.
	TransactionTax *AmountAndDirection47 `xml:"TxTax,omitempty"`

	// Amount of money that will be withheld by a tax authority.
	WithholdingTax *AmountAndDirection47 `xml:"WhldgTax,omitempty"`

	// Amount of consumption tax.
	ConsumptionTax *AmountAndDirection47 `xml:"CsmptnTax,omitempty"`
}

func (o *OtherAmounts32) AddAccruedInterestAmount() *AmountAndDirection47 {
	o.AccruedInterestAmount = new(AmountAndDirection47)
	return o.AccruedInterestAmount
}

func (o *OtherAmounts32) AddChargesFees() *AmountAndDirection47 {
	o.ChargesFees = new(AmountAndDirection47)
	return o.ChargesFees
}

func (o *OtherAmounts32) AddTradeAmount() *AmountAndDirection47 {
	o.TradeAmount = new(AmountAndDirection47)
	return o.TradeAmount
}

func (o *OtherAmounts32) AddExecutingBrokerAmount() *AmountAndDirection47 {
	o.ExecutingBrokerAmount = new(AmountAndDirection47)
	return o.ExecutingBrokerAmount
}

func (o *OtherAmounts32) AddLocalTax() *AmountAndDirection47 {
	o.LocalTax = new(AmountAndDirection47)
	return o.LocalTax
}

func (o *OtherAmounts32) AddLocalBrokerCommission() *AmountAndDirection47 {
	o.LocalBrokerCommission = new(AmountAndDirection47)
	return o.LocalBrokerCommission
}

func (o *OtherAmounts32) AddOther() *AmountAndDirection47 {
	o.Other = new(AmountAndDirection47)
	return o.Other
}

func (o *OtherAmounts32) AddStampDuty() *AmountAndDirection47 {
	o.StampDuty = new(AmountAndDirection47)
	return o.StampDuty
}

func (o *OtherAmounts32) AddTransactionTax() *AmountAndDirection47 {
	o.TransactionTax = new(AmountAndDirection47)
	return o.TransactionTax
}

func (o *OtherAmounts32) AddWithholdingTax() *AmountAndDirection47 {
	o.WithholdingTax = new(AmountAndDirection47)
	return o.WithholdingTax
}

func (o *OtherAmounts32) AddConsumptionTax() *AmountAndDirection47 {
	o.ConsumptionTax = new(AmountAndDirection47)
	return o.ConsumptionTax
}
