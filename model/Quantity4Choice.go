package model

// Choice between different quantity of security formats.
type Quantity4Choice struct {

	// Signed face amount and amortised value of security.
	OriginalAndCurrentFaceAmount *OriginalAndCurrentQuantities2 `xml:"OrgnlAndCurFaceAmt"`

	// Signed quantity of security.
	SignedQuantity *SignedQuantityFormat2 `xml:"SgndQty"`
}

func (q *Quantity4Choice) AddOriginalAndCurrentFaceAmount() *OriginalAndCurrentQuantities2 {
	q.OriginalAndCurrentFaceAmount = new(OriginalAndCurrentQuantities2)
	return q.OriginalAndCurrentFaceAmount
}

func (q *Quantity4Choice) AddSignedQuantity() *SignedQuantityFormat2 {
	q.SignedQuantity = new(SignedQuantityFormat2)
	return q.SignedQuantity
}
