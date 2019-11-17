package model

// Closing balance for the statement period (final closing balance) or of this page (intermediary closing balance).
type ClosingBalance4 struct {

	// Indication that the position is short or long.
	ShortLongIndicator *ShortLong1Code `xml:"ShrtLngInd"`

	// Closing balance for the statement period (final closing balance) or of this page (intermediary closing balance).
	ClosingBalance *ClosingBalance5Choice `xml:"ClsgBal"`
}

func (c *ClosingBalance4) SetShortLongIndicator(value string) {
	c.ShortLongIndicator = (*ShortLong1Code)(&value)
}

func (c *ClosingBalance4) AddClosingBalance() *ClosingBalance5Choice {
	c.ClosingBalance = new(ClosingBalance5Choice)
	return c.ClosingBalance
}
