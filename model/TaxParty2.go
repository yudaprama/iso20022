package model

// Details about the entity involved in the tax paid or to be paid.
type TaxParty2 struct {

	// Tax identification number of the debtor.
	TaxIdentification *Max35Text `xml:"TaxId,omitempty"`

	// Unique identification, as assigned by an organisation, to unambiguously identify a party.
	RegistrationIdentification *Max35Text `xml:"RegnId,omitempty"`

	// Type of tax payer.
	TaxType *Max35Text `xml:"TaxTp,omitempty"`

	// Details of the authorised tax paying party.
	Authorisation *TaxAuthorisation1 `xml:"Authstn,omitempty"`
}

func (t *TaxParty2) SetTaxIdentification(value string) {
	t.TaxIdentification = (*Max35Text)(&value)
}

func (t *TaxParty2) SetRegistrationIdentification(value string) {
	t.RegistrationIdentification = (*Max35Text)(&value)
}

func (t *TaxParty2) SetTaxType(value string) {
	t.TaxType = (*Max35Text)(&value)
}

func (t *TaxParty2) AddAuthorisation() *TaxAuthorisation1 {
	t.Authorisation = new(TaxAuthorisation1)
	return t.Authorisation
}
