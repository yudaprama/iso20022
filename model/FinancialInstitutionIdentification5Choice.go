package model

// Unique and unambiguous identifier of a financial institution, as assigned under an internationally recognised or proprietary identification scheme.
type FinancialInstitutionIdentification5Choice struct {

	// Code allocated to a financial institution by the ISO 9362 Registration Authority as described in ISO 9362 "Banking - Banking telecommunication messages - Business identifier code (BIC)".
	BIC *BICIdentifier `xml:"BIC"`

	// Unique and unambiguous identifier of a clearing system member, as assigned by the system or system administrator.
	ClearingSystemMemberIdentification *ClearingSystemMemberIdentification3Choice `xml:"ClrSysMmbId"`

	// Identifies the name and address of a financial institution.
	NameAndAddress *NameAndAddress7 `xml:"NmAndAdr"`

	// Unique and unambiguous identifier, as assigned to a financial institution using a proprietary identification scheme.
	ProprietaryIdentification *GenericIdentification3 `xml:"PrtryId"`

	// Unique and unambiguous identification of a financial institution, through a combimation of globally recognised or proprietary identification scheme.
	CombinedIdentification *FinancialInstitutionIdentification3 `xml:"CmbndId"`
}

func (f *FinancialInstitutionIdentification5Choice) SetBIC(value string) {
	f.BIC = (*BICIdentifier)(&value)
}

func (f *FinancialInstitutionIdentification5Choice) AddClearingSystemMemberIdentification() *ClearingSystemMemberIdentification3Choice {
	f.ClearingSystemMemberIdentification = new(ClearingSystemMemberIdentification3Choice)
	return f.ClearingSystemMemberIdentification
}

func (f *FinancialInstitutionIdentification5Choice) AddNameAndAddress() *NameAndAddress7 {
	f.NameAndAddress = new(NameAndAddress7)
	return f.NameAndAddress
}

func (f *FinancialInstitutionIdentification5Choice) AddProprietaryIdentification() *GenericIdentification3 {
	f.ProprietaryIdentification = new(GenericIdentification3)
	return f.ProprietaryIdentification
}

func (f *FinancialInstitutionIdentification5Choice) AddCombinedIdentification() *FinancialInstitutionIdentification3 {
	f.CombinedIdentification = new(FinancialInstitutionIdentification3)
	return f.CombinedIdentification
}
