package model

// Stock lending deadline assigned to a borrower of the stock.
type BorrowerLendingDeadline2 struct {

	// Date/time set as the deadline to respond, with instructions, to an outstanding event, for which the underlying security is out on loan.
	StockLendingDeadline *DateFormat34Choice `xml:"StockLndgDdln"`

	// Party who has borrowed stocks on loan.
	Borrower *PartyIdentification103Choice `xml:"Brrwr"`
}

func (b *BorrowerLendingDeadline2) AddStockLendingDeadline() *DateFormat34Choice {
	b.StockLendingDeadline = new(DateFormat34Choice)
	return b.StockLendingDeadline
}

func (b *BorrowerLendingDeadline2) AddBorrower() *PartyIdentification103Choice {
	b.Borrower = new(PartyIdentification103Choice)
	return b.Borrower
}
