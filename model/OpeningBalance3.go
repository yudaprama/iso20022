package model

// Opening balance for the statement period (first opening balance) or of this page (intermediary opening balance).
type OpeningBalance3 struct {

	// Indication that the position is short or long.
	ShortLongIndicator *ShortLong1Code `xml:"ShrtLngInd"`

	// Opening balance for the statement period (first opening balance) or of this page (intermediary opening balance)
	OpeningBalance *OpeningBalance4Choice `xml:"OpngBal"`
}

func (o *OpeningBalance3) SetShortLongIndicator(value string) {
	o.ShortLongIndicator = (*ShortLong1Code)(&value)
}

func (o *OpeningBalance3) AddOpeningBalance() *OpeningBalance4Choice {
	o.OpeningBalance = new(OpeningBalance4Choice)
	return o.OpeningBalance
}
