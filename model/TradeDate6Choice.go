package model

// Choice of format for the trade date.
type TradeDate6Choice struct {

	// Date expressed as a ISO date.
	Date *DateAndDateTimeChoice `xml:"Dt"`

	// Date expressed as a code.
	DateCode *TradeDateCode4Choice `xml:"DtCd"`
}

func (t *TradeDate6Choice) AddDate() *DateAndDateTimeChoice {
	t.Date = new(DateAndDateTimeChoice)
	return t.Date
}

func (t *TradeDate6Choice) AddDateCode() *TradeDateCode4Choice {
	t.DateCode = new(TradeDateCode4Choice)
	return t.DateCode
}
