package model

// Choice between different quantity of security formats.
type Quantity23Choice struct {

	// Signed face amount and amortised value of security.
	OriginalAndCurrentFaceAmount *OriginalAndCurrentQuantities7 `xml:"OrgnlAndCurFaceAmt"`

	// Signed quantity of security.
	SignedQuantity *SignedQuantityFormat9 `xml:"SgndQty"`
}

func (q *Quantity23Choice) AddOriginalAndCurrentFaceAmount() *OriginalAndCurrentQuantities7 {
	q.OriginalAndCurrentFaceAmount = new(OriginalAndCurrentQuantities7)
	return q.OriginalAndCurrentFaceAmount
}

func (q *Quantity23Choice) AddSignedQuantity() *SignedQuantityFormat9 {
	q.SignedQuantity = new(SignedQuantityFormat9)
	return q.SignedQuantity
}
