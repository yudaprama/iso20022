package model

// Identifies other amounts pertaining to the transaction.
type OtherAmounts9 struct {

	// Interest amount that has accrued in between coupon payment periods.
	AccruedInterestAmount *AmountAndDirection9 `xml:"AcrdIntrstAmt,omitempty"`

	// Amount of money paid for the provision of financial services that cannot be categorised by another qualifier.
	ChargesFees *AmountAndDirection9 `xml:"ChrgsFees,omitempty"`

	// Amount of country, national or federal tax charged by the jurisdiction in which the account servicer is located.
	CountryNationalFederalTax *AmountAndDirection9 `xml:"CtryNtlFdrlTax,omitempty"`

	// Principal amount of a trade (price multiplied by quantity).
	TradeAmount *AmountAndDirection9 `xml:"TradAmt,omitempty"`

	// Amount of money paid to an executing broker as a commission.
	ExecutingBrokerAmount *AmountAndDirection9 `xml:"ExctgBrkrAmt,omitempty"`

	// Amount of money defined as a discount on a new issue or on a tranche of an existing issue.
	IssueDiscountAllowance *AmountAndDirection9 `xml:"IsseDscntAllwnc,omitempty"`

	// Amount of payment levy tax.
	PaymentLevyTax *AmountAndDirection9 `xml:"PmtLevyTax,omitempty"`

	// Tax charged by the jurisdiction in which the financial instrument settles.
	LocalTax *AmountAndDirection9 `xml:"LclTax,omitempty"`

	// Local tax country specific.
	LocalTaxCountrySpecific *AmountAndDirection9 `xml:"LclTaxCtrySpcfc,omitempty"`

	// Amount of commission paid to a local broker.
	LocalBrokerCommission *AmountAndDirection9 `xml:"LclBrkrComssn,omitempty"`

	// Amount of money deposited by the trading party in a margin account.
	Margin *AmountAndDirection9 `xml:"Mrgn,omitempty"`

	// An amount that is not indicated by a known business denomination.
	Other *AmountAndDirection9 `xml:"Othr,omitempty"`

	// Amount of money paid for delivery by regular post mail.
	PostageAmount *AmountAndDirection9 `xml:"PstgAmt,omitempty"`

	// Amount of money charged by a regulatory authority, for example, Securities and Exchange fees.
	RegulatoryAmount *AmountAndDirection9 `xml:"RgltryAmt,omitempty"`

	// Amount of money (including insurance) paid for delivery by carrier.
	ShippingAmount *AmountAndDirection9 `xml:"ShppgAmt,omitempty"`

	// Amount of drawdown or other reduction from or in addition to the deal price.
	SpecialConcession *AmountAndDirection9 `xml:"SpclCncssn,omitempty"`

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

	// Amount representing the difference between the cost and the current price of a security. In the context of securities settlement, it is the amount paid or received when the instructions are netted or paired off.
	NetGainLoss *AmountAndDirection9 `xml:"NetGnLoss,omitempty"`

	// Amount of consumption tax.
	ConsumptionTax *AmountAndDirection9 `xml:"CsmptnTax,omitempty"`

	// Amount of unpaid interest (on bonds which have defaulted and have subsequently
	// restructured), which is capitalized and added to the original principal amount of the bond.
	AccruedCapitalisationAmount *AmountAndDirection9 `xml:"AcrdCptlstnAmt,omitempty"`

	// Cost of the securities. May be requested in some countries for tax purposes.
	BookValue *AmountAndDirection9 `xml:"BookVal,omitempty"`

	// Value of the collateral available for the delivery settlement process at the account level.
	CollateralMonitorAmount *AmountAndDirection9 `xml:"CollMntrAmt,omitempty"`
}

func (o *OtherAmounts9) AddAccruedInterestAmount() *AmountAndDirection9 {
	o.AccruedInterestAmount = new(AmountAndDirection9)
	return o.AccruedInterestAmount
}

func (o *OtherAmounts9) AddChargesFees() *AmountAndDirection9 {
	o.ChargesFees = new(AmountAndDirection9)
	return o.ChargesFees
}

func (o *OtherAmounts9) AddCountryNationalFederalTax() *AmountAndDirection9 {
	o.CountryNationalFederalTax = new(AmountAndDirection9)
	return o.CountryNationalFederalTax
}

func (o *OtherAmounts9) AddTradeAmount() *AmountAndDirection9 {
	o.TradeAmount = new(AmountAndDirection9)
	return o.TradeAmount
}

