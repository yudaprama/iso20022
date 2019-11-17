package model

// Date and amount.
type DateAndAmount1 struct {

	// Date on which the amount is declared or registered.
	Date *ISODate `xml:"Dt"`

	// Amount of money.
	Amount *ActiveCurrencyAndAmount `xml:"Amt"`
}

func (d *DateAndAmount1) SetDate(value string) {
	d.Date = (*ISODate)(&value)
}

func (d *DateAndAmount1) SetAmount(value, currency string) {
	d.Amount = NewActiveCurrencyAndAmount(value, currency)
}
