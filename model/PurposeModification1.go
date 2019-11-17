package model

// Specifies the type of change to purpose.
type PurposeModification1 struct {

	// Specifies the type of change.
	ModificationCode *Modification1Code `xml:"ModCd,omitempty"`

	// Purpose.
	Purpose *Max140Text `xml:"Purp"`
}

func (p *PurposeModification1) SetModificationCode(value string) {
	p.ModificationCode = (*Modification1Code)(&value)
}

func (p *PurposeModification1) SetPurpose(value string) {
	p.Purpose = (*Max140Text)(&value)
}
