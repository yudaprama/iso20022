package model

// Reference information provided by the creditor to allow the identification of the underlying documents.
type CreditorReferenceInformation2 struct {

	// Specifies the type of creditor reference.
	Type *CreditorReferenceType2 `xml:"Tp,omitempty"`

	// Unique reference, as assigned by the creditor, to unambiguously refer to the payment transaction.
	//
	// Usage: If available, the initiating party should provide this reference in the structured remittance information, to enable reconciliation by the creditor upon receipt of the amount of money.
	//
	// If the business context requires the use of a creditor reference or a payment remit identification, and only one identifier can be passed through the end-to-end chain, the creditor's reference or payment remittance identification should be quoted in the end-to-end transaction identification.
	Reference *Max35Text `xml:"Ref,omitempty"`
}

func (c *CreditorReferenceInformation2) AddType() *CreditorReferenceType2 {
	c.Type = new(CreditorReferenceType2)
	return c.Type
}

func (c *CreditorReferenceInformation2) SetReference(value string) {
	c.Reference = (*Max35Text)(&value)
}
