package model

// Unique and unambiguous identifier of a financial institution, as assigned under an internationally recognised or proprietary identification scheme.
type FinancialInstitutionIdentification3 struct {

	// Code allocated to a financial institution by the ISO 9362 Registration Authority as described in ISO 9362 "Banking - Banking telecommunication messages - Business identifier code (BIC)".
	BIC *BICIdentifier `xml:"BIC,omitempty"`

	// Unique and unambiguous identifier of a clearing system member, as assigned by the system or system administrator.
	ClearingSystemMemberIdentification *ClearingSystemMemberIdentification3Choice `xml:"ClrSysMmbId,omitempty"`

	// Name by which a party is known and which is usually used to identify that party.
	Name *Max70Text `xml:"Nm,omitempty"`

	// Information that locates and identifies a specific address, as defined by postal services.
	PostalAddress *PostalAddress1 `xml:"PstlAdr,omitempty"`

	// Unique and unambiguous identifier, as assigned to a financial institution using a proprietary identification scheme.
	ProprietaryIdentification *GenericIdentification3 `xml:"PrtryId,omitempty"`
}

func (f *FinancialInstitutionIdentification3) SetBIC(value string) {
	f.BIC = (*BICIdentifier)(&value)
}

func (f *FinancialInstitutionIdentification3) AddClearingSystemMemberIdentification() *ClearingSystemMemberIdentification3Choice {
	f.ClearingSystemMemberIdentification = new(ClearingSystemMemberIdentification3Choice)
	return f.ClearingSystemMemberIdentification
}

func (f *FinancialInstitutionIdentification3) SetName(value string) {
	f.Name = (*Max70Text)(&value)
}

func (f *FinancialInstitutionIdentification3) AddPostalAddress() *PostalAddress1 {
	f.PostalAddress = new(PostalAddress1)
	return f.PostalAddress
}

func (f *FinancialInstitutionIdentification3) AddProprietaryIdentification() *GenericIdentification3 {
	f.ProprietaryIdentification = new(GenericIdentification3)
	return f.ProprietaryIdentification
}
