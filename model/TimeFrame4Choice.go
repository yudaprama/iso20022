package model

// TimeFrame elements that define a period as number of days after an activity.
type TimeFrame4Choice struct {

	// An agreed number of days after the Trade date (T) used to define standard timeframes e.g T+3 settlement period.
	//
	// Where = T is the date that the price is applied to a transaction.
	TradePlus *Number `xml:"TPlus"`

	// An agreed number of days after the renunciation of title documents are received used to define standard timeframes in Redemption e.g R+3 Redemption settlement cycle.
	RenunciationPlus *Number `xml:"RPlus"`
}

func (t *TimeFrame4Choice) SetTradePlus(value string) {
	t.TradePlus = (*Number)(&value)
}

func (t *TimeFrame4Choice) SetRenunciationPlus(value string) {
	t.RenunciationPlus = (*Number)(&value)
}
