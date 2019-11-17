package model

// Specifies the details of the first leg in a two leg transaction process.
type TwoLegTransactionDetails1 struct {

	// Specifies the date/time on which the trade was executed.
	TradeDate *TradeDate1Choice `xml:"TradDt,omitempty"`

	// Unambiguous identification of the reference assigned in the first leg of the transaction.
	OpeningLegIdentification *Max35Text `xml:"OpngLegId,omitempty"`

	// Unambiguous identification of the second leg of the transaction as known by the account owner (or the instructing party acting on its behalf).
	ClosingLegIdentification *Max35Text `xml:"ClsgLegId,omitempty"`

	// Principal amount of a trade (price multiplied by quantity).
	GrossTradeAmount *AmountAndDirection29 `xml:"GrssTradAmt,omitempty"`

	// Identifies other amounts pertaining to the transaction.
	OtherAmounts []*OtherAmounts16 `xml:"OthrAmts,omitempty"`

	// Provides additional information about the second leg in narrative form.
	SecondLegNarrative *Max140Text `xml:"ScndLegNrrtv,omitempty"`

	// Negotiated fixed price of the security to buy it back.
	EndPrice *Price4 `xml:"EndPric,omitempty"`

	// Closing date/time or maturity date/time of the transaction.
	ClosingDate *ClosingDate1Choice `xml:"ClsgDt,omitempty"`

	// Total amount of money to be paid or received in exchange for the securities.  The amount includes the principal with any commissions and fees or accrued interest.
	ClosingSettlementAmount *AmountAndDirection5 `xml:"ClsgSttlmAmt,omitempty"`

	// Processing date of the trading session.
	ProcessingDate *TradeDate4Choice `xml:"PrcgDt,omitempty"`

	// Specifies the type of the second leg transaction.
	TwoLegTransactionType *TwoLegTransactionType1Choice `xml:"TwoLegTxTp,omitempty"`
}

func (t *TwoLegTransactionDetails1) AddTradeDate() *TradeDate1Choice {
	t.TradeDate = new(TradeDate1Choice)
	return t.TradeDate
}

func (t *TwoLegTransactionDetails1) SetOpeningLegIdentification(value string) {
	t.OpeningLegIdentification = (*Max35Text)(&value)
}

func (t *TwoLegTransactionDetails1) SetClosingLegIdentification(value string) {
	t.ClosingLegIdentification = (*Max35Text)(&value)
}

func (t *TwoLegTransactionDetails1) AddGrossTradeAmount() *AmountAndDirection29 {
	t.GrossTradeAmount = new(AmountAndDirection29)
	return t.GrossTradeAmount
}

func (t *TwoLegTransactionDetails1) AddOtherAmounts() *OtherAmounts16 {
	newValue := new(OtherAmounts16)
	t.OtherAmounts = append(t.OtherAmounts, newValue)
	return newValue
}

func (t *TwoLegTransactionDetails1) SetSecondLegNarrative(value string) {
	t.SecondLegNarrative = (*Max140Text)(&value)
}

func (t *TwoLegTransactionDetails1) AddEndPrice() *Price4 {
	t.EndPrice = new(Price4)
	return t.EndPrice
}

func (t *TwoLegTransactionDetails1) AddClosingDate() *ClosingDate1Choice {
	t.ClosingDate = new(ClosingDate1Choice)
	return t.ClosingDate
}

func (t *TwoLegTransactionDetails1) AddClosingSettlementAmount() *AmountAndDirection5 {
	t.ClosingSettlementAmount = new(AmountAndDirection5)
	return t.ClosingSettlementAmount
}

func (t *TwoLegTransactionDetails1) AddProcessingDate() *TradeDate4Choice {
	t.ProcessingDate = new(TradeDate4Choice)
	return t.ProcessingDate
}

func (t *TwoLegTransactionDetails1) AddTwoLegTransactionType() *TwoLegTransactionType1Choice {
	t.TwoLegTransactionType = new(TwoLegTransactionType1Choice)
	return t.TwoLegTransactionType
}
