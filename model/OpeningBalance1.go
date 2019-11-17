package model

// Opening balance for the statement period (first opening balance) or of this page (intermediary opening balance).
type OpeningBalance1 struct {

	// Indication that the position is short or long.
	ShortLongIndicator *ShortLong1Code `xml:"ShrtLngInd"`

	// Opening balance for the statement period (first opening balance) or of this page (intermediary opening balance)
	OpeningBalance *OpeningBalance1Choice `xml:"OpngBal"`
}

func (o *OpeningBalance1) SetShortLongIndicator(value string) {
	o.ShortLongIndicator = (*ShortLong1Code)(&value)
}

func (o *OpeningBalance1) AddOpeningBalance() *OpeningBalance1Choice {
	o.OpeningBalance = new(OpeningBalance1Choice)
	return o.OpeningBalance
}
