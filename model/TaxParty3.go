package model

// Details about the entity involved in the tax paid or to be paid.
type TaxParty3 struct {

	// Number assigned by a tax authority to an entity.
	TaxIdentification *Max35Text `xml:"TaxId,omitempty"`

	// Type of tax payer.
	TaxType *Max35Text `xml:"TaxTp,omitempty"`

	// Unique identification, as assigned by an organisation, to unambiguously identify a party.
	RegistrationIdentification *Max35Text `xml:"RegnId,omitempty"`

	// Specification of the tax exemption reason.
	TaxExemptionReason []*TaxExemptionReasonFormatChoice `xml:"TaxXmptnRsn,omitempty"`
}

func (t *TaxParty3) SetTaxIdentification(value string) {
	t.TaxIdentification = (*Max35Text)(&value)
}

func (t *TaxParty3) SetTaxType(value string) {
	t.TaxType = (*Max35Text)(&value)
}

func (t *TaxParty3) SetRegistrationIdentification(value string) {
	t.RegistrationIdentification = (*Max35Text)(&value)
}

func (t *TaxParty3) AddTaxExemptionReason() *TaxExemptionReasonFormatChoice {
	newValue := new(TaxExemptionReasonFormatChoice)
	t.TaxExemptionReason = append(t.TaxExemptionReason, newValue)
	return newValue
}
