package model

// Specifies the type of change to the full legal name.
type FullLegalNameModification1 struct {

	// Specifies the type of change.
	ModificationCode *Modification1Code `xml:"ModCd,omitempty"`

	// Name by which a party is known and which is usually used to identify that party.
	FullLegalName *Max350Text `xml:"FullLglNm"`
}

func (f *FullLegalNameModification1) SetModificationCode(value string) {
	f.ModificationCode = (*Modification1Code)(&value)
}

func (f *FullLegalNameModification1) SetFullLegalName(value string) {
	f.FullLegalName = (*Max350Text)(&value)
}
