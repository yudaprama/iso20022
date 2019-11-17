package model

// Proprietary or domestic identification scheme that uniquely identifies a security.
type AlternateSecurityIdentification2 struct {

	// Identifies the type of financial instrument identifier type.
	Type *Max35Text `xml:"Tp"`

	// Unique and unambiguous identifier of a security.
	Identification *Max35Text `xml:"Id"`
}

func (a *AlternateSecurityIdentification2) SetType(value string) {
	a.Type = (*Max35Text)(&value)
}

func (a *AlternateSecurityIdentification2) SetIdentification(value string) {
	a.Identification = (*Max35Text)(&value)
}
