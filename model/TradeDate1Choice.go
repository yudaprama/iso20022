package model

// Choice of format for the trade date.
type TradeDate1Choice struct {

	// Date expressed as a ISO date.
	Date *DateAndDateTimeChoice `xml:"Dt"`

	// Date expressed as a code.
	DateCode *TradeDateCode1Choice `xml:"DtCd"`
}

func (t *TradeDate1Choice) AddDate() *DateAndDateTimeChoice {
	t.Date = new(DateAndDateTimeChoice)
	return t.Date
}

func (t *TradeDate1Choice) AddDateCode() *TradeDateCode1Choice {
	t.DateCode = new(TradeDateCode1Choice)
	return t.DateCode
}
