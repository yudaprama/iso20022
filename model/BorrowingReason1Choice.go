package model

// Choice of format for the borrowing reason.
type BorrowingReason1Choice struct {

	// Borrowing reason expressed as an ISO 20022 code.
	Code *BorrowingReason1Code `xml:"Cd"`

	// Borrowing reason expressed as a proprietary code.
	Proprietary *GenericIdentification38 `xml:"Prtry"`
}

func (b *BorrowingReason1Choice) SetCode(value string) {
	b.Code = (*BorrowingReason1Code)(&value)
}

func (b *BorrowingReason1Choice) AddProprietary() *GenericIdentification38 {
	b.Proprietary = new(GenericIdentification38)
	return b.Proprietary
}
