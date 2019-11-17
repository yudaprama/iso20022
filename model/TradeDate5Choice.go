package model

// Choice of format for the trade date.
type TradeDate5Choice struct {

	// Date expressed as a ISO date.
	Date *DateAndDateTimeChoice `xml:"Dt"`

	// Date expressed as a code.
	DateCode *TradeDateCode3Choice `xml:"DtCd"`
}

func (t *TradeDate5Choice) AddDate() *DateAndDateTimeChoice {
	t.Date = new(DateAndDateTimeChoice)
	return t.Date
}

func (t *TradeDate5Choice) AddDateCode() *TradeDateCode3Choice {
	t.DateCode = new(TradeDateCode3Choice)
	return t.DateCode
}
