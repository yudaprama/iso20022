package model

// Specifies the conditions to be filled in to obtain a valid power of attorney.
type PowerOfAttorneyRequirements2 struct {

	// Specifies whether the power of attorney needs to be validated by some authority.
	LegalRequirement []*PowerOfAttorneyLegalisation1Code `xml:"LglRqrmnt,omitempty"`

	// Specifies the documents needed to obtain a valid power of attorney.
	OtherDocumentation *Max350Text `xml:"OthrDcmnttn,omitempty"`
}

func (p *PowerOfAttorneyRequirements2) AddLegalRequirement(value string) {
	p.LegalRequirement = append(p.LegalRequirement, (*PowerOfAttorneyLegalisation1Code)(&value))
}

func (p *PowerOfAttorneyRequirements2) SetOtherDocumentation(value string) {
	p.OtherDocumentation = (*Max350Text)(&value)
}
