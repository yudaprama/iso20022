package model

// Extract of trade data for an investment fund switch order.
type FundOrderData6 struct {

	// Total amount of money paid /to be paid or received in exchange for the financial instrument in the individual order.
	SettlementAmount *ActiveCurrencyAndAmount `xml:"SttlmAmt,omitempty"`

	// Method by which the transaction is settled.
	SettlementMethod *DeliveryReceiptType2Code `xml:"SttlmMtd,omitempty"`

	// Choice between additional cash in or resulting cash out.
	AdditionalAmount *AdditionalAmount1Choice `xml:"AddtlAmt,omitempty"`

	// Currency from which the quoted currency is converted in an exchange rate calculation.
	// 1 x <UnitCcy> = <XchgRate> x <QtdCcy>.
	UnitCurrency *ActiveCurrencyCode `xml:"UnitCcy,omitempty"`

	// Currency into which the unit currency is converted in an exchange rate calculation.
	// 1 x <UnitCcy> = <XchgRate> x <QtdCcy>.
	QuotedCurrency *ActiveCurrencyCode `xml:"QtdCcy,omitempty"`
}

func (f *FundOrderData6) SetSettlementAmount(value, currency string) {
	f.SettlementAmount = NewActiveCurrencyAndAmount(value, currency)
}

func (f *FundOrderData6) SetSettlementMethod(value string) {
	f.SettlementMethod = (*DeliveryReceiptType2Code)(&value)
}

func (f *FundOrderData6) AddAdditionalAmount() *AdditionalAmount1Choice {
	f.AdditionalAmount = new(AdditionalAmount1Choice)
	return f.AdditionalAmount
}

func (f *FundOrderData6) SetUnitCurrency(value string) {
	f.UnitCurrency = (*ActiveCurrencyCode)(&value)
}

func (f *FundOrderData6) SetQuotedCurrency(value string) {
	f.QuotedCurrency = (*ActiveCurrencyCode)(&value)
}
