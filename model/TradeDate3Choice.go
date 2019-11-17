package model

// Choice of format for the trade date.
type TradeDate3Choice struct {

	// Trade date expressed as an ISO 20022 code.
	Date *ISODate `xml:"Dt"`

	// Date expressed as a code.
	DateCode *DateType1Code `xml:"DtCd"`
}

func (t *TradeDate3Choice) SetDate(value string) {
	t.Date = (*ISODate)(&value)
}

func (t *TradeDate3Choice) SetDateCode(value string) {
	t.DateCode = (*DateType1Code)(&value)
}
