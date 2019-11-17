package model

// Specifies the identification of a billing subservice.
type BillingSubServiceIdentification1 struct {

	// Specifies the qualifier of the sub service.
	Issuer *BillingSubServiceQualifier1Choice `xml:"Issr"`

	// Further defines a financial institution service, through the provision of the value required by the sub service qualifier, such as the actual lockbox number or store number.
	Identification *Max35Text `xml:"Id"`
}

func (b *BillingSubServiceIdentification1) AddIssuer() *BillingSubServiceQualifier1Choice {
	b.Issuer = new(BillingSubServiceQualifier1Choice)
	return b.Issuer
}

func (b *BillingSubServiceIdentification1) SetIdentification(value string) {
	b.Identification = (*Max35Text)(&value)
}
