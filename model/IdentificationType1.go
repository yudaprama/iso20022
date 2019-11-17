package model

// Specification of the identification type.
type IdentificationType1 struct {

	// Structured format.
	Structured *PersonIdentificationType1Code `xml:"Strd"`

	// Additional information about the type of identification
	AdditionalInformation *Max350Text `xml:"AddtlInf,omitempty"`
}

func (i *IdentificationType1) SetStructured(value string) {
	i.Structured = (*PersonIdentificationType1Code)(&value)
}

func (i *IdentificationType1) SetAdditionalInformation(value string) {
	i.AdditionalInformation = (*Max350Text)(&value)
}
