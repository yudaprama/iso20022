package model

// Provides the additional information for an NDF as supplied on a fixing instruction.
type FixingConditions1 struct {

	// The date on which the trade was executed.
	TradeDate *ISODate `xml:"TradDt"`

	// Represents the original reference of the instruction for which the status is given, as assigned by the participant that submitted the foreign exchange trade.
	OriginatorReference *Max35Text `xml:"OrgtrRef"`

	// Reference common to both parties of the trade.
	CommonReference *Max35Text `xml:"CmonRef,omitempty"`

	// Reference to the identification of a previous event in the life of a trade which is amended or cancelled.
	RelatedReference *Max35Text `xml:"RltdRef,omitempty"`

	// Currency and amount bought in a foreign exchange trade.
	TradingSideBuyAmount *ActiveOrHistoricCurrencyAndAmount `xml:"TradgSdBuyAmt"`

	// Currency and amount sold in a foreign exchange trade.
	TradingSideSellAmount *ActiveOrHistoricCurrencyAndAmount `xml:"TradgSdSellAmt"`

	// The value of one currency expressed in relation to another currency. ExchangeRate expresses the ratio between UnitCurrency and QuotedCurrency (ExchangeRate = UnitCurrency/QuotedCurrency).
	ExchangeRate *BaseOneRate `xml:"XchgRate"`
}

func (f *FixingConditions1) SetTradeDate(value string) {
	f.TradeDate = (*ISODate)(&value)
}

func (f *FixingConditions1) SetOriginatorReference(value string) {
	f.OriginatorReference = (*Max35Text)(&value)
}

func (f *FixingConditions1) SetCommonReference(value string) {
	f.CommonReference = (*Max35Text)(&value)
}

func (f *FixingConditions1) SetRelatedReference(value string) {
	f.RelatedReference = (*Max35Text)(&value)
}

func (f *FixingConditions1) SetTradingSideBuyAmount(value, currency string) {
	f.TradingSideBuyAmount = NewActiveOrHistoricCurrencyAndAmount(value, currency)
}

func (f *FixingConditions1) SetTradingSideSellAmount(value, currency string) {
	f.TradingSideSellAmount = NewActiveOrHistoricCurrencyAndAmount(value, currency)
}

func (f *FixingConditions1) SetExchangeRate(value string) {
	f.ExchangeRate = (*BaseOneRate)(&value)
}
