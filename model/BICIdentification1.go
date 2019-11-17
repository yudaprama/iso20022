package model

// Unique and unambiguous identifier of a financial institution, as assigned under an internationally recognised or proprietary identification scheme.
type BICIdentification1 struct {

	// Code allocated to a financial institution by the ISO 9362 Registration Authority as described in ISO 9362 "Banking - Banking telecommunication messages - Business identifier code (BIC)".
	BIC *BICIdentifier `xml:"BIC"`
}

func (b *BICIdentification1) SetBIC(value string) {
	b.BIC = (*BICIdentifier)(&value)
}
