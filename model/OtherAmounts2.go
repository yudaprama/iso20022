package model

// Identifies other amounts pertaining to the transaction.
type OtherAmounts2 struct {

	// Interest amount that has accrued in between coupon payment periods.
	AccruedInterestAmount *AmountAndDirection9 `xml:"AcrdIntrstAmt,omitempty"`

	// Amount of money paid for the provision of financial services that cannot be categorised by another qualifier.
	ChargesFees *AmountAndDirection9 `xml:"ChrgsFees,omitempty"`

	// Amount of country, national or federal tax charged by the jurisdiction in which the account servicer is located.
	CountryNationalFederalTax *AmountAndDirection9 `xml:"CtryNtlFdrlTax,omitempty"`

	// Amount of payment levy tax.
	PaymentLevyTax *AmountAndDirection9 `xml:"PmtLevyTax,omitempty"`

	// Tax charged by the jurisdiction in which the financial instrument settles.
	LocalTax *AmountAndDirection9 `xml:"LclTax,omitempty"`

	// An amount that is not indicated by a known business denomination.
	Other *AmountAndDirection9 `xml:"Othr,omitempty"`

	// Amount of money paid for delivery by regular post mail.
	PostageAmount *AmountAndDirection9 `xml:"PstgAmt,omitempty"`

	// Amount of money charged by a regulatory authority, for example, Securities and Exchange fees.
	RegulatoryAmount *AmountAndDirection9 `xml:"RgltryAmt,omitempty"`

	// Amount of money (including insurance) paid for delivery by carrier.
	ShippingAmount *AmountAndDirection9 `xml:"ShppgAmt,omitempty"`

	// Amount of stamp duty.
	StampDuty *AmountAndDirection9 `xml:"StmpDty,omitempty"`

	// Amount of stock exchange tax.
	StockExchangeTax *AmountAndDirection9 `xml:"StockXchgTax,omitempty"`

	// Amount of tax levied on a transfer of ownership of financial instrument.
	TransferTax *AmountAndDirection9 `xml:"TrfTax,omitempty"`

	// Amount of transaction tax.
	TransactionTax *AmountAndDirection9 `xml:"TxTax,omitempty"`

	// Amount of value-added tax.
	ValueAddedTax *AmountAndDirection9 `xml:"ValAddedTax,omitempty"`

	// Amount of money that will be withheld by a tax authority.
	WithholdingTax *AmountAndDirection9 `xml:"WhldgTax,omitempty"`

	// Amount of consumption tax.
	ConsumptionTax *AmountAndDirection9 `xml:"CsmptnTax,omitempty"`

	// Amount of unpaid interest (on bonds which have defaulted and have subsequently
	// restructured), which is capitalized and added to the original principal amount of the bond.
	AccruedCapitalisationAmount *AmountAndDirection9 `xml:"AcrdCptlstnAmt,omitempty"`
}

func (o *OtherAmounts2) AddAccruedInterestAmount() *AmountAndDirection9 {
	o.AccruedInterestAmount = new(AmountAndDirection9)
	return o.AccruedInterestAmount
}

func (o *OtherAmounts2) AddChargesFees() *AmountAndDirection9 {
	o.ChargesFees = new(AmountAndDirection9)
	return o.ChargesFees
}

func (o *OtherAmounts2) AddCountryNationalFederalTax() *AmountAndDirection9 {
	o.CountryNationalFederalTax = new(AmountAndDirection9)
	return o.CountryNationalFederalTax
}

func (o *OtherAmounts2) AddPaymentLevyTax() *AmountAndDirection9 {
	o.PaymentLevyTax = new(AmountAndDirection9)
	return o.PaymentLevyTax
}

func (o *OtherAmounts2) AddLocalTax() *AmountAndDirection9 {
	o.LocalTax = new(AmountAndDirection9)
	return o.LocalTax
}

func (o *OtherAmounts2) AddOther() *AmountAndDirection9 {
	o.Other = new(AmountAndDirection9)
	return o.Other
}

func (o *OtherAmounts2) AddPostageAmount() *AmountAndDirection9 {
	o.PostageAmount = new(AmountAndDirection9)
	return o.PostageAmount
}

func (o *OtherAmounts2) AddRegulatoryAmount() *AmountAndDirection9 {
	o.RegulatoryAmount = new(AmountAndDirection9)
	return o.RegulatoryAmount
}

func (o *OtherAmounts2) AddShippingAmount() *AmountAndDirection9 {
	o.ShippingAmount = new(AmountAndDirection9)
	return o.ShippingAmount
}

func (o *OtherAmounts2) AddStampDuty() *AmountAndDirection9 {
	o.StampDuty = new(AmountAndDirection9)
	return o.StampDuty
}

func (o *OtherAmounts2) AddStockExchangeTax() *AmountAndDirection9 {
	o.StockExchangeTax = new(AmountAndDirection9)
	return o.StockExchangeTax
}

func (o *OtherAmounts2) AddTransferTax() *AmountAndDirection9 {
	o.TransferTax = new(AmountAndDirection9)
	return o.TransferTax
}

func (o *OtherAmounts2) AddTransactionTax() *AmountAndDirection9 {
	o.TransactionTax = new(AmountAndDirection9)
	return o.TransactionTax
}

func (o *OtherAmounts2) AddValueAddedTax() *AmountAndDirection9 {
	o.ValueAddedTax = new(AmountAndDirection9)
	return o.ValueAddedTax
}

func (o *OtherAmounts2) AddWithholdingTax() *AmountAndDirection9 {
	o.WithholdingTax = new(AmountAndDirection9)
	return o.WithholdingTax
}

func (o *OtherAmounts2) AddConsumptionTax() *AmountAndDirection9 {
	o.ConsumptionTax = new(AmountAndDirection9)
	return o.ConsumptionTax
}

func (o *OtherAmounts2) AddAccruedCapitalisationAmount() *AmountAndDirection9 {
	o.AccruedCapitalisationAmount = new(AmountAndDirection9)
	return o.AccruedCapitalisationAmount
}
