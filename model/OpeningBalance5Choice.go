package model

// Choice of opening balance.
type OpeningBalance5Choice struct {

	// Opening balance for the statement period. It always equals the closing balance of the previous statement.
	First *BalanceQuantity12Choice `xml:"Frst"`

	// Opening balance of this page only. This balance must be the intermediary closing balance of the previous page of the same statement.
	Intermediary *BalanceQuantity12Choice `xml:"Intrmy"`
}

func (o *OpeningBalance5Choice) AddFirst() *BalanceQuantity12Choice {
	o.First = new(BalanceQuantity12Choice)
	return o.First
}

func (o *OpeningBalance5Choice) AddIntermediary() *BalanceQuantity12Choice {
	o.Intermediary = new(BalanceQuantity12Choice)
	return o.Intermediary
}
