package model

// Choice of closing balance.
type ClosingBalance1Choice struct {

	// Sum of the opening balance and all entries booked to the account at the close of the statement period.
	Final *BalanceQuantity5Choice `xml:"Fnl"`

	// Closing balance of this page only. This balance must be the intermediary opening balance of the next page of the same statement.
	Intermediary *BalanceQuantity5Choice `xml:"Intrmy"`
}

func (c *ClosingBalance1Choice) AddFinal() *BalanceQuantity5Choice {
	c.Final = new(BalanceQuantity5Choice)
	return c.Final
}

func (c *ClosingBalance1Choice) AddIntermediary() *BalanceQuantity5Choice {
	c.Intermediary = new(BalanceQuantity5Choice)
	return c.Intermediary
}
