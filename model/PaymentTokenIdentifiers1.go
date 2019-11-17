package model

// Identifier of a token provider requestor.
type PaymentTokenIdentifiers1 struct {

	// Identifier of the token provider.
	ProviderIdentification *Max35Text `xml:"PrvdrId"`

	// Identifier of the token requestor.
	RequestorIdentification *Max35Text `xml:"RqstrId"`
}

func (p *PaymentTokenIdentifiers1) SetProviderIdentification(value string) {
	p.ProviderIdentification = (*Max35Text)(&value)
}

func (p *PaymentTokenIdentifiers1) SetRequestorIdentification(value string) {
	p.RequestorIdentification = (*Max35Text)(&value)
}
