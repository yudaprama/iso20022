package model

// Extract of trade data for an investment fund switch order.
type FundOrderData2 struct {

	// Amount of money used to derive the quantity of investment fund units to be redeemed.
	TotalRedemptionAmount *ActiveOrHistoricCurrencyAndAmount `xml:"TtlRedAmt,omitempty"`

	// Amount of money used to derive the quantity of investment fund units to be subscribed.
	TotalSubscriptionAmount *ActiveOrHistoricCurrencyAndAmount `xml:"TtlSbcptAmt,omitempty"`

	// Amount of money to be transferred between the debtor and creditor before bank transaction charges.
	SettlementAmount *ActiveCurrencyAndAmount `xml:"SttlmAmt,omitempty"`

	// Method by which the transaction is settled.
	SettlementMethod *DeliveryReceiptType2Code `xml:"SttlmMtd,omitempty"`

	// Additional amount of money paid by the investor in addition to the switch redemption amount.
	AdditionalCashIn *ActiveOrHistoricCurrencyAndAmount `xml:"AddtlCshIn,omitempty"`

	// Amount of money that results from a switch-out, that is not reinvested in another investment fund, and is repaid to the investor.
	ResultingCashOut *ActiveOrHistoricCurrencyAndAmount `xml:"RsltgCshOut,omitempty"`

	// Currency in which the rate of exchange is expressed in a currency exchange. In the example 1GBP = xxxCUR, the unit currency is GBP.
	UnitCurrency *ActiveOrHistoricCurrencyCode `xml:"UnitCcy,omitempty"`

	// Currency into which the base currency is converted, in a currency exchange.
	QuotedCurrency *ActiveOrHistoricCurrencyCode `xml:"QtdCcy,omitempty"`
}

func (f *FundOrderData2) SetTotalRedemptionAmount(value, currency string) {
	f.TotalRedemptionAmount = NewActiveOrHistoricCurrencyAndAmount(value, currency)
}

func (f *FundOrderData2) SetTotalSubscriptionAmount(value, currency string) {
	f.TotalSubscriptionAmount = NewActiveOrHistoricCurrencyAndAmount(value, currency)
}

func (f *FundOrderData2) SetSettlementAmount(value, currency string) {
	f.SettlementAmount = NewActiveCurrencyAndAmount(value, currency)
}

func (f *FundOrderData2) SetSettlementMethod(value string) {
	f.SettlementMethod = (*DeliveryReceiptType2Code)(&value)
}

func (f *FundOrderData2) SetAdditionalCashIn(value, currency string) {
	f.AdditionalCashIn = NewActiveOrHistoricCurrencyAndAmount(value, currency)
}

func (f *FundOrderData2) SetResultingCashOut(value, currency string) {
	f.ResultingCashOut = NewActiveOrHistoricCurrencyAndAmount(value, currency)
}

func (f *FundOrderData2) SetUnitCurrency(value string) {
	f.UnitCurrency = (*ActiveOrHistoricCurrencyCode)(&value)
}

func (f *FundOrderData2) SetQuotedCurrency(value string) {
	f.QuotedCurrency = (*ActiveOrHistoricCurrencyCode)(&value)
}
