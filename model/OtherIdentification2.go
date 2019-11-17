package model

// Other accepted financial instrument's identification than ISIN.
type OtherIdentification2 struct {

	// Identification of a security.
	Identification *RestrictedFINXMax31Text `xml:"Id"`

	// Identifies the suffix of the security identification.
	Suffix *Max16Text `xml:"Sfx,omitempty"`

	// Type of the identification.
	Type *IdentificationSource4Choice `xml:"Tp"`
}

func (o *OtherIdentification2) SetIdentification(value string) {
	o.Identification = (*RestrictedFINXMax31Text)(&value)
}

func (o *OtherIdentification2) SetSuffix(value string) {
	o.Suffix = (*Max16Text)(&value)
}

func (o *OtherIdentification2) AddType() *IdentificationSource4Choice {
	o.Type = new(IdentificationSource4Choice)
	return o.Type
}
