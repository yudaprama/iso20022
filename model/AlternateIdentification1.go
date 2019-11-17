package model

// Alternate identification of a security.
type AlternateIdentification1 struct {

	// Unique and unambiguous identifier of a security.
	Identification *Max35Text `xml:"Id"`

	// Source of the security identification.
	IdentificationSource *IdentificationSource1Choice `xml:"IdSrc"`
}

func (a *AlternateIdentification1) SetIdentification(value string) {
	a.Identification = (*Max35Text)(&value)
}

func (a *AlternateIdentification1) AddIdentificationSource() *IdentificationSource1Choice {
	a.IdentificationSource = new(IdentificationSource1Choice)
	return a.IdentificationSource
}
