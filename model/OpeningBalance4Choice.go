package model

// Choice of opening balance.
type OpeningBalance4Choice struct {

	// Opening balance for the statement period. It always equals the closing balance of the previous statement.
	First *BalanceQuantity8Choice `xml:"Frst"`

	// Opening balance of this page only. This balance must be the intermediary closing balance of the previous page of the same statement.
	Intermediary *BalanceQuantity8Choice `xml:"Intrmy"`
}

func (o *OpeningBalance4Choice) AddFirst() *BalanceQuantity8Choice {
	o.First = new(BalanceQuantity8Choice)
	return o.First
}

func (o *OpeningBalance4Choice) AddIntermediary() *BalanceQuantity8Choice {
	o.Intermediary = new(BalanceQuantity8Choice)
	return o.Intermediary
}
