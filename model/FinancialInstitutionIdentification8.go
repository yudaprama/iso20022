package model

// Set of elements used to identify a financial institution.
type FinancialInstitutionIdentification8 struct {

	// Code allocated to a financial institution by the ISO 9362 Registration Authority as described in ISO 9362 "Banking - Banking telecommunication messages - Business identifier code (BIC)".
	BICFI *BICFIIdentifier `xml:"BICFI,omitempty"`

	// Information used to identify a member within a clearing system.
	ClearingSystemMemberIdentification *ClearingSystemMemberIdentification2 `xml:"ClrSysMmbId,omitempty"`

	// Name by which an agent is known and which is usually used to identify that agent.
	Name *Max140Text `xml:"Nm,omitempty"`

	// Information that locates and identifies a specific address, as defined by postal services.
	PostalAddress *PostalAddress6 `xml:"PstlAdr,omitempty"`

	// Unique identification of an agent, as assigned by an institution, using an identification scheme.
	Other *GenericFinancialIdentification1 `xml:"Othr,omitempty"`
}

func (f *FinancialInstitutionIdentification8) SetBICFI(value string) {
	f.BICFI = (*BICFIIdentifier)(&value)
}

func (f *FinancialInstitutionIdentification8) AddClearingSystemMemberIdentification() *ClearingSystemMemberIdentification2 {
	f.ClearingSystemMemberIdentification = new(ClearingSystemMemberIdentification2)
	return f.ClearingSystemMemberIdentification
}

func (f *FinancialInstitutionIdentification8) SetName(value string) {
	f.Name = (*Max140Text)(&value)
}

func (f *FinancialInstitutionIdentification8) AddPostalAddress() *PostalAddress6 {
	f.PostalAddress = new(PostalAddress6)
	return f.PostalAddress
}

func (f *FinancialInstitutionIdentification8) AddOther() *GenericFinancialIdentification1 {
	f.Other = new(GenericFinancialIdentification1)
	return f.Other
}
