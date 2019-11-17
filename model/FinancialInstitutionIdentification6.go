package model

// Unique and unambiguous identifier of a financial institution, as assigned under an internationally recognised or proprietary identification scheme.
type FinancialInstitutionIdentification6 struct {

	// Unique and unambiguous identifier of a clearing system member, as assigned by the system or system administrator.
	ClearingSystemMemberIdentification *ClearingSystemMemberIdentification2Choice `xml:"ClrSysMmbId,omitempty"`

	// Unique and unambiguous identifier, as assigned to a financial institution using a proprietary identification scheme.
	ProprietaryIdentification *GenericIdentification4 `xml:"PrtryId,omitempty"`

	// Code allocated to a financial institution by the ISO 9362 Registration Authority as described in ISO 9362 "Banking - Banking telecommunication messages - Business identifier code (BIC)".
	BIC *BICIdentifier `xml:"BIC,omitempty"`
}

func (f *FinancialInstitutionIdentification6) AddClearingSystemMemberIdentification() *ClearingSystemMemberIdentification2Choice {
	f.ClearingSystemMemberIdentification = new(ClearingSystemMemberIdentification2Choice)
	return f.ClearingSystemMemberIdentification
}

func (f *FinancialInstitutionIdentification6) AddProprietaryIdentification() *GenericIdentification4 {
	f.ProprietaryIdentification = new(GenericIdentification4)
	return f.ProprietaryIdentification
}

func (f *FinancialInstitutionIdentification6) SetBIC(value string) {
	f.BIC = (*BICIdentifier)(&value)
}
