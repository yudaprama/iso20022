package model

// Specifies the value date and the amounts traded in a foreign exchange transaction.
type AmountsAndValueDate1 struct {

	// Currency and amount bought in a foreign exchange trade.
	TradingSideBuyAmount *ActiveOrHistoricCurrencyAndAmount `xml:"TradgSdBuyAmt"`

	// Currency and amount sold in a foreign exchange trade.
	TradingSideSellAmount *ActiveOrHistoricCurrencyAndAmount `xml:"TradgSdSellAmt"`

	// Date on which the trade is settled, ie, the amounts are due.
	SettlementDate *ISODate `xml:"SttlmDt"`
}

func (a *AmountsAndValueDate1) SetTradingSideBuyAmount(value, currency string) {
	a.TradingSideBuyAmount = NewActiveOrHistoricCurrencyAndAmount(value, currency)
}

func (a *AmountsAndValueDate1) SetTradingSideSellAmount(value, currency string) {
	a.TradingSideSellAmount = NewActiveOrHistoricCurrencyAndAmount(value, currency)
}

func (a *AmountsAndValueDate1) SetSettlementDate(value string) {
	a.SettlementDate = (*ISODate)(&value)
}
