package model

// Specifies the tax identification related to a service to be billed.
type BillingTaxIdentification1 struct {

	// Value added tax (VAT) registration number as provided by the  region’s local taxing authority.
	VATRegistrationNumber *Max35Text `xml:"VATRegnNb,omitempty"`

	// Tax registration number (TRN) as provided by the tax region’s local taxing authority.
	TaxRegistrationNumber *Max35Text `xml:"TaxRegnNb,omitempty"`

	// Specifies financial institution's contact details for the tax region. This contact works for the financial institution, not the tax region.
	TaxContact *ContactDetails3 `xml:"TaxCtct,omitempty"`
}

func (b *BillingTaxIdentification1) SetVATRegistrationNumber(value string) {
	b.VATRegistrationNumber = (*Max35Text)(&value)
}

func (b *BillingTaxIdentification1) SetTaxRegistrationNumber(value string) {
	b.TaxRegistrationNumber = (*Max35Text)(&value)
}

func (b *BillingTaxIdentification1) AddTaxContact() *ContactDetails3 {
	b.TaxContact = new(ContactDetails3)
	return b.TaxContact
}
