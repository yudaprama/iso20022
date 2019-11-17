package model

// Choice between formats for the identification of the financial institution.
type FinancialInstitutionIdentification3Choice struct {

	// Identifies the name and (long) postal address of a financial institution.
	NameAndAddress *NameAndAddress5 `xml:"NmAndAdr"`

	// Code allocated to a financial institution by the ISO 9362 Registration Authority as described in ISO 9362 "Banking - Banking telecommunication messages - Business identifier code (BIC)".
	BIC *BICIdentifier `xml:"BIC"`

	// Choice of identifier for a clearing system member, as assigned by the clearing system. In some clearing systems, the accounts of the clearing system members are also assigned an identifier.
	ClearingSystemMemberIdentification *ClearingSystemMemberIdentificationChoice `xml:"ClrSysMmbId"`

	// Unique and unambiguous identifier, as assigned to a financial institution using a proprietary identification scheme.
	ProprietaryIdentification *SimpleIdentificationInformation `xml:"PrtryId"`
}

func (f *FinancialInstitutionIdentification3Choice) AddNameAndAddress() *NameAndAddress5 {
	f.NameAndAddress = new(NameAndAddress5)
	return f.NameAndAddress
}

func (f *FinancialInstitutionIdentification3Choice) SetBIC(value string) {
	f.BIC = (*BICIdentifier)(&value)
}

func (f *FinancialInstitutionIdentification3Choice) AddClearingSystemMemberIdentification() *ClearingSystemMemberIdentificationChoice {
	f.ClearingSystemMemberIdentification = new(ClearingSystemMemberIdentificationChoice)
	return f.ClearingSystemMemberIdentification
}

func (f *FinancialInstitutionIdentification3Choice) AddProprietaryIdentification() *SimpleIdentificationInformation {
	f.ProprietaryIdentification = new(SimpleIdentificationInformation)
	return f.ProprietaryIdentification
}
