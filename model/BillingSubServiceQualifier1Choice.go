package model

// Specifies the qualifier for a subservice.
type BillingSubServiceQualifier1Choice struct {

	// Specifies the contents of the sub service qualifier, in a coded form.
	Code *BillingSubServiceQualifier1Code `xml:"Cd"`

	// Specifies the contents of the sub service qualifier, in a free-text form.
	Proprietary *Max35Text `xml:"Prtry"`
}

func (b *BillingSubServiceQualifier1Choice) SetCode(value string) {
	b.Code = (*BillingSubServiceQualifier1Code)(&value)
}

func (b *BillingSubServiceQualifier1Choice) SetProprietary(value string) {
	b.Proprietary = (*Max35Text)(&value)
}
