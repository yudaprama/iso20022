package model

// Set of elements used to identify a financial institution.
type FinancialInstitutionIdentification9 struct {

	// Code allocated to a financial institution by the ISO 9362 Registration Authority as described in ISO 9362 "Banking - Banking telecommunication messages - Business identifier code (BIC)".
	BICFI *BICFIIdentifier `xml:"BICFI,omitempty"`

	// Information used to identify a member within a clearing system.
	ClearingSystemMemberIdentification *ClearingSystemMemberIdentification2 `xml:"ClrSysMmbId,omitempty"`

	// Unique identification of an agent, as assigned by an institution, using an identification scheme.
	Other *GenericFinancialIdentification1 `xml:"Othr,omitempty"`
}

func (f *FinancialInstitutionIdentification9) SetBICFI(value string) {
	f.BICFI = (*BICFIIdentifier)(&value)
}

func (f *FinancialInstitutionIdentification9) AddClearingSystemMemberIdentification() *ClearingSystemMemberIdentification2 {
	f.ClearingSystemMemberIdentification = new(ClearingSystemMemberIdentification2)
	return f.ClearingSystemMemberIdentification
}

func (f *FinancialInstitutionIdentification9) AddOther() *GenericFinancialIdentification1 {
	f.Other = new(GenericFinancialIdentification1)
	return f.Other
}
