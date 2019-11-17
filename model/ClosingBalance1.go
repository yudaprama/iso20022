package model

// Closing balance for the statement period (final closing balance) or of this page (intermediary closing balance).
type ClosingBalance1 struct {

	// Indication that the position is short or long.
	ShortLongIndicator *ShortLong1Code `xml:"ShrtLngInd"`

	// Closing balance for the statement period (final closing balance) or of this page (intermediary closing balance).
	ClosingBalance *ClosingBalance1Choice `xml:"ClsgBal"`
}

func (c *ClosingBalance1) SetShortLongIndicator(value string) {
	c.ShortLongIndicator = (*ShortLong1Code)(&value)
}

func (c *ClosingBalance1) AddClosingBalance() *ClosingBalance1Choice {
	c.ClosingBalance = new(ClosingBalance1Choice)
	return c.ClosingBalance
}