func (o *OtherAmounts9) AddExecutingBrokerAmount() *AmountAndDirection9 {
	o.ExecutingBrokerAmount = new(AmountAndDirection9)
	return o.ExecutingBrokerAmount
}

func (o *OtherAmounts9) AddIssueDiscountAllowance() *AmountAndDirection9 {
	o.IssueDiscountAllowance = new(AmountAndDirection9)
	return o.IssueDiscountAllowance
}

func (o *OtherAmounts9) AddPaymentLevyTax() *AmountAndDirection9 {
	o.PaymentLevyTax = new(AmountAndDirection9)
	return o.PaymentLevyTax
}

func (o *OtherAmounts9) AddLocalTax() *AmountAndDirection9 {
	o.LocalTax = new(AmountAndDirection9)
	return o.LocalTax
}

func (o *OtherAmounts9) AddLocalTaxCountrySpecific() *AmountAndDirection9 {
	o.LocalTaxCountrySpecific = new(AmountAndDirection9)
	return o.LocalTaxCountrySpecific
}

func (o *OtherAmounts9) AddLocalBrokerCommission() *AmountAndDirection9 {
	o.LocalBrokerCommission = new(AmountAndDirection9)
	return o.LocalBrokerCommission
}

func (o *OtherAmounts9) AddMargin() *AmountAndDirection9 {
	o.Margin = new(AmountAndDirection9)
	return o.Margin
}

func (o *OtherAmounts9) AddOther() *AmountAndDirection9 {
	o.Other = new(AmountAndDirection9)
	return o.Other
}

func (o *OtherAmounts9) AddPostageAmount() *AmountAndDirection9 {
	o.PostageAmount = new(AmountAndDirection9)
	return o.PostageAmount
}

func (o *OtherAmounts9) AddRegulatoryAmount() *AmountAndDirection9 {
	o.RegulatoryAmount = new(AmountAndDirection9)
	return o.RegulatoryAmount
}

func (o *OtherAmounts9) AddShippingAmount() *AmountAndDirection9 {
	o.ShippingAmount = new(AmountAndDirection9)
	return o.ShippingAmount
}

func (o *OtherAmounts9) AddSpecialConcession() *AmountAndDirection9 {
	o.SpecialConcession = new(AmountAndDirection9)
	return o.SpecialConcession
}

func (o *OtherAmounts9) AddStampDuty() *AmountAndDirection9 {
	o.StampDuty = new(AmountAndDirection9)
	return o.StampDuty
}

func (o *OtherAmounts9) AddStockExchangeTax() *AmountAndDirection9 {
	o.StockExchangeTax = new(AmountAndDirection9)
	return o.StockExchangeTax
}

func (o *OtherAmounts9) AddTransferTax() *AmountAndDirection9 {
	o.TransferTax = new(AmountAndDirection9)
	return o.TransferTax
}

func (o *OtherAmounts9) AddTransactionTax() *AmountAndDirection9 {
	o.TransactionTax = new(AmountAndDirection9)
	return o.TransactionTax
}

func (o *OtherAmounts9) AddValueAddedTax() *AmountAndDirection9 {
	o.ValueAddedTax = new(AmountAndDirection9)
	return o.ValueAddedTax
}

func (o *OtherAmounts9) AddWithholdingTax() *AmountAndDirection9 {
	o.WithholdingTax = new(AmountAndDirection9)
	return o.WithholdingTax
}

func (o *OtherAmounts9) AddNetGainLoss() *AmountAndDirection9 {
	o.NetGainLoss = new(AmountAndDirection9)
	return o.NetGainLoss
}

func (o *OtherAmounts9) AddConsumptionTax() *AmountAndDirection9 {
	o.ConsumptionTax = new(AmountAndDirection9)
	return o.ConsumptionTax
}

func (o *OtherAmounts9) AddAccruedCapitalisationAmount() *AmountAndDirection9 {
	o.AccruedCapitalisationAmount = new(AmountAndDirection9)
	return o.AccruedCapitalisationAmount
}

func (o *OtherAmounts9) AddBookValue() *AmountAndDirection9 {
	o.BookValue = new(AmountAndDirection9)
	return o.BookValue
}

func (o *OtherAmounts9) AddCollateralMonitorAmount() *AmountAndDirection9 {
	o.CollateralMonitorAmount = new(AmountAndDirection9)
	return o.CollateralMonitorAmount
}
