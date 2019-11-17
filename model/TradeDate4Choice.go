package model

// expressed as a ISO20022 code.
type TradeDate4Choice struct {

	// Date and time at which the securities are to be traded.
	Date *DateAndDateTime1Choice `xml:"Dt"`

	// Date and time at which the securities are to be traded expressed as a ISO20022 code.
	Value *TradingDateCode1Choice `xml:"Val"`
}

func (t *TradeDate4Choice) AddDate() *DateAndDateTime1Choice {
	t.Date = new(DateAndDateTime1Choice)
	return t.Date
}

func (t *TradeDate4Choice) AddValue() *TradingDateCode1Choice {
	t.Value = new(TradingDateCode1Choice)
	return t.Value
}
