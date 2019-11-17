package model

// Unique and unambiguous identifier of a financial institution, as assigned under an internationally recognised or proprietary identification scheme.
type FinancialInstitutionIdentification4Choice struct {

	// Code allocated to a financial institution by the ISO 9362 Registration Authority as described in ISO 9362 "Banking - Banking telecommunication messages - Business identifier code (BIC)".
	BIC *BICIdentifier `xml:"BIC"`

	// Name and address of an institution.
	NameAndAddress *NameAndAddress6 `xml:"NmAndAdr"`
}

func (f *FinancialInstitutionIdentification4Choice) SetBIC(value string) {
	f.BIC = (*BICIdentifier)(&value)
}

func (f *FinancialInstitutionIdentification4Choice) AddNameAndAddress() *NameAndAddress6 {
	f.NameAndAddress = new(NameAndAddress6)
	return f.NameAndAddress
}
