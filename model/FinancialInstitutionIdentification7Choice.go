package model

// Choice between formats for the identification of the financial institution.
type FinancialInstitutionIdentification7Choice struct {

	// Identifies the name and (long) postal address of a financial institution.
	NameAndAddress *NameAndAddress5 `xml:"NmAndAdr"`

	// Identification of the financial institution expressed as a BIC.
	BICFI *BICFIIdentifier `xml:"BICFI"`

	// Choice of identifier for a clearing system member, as assigned by the clearing system. In some clearing systems, the accounts of the clearing system members are also assigned an identifier.
	ClearingSystemMemberIdentification *ClearingSystemMemberIdentificationChoice `xml:"ClrSysMmbId"`

	// Unique and unambiguous identifier, as assigned to a financial institution using a proprietary identification scheme.
	ProprietaryIdentification *SimpleIdentificationInformation4 `xml:"PrtryId"`
}

func (f *FinancialInstitutionIdentification7Choice) AddNameAndAddress() *NameAndAddress5 {
	f.NameAndAddress = new(NameAndAddress5)
	return f.NameAndAddress
}

func (f *FinancialInstitutionIdentification7Choice) SetBICFI(value string) {
	f.BICFI = (*BICFIIdentifier)(&value)
}

func (f *FinancialInstitutionIdentification7Choice) AddClearingSystemMemberIdentification() *ClearingSystemMemberIdentificationChoice {
	f.ClearingSystemMemberIdentification = new(ClearingSystemMemberIdentificationChoice)
	return f.ClearingSystemMemberIdentification
}

func (f *FinancialInstitutionIdentification7Choice) AddProprietaryIdentification() *SimpleIdentificationInformation4 {
	f.ProprietaryIdentification = new(SimpleIdentificationInformation4)
	return f.ProprietaryIdentification
}
