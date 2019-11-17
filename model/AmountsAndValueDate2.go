package model

// Specifies the value date and the amounts traded in a foreign exchange option trade.
type AmountsAndValueDate2 struct {

	// Call amount and currency of a foreign exchange option trade.
	CallAmount *ActiveOrHistoricCurrencyAndAmount `xml:"CallAmt"`

	// Put amount and currency of a foreign exchange option trade.
	PutAmount *ActiveOrHistoricCurrencyAndAmount `xml:"PutAmt"`

	// Date on which the trade is settled, ie, the amounts are due.
	FinalSettlementDate *ISODate `xml:"FnlSttlmDt"`
}

func (a *AmountsAndValueDate2) SetCallAmount(value, currency string) {
	a.CallAmount = NewActiveOrHistoricCurrencyAndAmount(value, currency)
}

func (a *AmountsAndValueDate2) SetPutAmount(value, currency string) {
	a.PutAmount = NewActiveOrHistoricCurrencyAndAmount(value, currency)
}

func (a *AmountsAndValueDate2) SetFinalSettlementDate(value string) {
	a.FinalSettlementDate = (*ISODate)(&value)
}
