package model

// Choice of opening balance.
type OpeningBalance1Choice struct {

	// Opening balance for the statement period. It always equals the closing balance of the previous statement.
	First *BalanceQuantity5Choice `xml:"Frst"`

	// Opening balance of this page only. This balance must be the intermediary closing balance of the previous page of the same statement.
	Intermediary *BalanceQuantity5Choice `xml:"Intrmy"`
}

func (o *OpeningBalance1Choice) AddFirst() *BalanceQuantity5Choice {
	o.First = new(BalanceQuantity5Choice)
	return o.First
}

func (o *OpeningBalance1Choice) AddIntermediary() *BalanceQuantity5Choice {
	o.Intermediary = new(BalanceQuantity5Choice)
	return o.Intermediary
}
