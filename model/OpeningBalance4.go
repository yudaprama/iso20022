package model

// Opening balance for the statement period (first opening balance) or of this page (intermediary opening balance).
type OpeningBalance4 struct {

	// Indication that the position is short or long.
	ShortLongIndicator *ShortLong1Code `xml:"ShrtLngInd"`

	// Opening balance for the statement period (first opening balance) or of this page (intermediary opening balance)
	OpeningBalance *OpeningBalance5Choice `xml:"OpngBal"`
}

func (o *OpeningBalance4) SetShortLongIndicator(value string) {
	o.ShortLongIndicator = (*ShortLong1Code)(&value)
}

func (o *OpeningBalance4) AddOpeningBalance() *OpeningBalance5Choice {
	o.OpeningBalance = new(OpeningBalance5Choice)
	return o.OpeningBalance
}
