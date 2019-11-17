package model

// Choice between different quantity of security formats.
type Quantity18Choice struct {

	// Signed face amount and amortised value of security.
	OriginalAndCurrentFaceAmount *OriginalAndCurrentQuantities6 `xml:"OrgnlAndCurFaceAmt"`

	// Signed quantity of security.
	SignedQuantity *SignedQuantityFormat6 `xml:"SgndQty"`
}

func (q *Quantity18Choice) AddOriginalAndCurrentFaceAmount() *OriginalAndCurrentQuantities6 {
	q.OriginalAndCurrentFaceAmount = new(OriginalAndCurrentQuantities6)
	return q.OriginalAndCurrentFaceAmount
}

func (q *Quantity18Choice) AddSignedQuantity() *SignedQuantityFormat6 {
	q.SignedQuantity = new(SignedQuantityFormat6)
	return q.SignedQuantity
}
