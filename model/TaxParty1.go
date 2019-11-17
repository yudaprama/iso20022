package model

// Details about the entity involved in the tax paid or to be paid.
type TaxParty1 struct {

	// Tax identification number of the creditor.
	TaxIdentification *Max35Text `xml:"TaxId,omitempty"`

	// Unique identification, as assigned by an organisation, to unambiguously identify a party.
	RegistrationIdentification *Max35Text `xml:"RegnId,omitempty"`

	// Type of tax payer.
	TaxType *Max35Text `xml:"TaxTp,omitempty"`
}

func (t *TaxParty1) SetTaxIdentification(value string) {
	t.TaxIdentification = (*Max35Text)(&value)
}

func (t *TaxParty1) SetRegistrationIdentification(value string) {
	t.RegistrationIdentification = (*Max35Text)(&value)
}

func (t *TaxParty1) SetTaxType(value string) {
	t.TaxType = (*Max35Text)(&value)
}
