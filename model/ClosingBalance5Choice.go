package model

// Choice of closing balance.
type ClosingBalance5Choice struct {

	// Sum of the opening balance and all entries booked to the account at the close of the statement period.
	Final *BalanceQuantity12Choice `xml:"Fnl"`

	// Closing balance of this page only. This balance must be the intermediary opening balance of the next page of the same statement.
	Intermediary *BalanceQuantity12Choice `xml:"Intrmy"`
}

func (c *ClosingBalance5Choice) AddFinal() *BalanceQuantity12Choice {
	c.Final = new(BalanceQuantity12Choice)
	return c.Final
}

func (c *ClosingBalance5Choice) AddIntermediary() *BalanceQuantity12Choice {
	c.Intermediary = new(BalanceQuantity12Choice)
	return c.Intermediary
}
