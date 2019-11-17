package model

// Identifies other amounts pertaining to the transaction.
type OtherAmounts33 struct {

	// Interest amount that has accrued in between coupon payment periods.
	AccruedInterestAmount *AmountAndDirection58 `xml:"AcrdIntrstAmt,omitempty"`

	// Amount of money paid for the provision of financial services that cannot be categorised by another qualifier.
	ChargesFees *AmountAndDirection58 `xml:"ChrgsFees,omitempty"`

	// Amount of country, national or federal tax charged by the jurisdiction in which the account servicer is located.
	CountryNationalFederalTax *AmountAndDirection58 `xml:"CtryNtlFdrlTax,omitempty"`

	// Amount of payment levy tax.
	PaymentLevyTax *AmountAndDirection58 `xml:"PmtLevyTax,omitempty"`

	// Tax charged by the jurisdiction in which the financial instrument settles.
	LocalTax *AmountAndDirection58 `xml:"LclTax,omitempty"`

	// An amount that is not indicated by a known business denomination.
	Other *AmountAndDirection58 `xml:"Othr,omitempty"`

	// Amount of money charged by a regulatory authority, for example, Securities and Exchange fees.
	RegulatoryAmount *AmountAndDirection58 `xml:"RgltryAmt,omitempty"`

	// All costs related to the physical delivery of documents such as stamps, postage, carrier fees, insurances or messenger services.
	ShippingAmount *AmountAndDirection58 `xml:"ShppgAmt,omitempty"`

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

	// Amount of consumption tax.
	ConsumptionTax *AmountAndDirection58 `xml:"CsmptnTax,omitempty"`

	// Amount of unpaid interest (on bonds which have defaulted and have subsequently
	// restructured), which is capitalized and added to the original principal amount of the bond.
	AccruedCapitalisationAmount *AmountAndDirection58 `xml:"AcrdCptlstnAmt,omitempty"`
}

func (o *OtherAmounts33) AddAccruedInterestAmount() *AmountAndDirection58 {
	o.AccruedInterestAmount = new(AmountAndDirection58)
	return o.AccruedInterestAmount
}

func (o *OtherAmounts33) AddChargesFees() *AmountAndDirection58 {
	o.ChargesFees = new(AmountAndDirection58)
	return o.ChargesFees
}

func (o *OtherAmounts33) AddCountryNationalFederalTax() *AmountAndDirection58 {
	o.CountryNationalFederalTax = new(AmountAndDirection58)
	return o.CountryNationalFederalTax
}

func (o *OtherAmounts33) AddPaymentLevyTax() *AmountAndDirection58 {
	o.PaymentLevyTax = new(AmountAndDirection58)
	return o.PaymentLevyTax
}

func (o *OtherAmounts33) AddLocalTax() *AmountAndDirection58 {
	o.LocalTax = new(AmountAndDirection58)
	return o.LocalTax
}

func (o *OtherAmounts33) AddOther() *AmountAndDirection58 {
	o.Other = new(AmountAndDirection58)
	return o.Other
}

func (o *OtherAmounts33) AddRegulatoryAmount() *AmountAndDirection58 {
	o.RegulatoryAmount = new(AmountAndDirection58)
	return o.RegulatoryAmount
}

func (o *OtherAmounts33) AddShippingAmount() *AmountAndDirection58 {
	o.ShippingAmount = new(AmountAndDirection58)
	return o.ShippingAmount
}

func (o *OtherAmounts33) AddStampDuty() *AmountAndDirection58 {
	o.StampDuty = new(AmountAndDirection58)
	return o.StampDuty
}

func (o *OtherAmounts33) AddStockExchangeTax() *AmountAndDirection58 {
	o.StockExchangeTax = new(AmountAndDirection58)
	return o.StockExchangeTax
}

func (o *OtherAmounts33) AddTransferTax() *AmountAndDirection58 {
	o.TransferTax = new(AmountAndDirection58)
	return o.TransferTax
}

func (o *OtherAmounts33) AddTransactionTax() *AmountAndDirection58 {
	o.TransactionTax = new(AmountAndDirection58)
	return o.TransactionTax
}

func (o *OtherAmounts33) AddValueAddedTax() *AmountAndDirection58 {
	o.ValueAddedTax = new(AmountAndDirection58)
	return o.ValueAddedTax
}

func (o *OtherAmounts33) AddWithholdingTax() *AmountAndDirection58 {
	o.WithholdingTax = new(AmountAndDirection58)
	return o.WithholdingTax
}

func (o *OtherAmounts33) AddConsumptionTax() *AmountAndDirection58 {
	o.ConsumptionTax = new(AmountAndDirection58)
	return o.ConsumptionTax
}

func (o *OtherAmounts33) AddAccruedCapitalisationAmount() *AmountAndDirection58 {
	o.AccruedCapitalisationAmount = new(AmountAndDirection58)
	return o.AccruedCapitalisationAmount
}
