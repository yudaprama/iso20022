package model

// Identifies other amounts pertaining to the transaction.
type OtherAmounts29 struct {

	// Interest amount that has accrued in between coupon payment periods.
	AccruedInterestAmount *AmountAndDirection44 `xml:"AcrdIntrstAmt,omitempty"`

	// Amount of money paid for the provision of financial services that cannot be categorised by another qualifier.
	ChargesFees *AmountAndDirection44 `xml:"ChrgsFees,omitempty"`

	// Amount of country, national or federal tax charged by the jurisdiction in which the account servicer is located.
	CountryNationalFederalTax *AmountAndDirection44 `xml:"CtryNtlFdrlTax,omitempty"`

	// Amount of payment levy tax.
	PaymentLevyTax *AmountAndDirection44 `xml:"PmtLevyTax,omitempty"`

	// Tax charged by the jurisdiction in which the financial instrument settles.
	LocalTax *AmountAndDirection44 `xml:"LclTax,omitempty"`

	// An amount that is not indicated by a known business denomination.
	Other *AmountAndDirection44 `xml:"Othr,omitempty"`

	// Amount of money charged by a regulatory authority, for example, Securities and Exchange fees.
	RegulatoryAmount *AmountAndDirection44 `xml:"RgltryAmt,omitempty"`

	// All costs related to the physical delivery of documents such as stamps, postage, carrier fees, insurances or messenger services.
	ShippingAmount *AmountAndDirection44 `xml:"ShppgAmt,omitempty"`

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

	// Amount of consumption tax.
	ConsumptionTax *AmountAndDirection44 `xml:"CsmptnTax,omitempty"`

	// Amount of unpaid interest (on bonds which have defaulted and have subsequently
	// restructured), which is capitalized and added to the original principal amount of the bond.
	AccruedCapitalisationAmount *AmountAndDirection44 `xml:"AcrdCptlstnAmt,omitempty"`
}

func (o *OtherAmounts29) AddAccruedInterestAmount() *AmountAndDirection44 {
	o.AccruedInterestAmount = new(AmountAndDirection44)
	return o.AccruedInterestAmount
}

func (o *OtherAmounts29) AddChargesFees() *AmountAndDirection44 {
	o.ChargesFees = new(AmountAndDirection44)
	return o.ChargesFees
}

func (o *OtherAmounts29) AddCountryNationalFederalTax() *AmountAndDirection44 {
	o.CountryNationalFederalTax = new(AmountAndDirection44)
	return o.CountryNationalFederalTax
}

func (o *OtherAmounts29) AddPaymentLevyTax() *AmountAndDirection44 {
	o.PaymentLevyTax = new(AmountAndDirection44)
	return o.PaymentLevyTax
}

func (o *OtherAmounts29) AddLocalTax() *AmountAndDirection44 {
	o.LocalTax = new(AmountAndDirection44)
	return o.LocalTax
}

func (o *OtherAmounts29) AddOther() *AmountAndDirection44 {
	o.Other = new(AmountAndDirection44)
	return o.Other
}

func (o *OtherAmounts29) AddRegulatoryAmount() *AmountAndDirection44 {
	o.RegulatoryAmount = new(AmountAndDirection44)
	return o.RegulatoryAmount
}

func (o *OtherAmounts29) AddShippingAmount() *AmountAndDirection44 {
	o.ShippingAmount = new(AmountAndDirection44)
	return o.ShippingAmount
}

func (o *OtherAmounts29) AddStampDuty() *AmountAndDirection44 {
	o.StampDuty = new(AmountAndDirection44)
	return o.StampDuty
}

func (o *OtherAmounts29) AddStockExchangeTax() *AmountAndDirection44 {
	o.StockExchangeTax = new(AmountAndDirection44)
	return o.StockExchangeTax
}

func (o *OtherAmounts29) AddTransferTax() *AmountAndDirection44 {
	o.TransferTax = new(AmountAndDirection44)
	return o.TransferTax
}

func (o *OtherAmounts29) AddTransactionTax() *AmountAndDirection44 {
	o.TransactionTax = new(AmountAndDirection44)
	return o.TransactionTax
}

func (o *OtherAmounts29) AddValueAddedTax() *AmountAndDirection44 {
	o.ValueAddedTax = new(AmountAndDirection44)
	return o.ValueAddedTax
}

func (o *OtherAmounts29) AddWithholdingTax() *AmountAndDirection44 {
	o.WithholdingTax = new(AmountAndDirection44)
	return o.WithholdingTax
}

func (o *OtherAmounts29) AddConsumptionTax() *AmountAndDirection44 {
	o.ConsumptionTax = new(AmountAndDirection44)
	return o.ConsumptionTax
}

func (o *OtherAmounts29) AddAccruedCapitalisationAmount() *AmountAndDirection44 {
	o.AccruedCapitalisationAmount = new(AmountAndDirection44)
	return o.AccruedCapitalisationAmount
}
