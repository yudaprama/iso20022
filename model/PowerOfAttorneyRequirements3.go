package model

// Specifies the conditions to be filled in to obtain a valid power of attorney.
type PowerOfAttorneyRequirements3 struct {

	// Specifies whether the power of attorney needs to be validated by some authority.
	LegalRequirement []*PowerOfAttorneyLegalisation1Code `xml:"LglRqrmnt,omitempty"`

	// Specifies the documents needed to obtain a valid power of attorney.
	OtherDocumentation *Max350Text `xml:"OthrDcmnttn,omitempty"`

	// Date by which the requested documents must be provided.
	DocumentSubmissionDeadline *DateFormat29Choice `xml:"DocSubmissnDdln,omitempty"`
}

func (p *PowerOfAttorneyRequirements3) AddLegalRequirement(value string) {
	p.LegalRequirement = append(p.LegalRequirement, (*PowerOfAttorneyLegalisation1Code)(&value))
}

func (p *PowerOfAttorneyRequirements3) SetOtherDocumentation(value string) {
	p.OtherDocumentation = (*Max350Text)(&value)
}

func (p *PowerOfAttorneyRequirements3) AddDocumentSubmissionDeadline() *DateFormat29Choice {
	p.DocumentSubmissionDeadline = new(DateFormat29Choice)
	return p.DocumentSubmissionDeadline
}
