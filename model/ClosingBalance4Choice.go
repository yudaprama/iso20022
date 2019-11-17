package model

// Choice of closing balance.
type ClosingBalance4Choice struct {

	// Sum of the opening balance and all entries booked to the account at the close of the statement period.
	Final *BalanceQuantity8Choice `xml:"Fnl"`

	// Closing balance of this page only. This balance must be the intermediary opening balance of the next page of the same statement.
	Intermediary *BalanceQuantity8Choice `xml:"Intrmy"`
}

func (c *ClosingBalance4Choice) AddFinal() *BalanceQuantity8Choice {
	c.Final = new(BalanceQuantity8Choice)
	return c.Final
}

func (c *ClosingBalance4Choice) AddIntermediary() *BalanceQuantity8Choice {
	c.Intermediary = new(BalanceQuantity8Choice)
	return c.Intermediary
}
