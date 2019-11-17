package model

// Proprietary or domestic identification scheme that uniquely identifies a security.
type AlternateSecurityIdentification7 struct {

	// Unique and unambiguous identifier of a security.
	Identification *Max35Text `xml:"Id"`

	// Source of the identification, that is, domestic (national) or proprietary.
	IdentificationSource *IdentificationSource1Choice `xml:"IdSrc"`
}

func (a *AlternateSecurityIdentification7) SetIdentification(value string) {
	a.Identification = (*Max35Text)(&value)
}

func (a *AlternateSecurityIdentification7) AddIdentificationSource() *IdentificationSource1Choice {
	a.IdentificationSource = new(IdentificationSource1Choice)
	return a.IdentificationSource
}
