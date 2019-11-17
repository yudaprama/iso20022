package model

// Choice between TimeFrame elements that define a period as number of days after an activity.
type TimeFrame5Choice struct {

	// An agreed number of days after the Trade date (T) used to define standard timeframes e.g T+3 settlement period.
	//
	// Where = T is the date that the price is applied to a transaction.
	TradePlus *Number `xml:"TPlus"`

	// Indicates whether pre-payment is necessary.
	Prepayment *YesNoIndicator `xml:"Prepmt"`
}

func (t *TimeFrame5Choice) SetTradePlus(value string) {
	t.TradePlus = (*Number)(&value)
}

func (t *TimeFrame5Choice) SetPrepayment(value string) {
	t.Prepayment = (*YesNoIndicator)(&value)
}
